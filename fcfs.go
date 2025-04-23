package main

import (
	"fmt"
	"sort"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func float64Perser(a string) float64 {
	temp, _ := strconv.ParseFloat(a, 32)
	return temp
}
func main() {
	content := container.NewVBox()
	myApp := app.New()
	myWindow := myApp.NewWindow("FCFS Algorithm")

	input := widget.NewEntry()
	input.SetPlaceHolder("how many proccess...")
	button := widget.NewButton("Add", func() {
		text := input.Text
		n, _ := strconv.Atoi(text)
		grid := container.New(layout.NewGridLayout(3), widget.NewLabel("PID"), widget.NewLabel("Arrival Time"), widget.NewLabel("Burst Time"))
		content.Add(grid)
		var entries [][]*widget.Entry
		var properties [][]float64
		for i := 0; i < n; i++ {
			rowSlice := []*widget.Entry{widget.NewEntry(), widget.NewEntry()}
			entries = append(entries, rowSlice)
			grid := container.New(layout.NewGridLayout(3), widget.NewLabel(strconv.Itoa(i)), rowSlice[0], rowSlice[1])
			content.Add(grid)
		}
		content.Add(widget.NewButton("Calculate", func() {

			for i := 0; i < n; i++ {
				var temp []float64
				temp = append(temp, float64(i))                        //(0)PID
				temp = append(temp, float64Perser(entries[i][0].Text)) //(1)arrival time
				temp = append(temp, float64Perser(entries[i][1].Text)) //(2)burst time
				temp = append(temp, 0.0, 0.0, 0.0)                     //(3)CT, (4)TAT, (5)WT
				properties = append(properties, temp)

			}
			sort.SliceStable(properties, func(i, j int) bool {
				return properties[i][1] < properties[j][1] //sorting in ascending order as arrival time
			})
			content.Add(container.New(layout.NewGridLayout(6), widget.NewLabel("PID"), widget.NewLabel("AT"), widget.NewLabel("BT"), widget.NewLabel("CT"), widget.NewLabel("TAT"), widget.NewLabel("WT")))
			var totalTime, totalWT, totalTAT float64 = 0.0, 0.0, 0.0
			for i := 0; i < n; i++ {
				var idle float64 = 0.0
				if properties[i][1] > totalTime {
					idle = properties[i][1] - totalTime
					totalTime += idle
				}

				if totalTime > properties[i][1] {
					properties[i][5] = totalTime - properties[i][1] //WT=Tot-AT
				}
				properties[i][4] = properties[i][5] + properties[i][2] //TAT=WT+BT
				totalTime += properties[i][2]                          //+= BT
				properties[i][3] = totalTime                           //CT

				totalTAT += properties[i][4]
				totalWT += properties[i][5]

				content.Add(container.New(layout.NewGridLayout(6), widget.NewLabel(strconv.FormatFloat(properties[i][0], 'f', -1, 32)), widget.NewLabel(strconv.FormatFloat(properties[i][1], 'f', -1, 32)), widget.NewLabel(strconv.FormatFloat(properties[i][2], 'f', -1, 32)), widget.NewLabel(strconv.FormatFloat(properties[i][3], 'f', -1, 32)), widget.NewLabel(strconv.FormatFloat(properties[i][4], 'f', -1, 32)), widget.NewLabel(strconv.FormatFloat(properties[i][5], 'f', -1, 32))))
			}
			content.Add(widget.NewLabel(fmt.Sprintf("Average awaiting time: %.2f", totalWT/float64(n))))
			content.Add(widget.NewLabel(fmt.Sprintf("Average turn around time: %.2f", totalTAT/float64(n))))
			fmt.Println(properties)
		}))
	})
	// arr := [...]*widget.Entry{widget.NewEntry(), widget.NewEntry()}
	content.Add(input)
	content.Add(button)
	/*content.Add(widget.NewButton("Refresh", func() {
		myWindow.Close()

	}))*/

	scrollContainer := container.NewScroll(content)
	scrollContainer.SetMinSize(fyne.NewSize(700, 700))

	myWindow.SetContent(scrollContainer)
	myWindow.ShowAndRun()
}

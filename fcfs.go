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

func float32Perser(a string) float32 {
	temp, _ := strconv.ParseFloat(a, 32)
	temp2 := float32(temp)
	return temp2
}
func main() {
	content := container.NewVBox()
	myApp := app.New()
	myWindow := myApp.NewWindow("Stack Layout")

	input := widget.NewEntry()
	input.SetPlaceHolder("how many proccess...")
	button := widget.NewButton("Add", func() {
		text := input.Text
		n, _ := strconv.Atoi(text)
		grid := container.New(layout.NewGridLayout(3), widget.NewLabel("PID"), widget.NewLabel("Arrival Time"), widget.NewLabel("Burst Time"))
		content.Add(grid)
		var entries [][]*widget.Entry
		var properties [][]float32
		for i := 0; i < n; i++ {
			rowSlice := []*widget.Entry{widget.NewEntry(), widget.NewEntry()}
			entries = append(entries, rowSlice)
			grid := container.New(layout.NewGridLayout(3), widget.NewLabel(strconv.Itoa(i)), rowSlice[0], rowSlice[1])
			content.Add(grid)
		}
		content.Add(widget.NewButton("Calculate", func() {

			for i := 0; i < n; i++ {
				var temp []float32
				temp = append(temp, float32(i))                        //(0)PID
				temp = append(temp, float32Perser(entries[i][0].Text)) //(1)arrival time
				temp = append(temp, float32Perser(entries[i][1].Text)) //(2)burst time
				temp = append(temp, 0.0, 0.0, 0.0)                     //(3)CT, (4)TAT, (5)WT
				properties = append(properties, temp)

			}
			sort.SliceStable(properties, func(i, j int) bool {
				return properties[i][1] < properties[j][1] //sorting in ascending order as arrival time
			})
			var totalTime float32 = 0.0
			for i := 0; i < n; i++ {
				var idle float32 = 0.0
				if properties[i][1] > totalTime {
					idle = properties[0][1] - totalTime
				}

				if totalTime > properties[0][1] {
					properties[i][5] = totalTime - properties[i][1] //WT=Tot-AT
				}
				properties[i][4] = properties[i][5] + properties[i][1] //TAT=WT+AT
				totalTime += idle + properties[i][1]
				properties[i][3] = +totalTime //CT

			}

			fmt.Println(properties)
		}))
	})
	// arr := [...]*widget.Entry{widget.NewEntry(), widget.NewEntry()}
	content.Add(input)
	content.Add(button)

	scrollContainer := container.NewScroll(content)
	scrollContainer.SetMinSize(fyne.NewSize(300, 300))

	myWindow.SetContent(scrollContainer)
	myWindow.ShowAndRun()
}

package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

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
		//		var properties [][]float32
		for i := 0; i < n; i++ {
			rowSlice := []*widget.Entry{widget.NewEntry(), widget.NewEntry()}
			entries = append(entries, rowSlice)
			grid := container.New(layout.NewGridLayout(3), widget.NewLabel(strconv.Itoa(i)), rowSlice[0], rowSlice[1])
			content.Add(grid)
		}
		content.Add(widget.NewButton("Calculate", func() {

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

package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type DirectGeodProbTab struct {
	*fyne.Container
}

func NewDirectGeodProbTab() *DirectGeodProbTab {
	// define coordinates container
	xEntry := widget.NewEntry()
	xEntry.SetPlaceHolder("x")
	yEntry := widget.NewEntry()
	yEntry.SetPlaceHolder("y")
	coordinatesLabel := widget.NewLabel("Координати")
	coordinatesContainer := container.NewBorder(coordinatesLabel, nil, xEntry, yEntry)

	// define length container
	lengthEntry := widget.NewEntry()
	lengthLabel := widget.NewLabel("Відстань")
	lengthContainer := container.NewBorder(lengthLabel, lengthEntry, nil, nil)

	// define azimuth container
	azimuthEntry := widget.NewEntry()
	azimuthLabel := widget.NewLabel("Азимут")
	azimuthContainer := container.NewBorder(azimuthLabel, azimuthEntry, nil, nil)

	// define Result container
	resultLabel := widget.NewLabel("Результат")
	resultValueLabel := widget.NewLabel("Mock result")
	resultContainer := container.NewBorder(resultLabel, resultValueLabel, nil, nil)

	processButton := widget.NewButton("Рахувати", func() {

	})
	c := container.New(layout.NewAdaptiveGridLayout(2),
		coordinatesContainer,
		lengthContainer,
		azimuthContainer,
		resultContainer,
		processButton,
	)
	return &DirectGeodProbTab{c}
}

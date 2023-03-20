package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/containers"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
)

type DirectGeodProbTab struct {
	*fyne.Container
}

func NewDirectGeodProbTab() *DirectGeodProbTab {

	coordinatesContainer := containers.NewCoordinatesContainer()

	// define length container
	lengthEntry := entries.NewNumericalEntry()
	lengthLabel := widget.NewLabel("Відстань")
	lengthContainer := container.NewBorder(lengthLabel, lengthEntry, nil, nil)

	// define azimuth container
	azimuthEntry := entries.NewNumericalEntry()
	azimuthLabel := widget.NewLabel("Азимут")
	azimuthContainer := container.NewBorder(azimuthLabel, azimuthEntry, nil, nil)

	// define Result container
	resultLabel := widget.NewLabel("Результат")
	resultValueLabel := widget.NewLabel("Mock result")
	resultContainer := container.NewBorder(resultLabel, resultValueLabel, nil, nil)

	processButton := widget.NewButton("Рахувати", func() {

	})
	c := container.New(layout.NewGridLayout(2),
		coordinatesContainer.Container,
		lengthContainer,
		azimuthContainer,
		resultContainer,
		processButton,
	)
	return &DirectGeodProbTab{c}
}

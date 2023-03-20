package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
)

type CoordinatesContainer struct {
	*fyne.Container
}

func NewCoordinatesContainer() *CoordinatesContainer {
	xEntry := entries.NewNumericalEntry()
	yEntry := entries.NewNumericalEntry()

	xLabel := widget.NewLabel("x")
	yLabel := widget.NewLabel("y")

	coordsLabel := widget.NewLabel("Координати")

	xContainer := container.New(layout.NewHBoxLayout(), xLabel, xEntry)
	yContainer := container.New(layout.NewHBoxLayout(), yLabel, yEntry)

	coordsContainer := container.New(layout.NewHBoxLayout(), xContainer, yContainer)

	coordsLabelContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), coordsLabel, layout.NewSpacer())

	c := container.New(layout.NewVBoxLayout(), coordsLabelContainer, coordsContainer)
	return &CoordinatesContainer{c}
}

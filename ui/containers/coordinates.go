package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
)

type CoordinatesEntryContainer struct {
	*fyne.Container
	xEntry *entries.NumericalEntry
	yEntry *entries.NumericalEntry
}

func NewCoordinatesEntryContainer() *CoordinatesEntryContainer {
	cont := &CoordinatesEntryContainer{}

	cont.xEntry = entries.NewNumericalEntry()
	cont.yEntry = entries.NewNumericalEntry()

	xLabel := widget.NewLabel("x")
	yLabel := widget.NewLabel("y")

	xContainer := container.New(layout.NewHBoxLayout(), xLabel, cont.xEntry)
	yContainer := container.New(layout.NewHBoxLayout(), yLabel, cont.yEntry)

	coordsContainer := container.New(layout.NewHBoxLayout(), xContainer, yContainer)

	cont.Container = NewCentralNamedContainer("Координати", coordsContainer).Container
	return cont
}

func (c *CoordinatesEntryContainer) GetCoordinates() {

}

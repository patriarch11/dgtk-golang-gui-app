package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
	"strconv"
)

type CoordinatesEntryContainer struct {
	*fyne.Container
	latEntry *entries.NumericalEntry
	lonEntry *entries.NumericalEntry
}

func NewCoordinatesEntryContainer(text string) *CoordinatesEntryContainer {
	cont := &CoordinatesEntryContainer{}

	cont.latEntry = entries.NewNumericalEntry()
	cont.lonEntry = entries.NewNumericalEntry()

	latLabel := widget.NewLabel("Довгота")
	lonLabel := widget.NewLabel("Широта")

	xContainer := container.New(layout.NewHBoxLayout(), latLabel, cont.latEntry)
	yContainer := container.New(layout.NewHBoxLayout(), lonLabel, cont.lonEntry)

	coordsContainer := container.New(layout.NewHBoxLayout(), xContainer, yContainer)

	cont.Container = NewCentralNamedContainer(text, coordsContainer).Container
	return cont
}

func (c *CoordinatesEntryContainer) GetCoordinates() (lat, lon float64) {
	lat, _ = strconv.ParseFloat(c.latEntry.Entry.Text, 64)

	lon, _ = strconv.ParseFloat(c.lonEntry.Entry.Text, 64)
	return
}

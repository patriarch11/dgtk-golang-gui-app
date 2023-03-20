package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
	"log"
	"strconv"
)

type CoordinatesEntryContainer struct {
	*fyne.Container
	latEntry *entries.NumericalEntry
	lonEntry *entries.NumericalEntry
}

func NewCoordinatesEntryContainer() *CoordinatesEntryContainer {
	cont := &CoordinatesEntryContainer{}

	cont.latEntry = entries.NewNumericalEntry()
	cont.lonEntry = entries.NewNumericalEntry()

	latLabel := widget.NewLabel("Довгота")
	lonLabel := widget.NewLabel("Широта")

	xContainer := container.New(layout.NewHBoxLayout(), latLabel, cont.latEntry)
	yContainer := container.New(layout.NewHBoxLayout(), lonLabel, cont.lonEntry)

	coordsContainer := container.New(layout.NewHBoxLayout(), xContainer, yContainer)

	cont.Container = NewCentralNamedContainer("Координати", coordsContainer).Container
	return cont
}

func (c *CoordinatesEntryContainer) GetCoordinates() (lat, lon float64) {
	var err error
	lat, err = strconv.ParseFloat(c.latEntry.Entry.Text, 64)
	if err != nil {
		log.Print(err)
	}
	lon, err = strconv.ParseFloat(c.lonEntry.Entry.Text, 64)
	if err != nil {
		log.Print(err)
	}
	return
}

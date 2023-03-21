package tabs

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/compute"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/containers"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
	"strconv"
)

type DirectGeodesicProbTab struct {
	*fyne.Container
	resultLabel          *containers.DynamicTextLabel
	coordinatesContainer *containers.CoordinatesEntryContainer
	azimuthEntry         *entries.NumericalEntry
	lengthEntry          *entries.NumericalEntry
}

func NewDirectGeodesicProbTab() *DirectGeodesicProbTab {
	tab := &DirectGeodesicProbTab{}

	tab.coordinatesContainer = containers.NewCoordinatesEntryContainer("Known coordinates")

	tab.lengthEntry = entries.NewNumericalEntry()
	lengthContainer := containers.NewLeftNamedContainer("Distance to point", tab.lengthEntry)

	tab.azimuthEntry = entries.NewNumericalEntry()
	azimuthContainer := containers.NewLeftNamedContainer("Azimuth", tab.azimuthEntry)

	tab.resultLabel = containers.NewDynamicTextLabel("Result:")

	processButton := widget.NewButton("Process", func() {
		tab.ResolveProblem()
	})
	compositeContainer := container.New(layout.NewHBoxLayout(),
		container.New(layout.NewVBoxLayout(), lengthContainer.Container, azimuthContainer.Container), // azimuth and length entry containers
		layout.NewSpacer(),
		container.New(layout.NewVBoxLayout(), tab.resultLabel.Container, layout.NewSpacer(), processButton),
	)
	tab.Container = container.New(layout.NewVBoxLayout(), tab.coordinatesContainer.Container, compositeContainer)
	return tab
}

func (t *DirectGeodesicProbTab) ResolveProblem() {
	lat, lon := t.coordinatesContainer.GetCoordinates()
	azimuth, _ := strconv.ParseFloat(t.azimuthEntry.Text, 64)
	dist, _ := strconv.ParseFloat(t.lengthEntry.Text, 64)
	latRes, lonRes := compute.ResolveDirectGeodesicProblem(lat, lon, azimuth, dist)
	t.resultLabel.SetValue(fmt.Sprintf("Coordinates of point:\n%.5f°, %.5f\n", latRes, lonRes))
}

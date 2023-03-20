package tabs

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/containers"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/entries"
	"github.com/pymaxion/geographiclib-go/geodesic"
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

	tab.coordinatesContainer = containers.NewCoordinatesEntryContainer("Координати відомої")

	tab.lengthEntry = entries.NewNumericalEntry()
	lengthContainer := containers.NewLeftNamedContainer("Відстань до точкки", tab.lengthEntry)

	tab.azimuthEntry = entries.NewNumericalEntry()
	azimuthContainer := containers.NewLeftNamedContainer("Азимут", tab.azimuthEntry)

	tab.resultLabel = containers.NewDynamicTextLabel("Результат:")

	processButton := widget.NewButton("Рахувати", func() {
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
	s, _ := strconv.ParseFloat(t.lengthEntry.Text, 64)
	res := geodesic.WGS84.Direct(lat, lon, azimuth, s)
	t.resultLabel.SetValue(fmt.Sprintf("Координати точки:\n%.5f, %.5f\n", res.Lat2, res.Lon2))
}

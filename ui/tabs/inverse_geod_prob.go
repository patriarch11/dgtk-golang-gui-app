package tabs

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/containers"
	"github.com/pymaxion/geographiclib-go/geodesic"
)

type InverseGeodesicProbTab struct {
	*fyne.Container
	resultLabel            *containers.DynamicTextLabel
	coordinatesT1Container *containers.CoordinatesEntryContainer
	coordinatesT2Container *containers.CoordinatesEntryContainer
}

func NewInverseGeodesicProbTab() *InverseGeodesicProbTab {
	tab := &InverseGeodesicProbTab{}

	tab.coordinatesT1Container = containers.NewCoordinatesEntryContainer(
		"Координати точки 1")
	tab.coordinatesT2Container = containers.NewCoordinatesEntryContainer(
		"Координати точки 2")
	tab.resultLabel = containers.NewDynamicTextLabel("Результат:")
	processButton := widget.NewButton("Рахувати", func() {
		tab.ResolveProblem()
	})
	compositeCoordsContainer := container.New(layout.NewVBoxLayout(), tab.coordinatesT1Container.Container,
		tab.coordinatesT2Container.Container)

	tab.Container = container.New(layout.NewVBoxLayout(),
		compositeCoordsContainer,

		container.New(layout.NewVBoxLayout(), tab.resultLabel.Container, processButton),
	)
	return tab
}

func (t *InverseGeodesicProbTab) ResolveProblem() {
	lat1, lon1 := t.coordinatesT1Container.GetCoordinates()
	lat2, lon2 := t.coordinatesT2Container.GetCoordinates()
	res := geodesic.WGS84.Inverse(lat1, lon1, lat2, lon2)

	t.resultLabel.SetValue(fmt.Sprintf("Азимут1: %.2f, Азимут2: %.2f\nВідстань: %.2f",
		res.Azi1, res.Azi2, res.S12))
}

/*
-41.32, 174.81, 40.96, -5.50
-32.06, 115.74, 225, 20000
*/

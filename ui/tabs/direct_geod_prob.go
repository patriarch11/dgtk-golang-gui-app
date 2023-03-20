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
	resultLabel          *containers.DynamicTextLabel
	coordinatesContainer *containers.CoordinatesEntryContainer
	azimuthEntry         *entries.NumericalEntry
	lengthEntry          *entries.NumericalEntry
}

func NewDirectGeodProbTab() *DirectGeodProbTab {
	tab := &DirectGeodProbTab{}

	coordinatesContainer := containers.NewCoordinatesEntryContainer()

	tab.lengthEntry = entries.NewNumericalEntry()
	lengthContainer := containers.NewLeftNamedContainer("Відстань", tab.lengthEntry)

	tab.azimuthEntry = entries.NewNumericalEntry()
	azimuthContainer := containers.NewLeftNamedContainer("Азимут", tab.azimuthEntry)

	tab.resultLabel = containers.NewDynamicTextLabel("Результат:")

	processButton := widget.NewButton("Рахувати", func() {
		tab.ResolveProbe()
	})

	compositeContainer := container.New(layout.NewHBoxLayout(),
		container.New(layout.NewVBoxLayout(), lengthContainer.Container, azimuthContainer.Container), // azimuth and length entry containers
		layout.NewSpacer(),
		container.New(layout.NewVBoxLayout(), tab.resultLabel.Container, layout.NewSpacer(), processButton),
	)
	tab.Container = container.New(layout.NewVBoxLayout(), coordinatesContainer.Container, compositeContainer)

	return tab
}

func (t *DirectGeodProbTab) ResolveProbe() {

}

package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/containers/head"
)

type App struct {
	fyne.App
	c *head.Container
}

func NewUIApp() *App {
	ap := app.New()
	directItem := widget.NewLabel("direct prob")
	reversedItem := widget.NewLabel("reversed prob")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	someItem := container.NewBorder(widget.NewLabel("some prob"), nil, input, nil, widget.NewLabel("china"))

	tabssIt := container.NewAppTabs(
		container.NewTabItem("Direct", directItem),
		container.NewTabItem("reversed", reversedItem),
		container.NewTabItem("some", someItem),
	)
	tabs := map[ui.TabKey]fyne.CanvasObject{
		ui.DIRECT_GEODESIC_PROBLEM:   directItem,
		ui.REVERSED_GEODESIC_PROBLEM: reversedItem,
		ui.SOME_KEY:                  someItem,
	}
	tabItems := []*head.TabItem{
		head.NewTabItem(ui.DIRECT_GEODESIC_PROBLEM, func(container ui.AbstractContainer) *widget.Button {
			b := widget.NewButton("direct", func() {
				container.SetCurrentTab(ui.DIRECT_GEODESIC_PROBLEM)
			})
			return b
		}),
		head.NewTabItem(ui.REVERSED_GEODESIC_PROBLEM, func(container ui.AbstractContainer) *widget.Button {
			b := widget.NewButton("reversed", func() {
				container.SetCurrentTab(ui.REVERSED_GEODESIC_PROBLEM)
			})
			return b
		}),
		head.NewTabItem(ui.SOME_KEY, func(container ui.AbstractContainer) *widget.Button {
			b := widget.NewButton("some key", func() {
				container.SetCurrentTab(ui.SOME_KEY)
			})
			return b
		}),
	}

	c := head.NewHeadContainer(ui.DIRECT_GEODESIC_PROBLEM, tabs, tabItems, tabssIt)
	a := &App{ap, c}
	window := a.NewWindow("test")
	window.SetContent(tabssIt)
	window.Show()
	return a
}

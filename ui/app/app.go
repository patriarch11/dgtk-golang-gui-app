package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/tabs"
)

type App struct {
	fyne.App
	tabs *container.AppTabs
}

func NewUIApp(title string) *App {
	a := &App{}
	a.App = app.New()
	directTab := tabs.NewDirectGeodProbTab()
	a.tabs = container.NewAppTabs(container.NewTabItem("direct", directTab.Container))

	window := a.NewWindow(title)
	window.SetContent(a.tabs)
	window.Show()
	return a
}

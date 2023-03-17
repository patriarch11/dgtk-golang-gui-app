package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

type App struct {
	fyne.App
	*widget.Toolbar
	cont *MainContainer
}

type MainContainerRender struct {
	c *MainContainer
}

func (r *MainContainerRender) Layout(size fyne.Size) {
	r.c.currentPage.Resize(size)
}

func (r *MainContainerRender) MinSize() fyne.Size {
	return r.c.currentPage.MinSize()
}

func (r *MainContainerRender) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *MainContainerRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.c.currentPage}
}

func (r *MainContainerRender) Destroy() {}

func (r *MainContainerRender) Refresh() {
	r.c.currentPage.Refresh()
}

type MainContainer struct {
	widget.BaseWidget
	tb          *widget.Toolbar
	pages       map[string]*fyne.Container
	currentPage *fyne.Container
}

func (c *MainContainer) CreateRenderer() fyne.WidgetRenderer {
	return &MainContainerRender{c: c}
}

func NewUIApp() (a *App) {
	a = &App{App: app.New()}

	a.Toolbar = widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Print("Document create")

			a.cont.currentPage = a.cont.pages["doc"]
			a.cont.Refresh()
		}),
		widget.NewToolbarAction(theme.ComputerIcon(), func() {
			log.Print("computer button")
			a.cont.currentPage = a.cont.pages["comp"]
			a.cont.Refresh()
		}),
	)

	docPage := container.NewBorder(a.Toolbar, nil, nil, nil, widget.NewLabel("document"))
	compPage := container.NewBorder(a.Toolbar, nil, nil, nil, widget.NewLabel("comp"))

	a.cont = &MainContainer{pages: map[string]*fyne.Container{
		"doc":  docPage,
		"comp": compPage,
	},
		currentPage: docPage,
	}

	window := a.NewWindow("Digital geodesic tool kit")
	window.SetContent(a.cont)
	window.Show()
	return a
}

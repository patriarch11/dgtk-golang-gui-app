package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type Renderer struct {
	container AbstractContainer
}

func NewRenderer(c AbstractContainer) *Renderer {
	return &Renderer{container: c}
}

func (r *Renderer) Layout(size fyne.Size) {
	r.container.GetCanvasObject().Resize(size)
}

func (r *Renderer) MinSize() fyne.Size {
	return r.container.GetCanvasObject().MinSize()
}

func (r *Renderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (r *Renderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container.GetCanvasObject()}
}

func (r *Renderer) Refresh() {
	r.container.GetCanvasObject().Refresh()
}

func (r *Renderer) Destroy() {}

package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type NamedEntryContainer struct {
	*fyne.Container
}

func NewCentralNamedContainer(text string, body fyne.CanvasObject) *NamedEntryContainer {
	label := widget.NewLabel(text)
	labelContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), label, layout.NewSpacer())
	c := container.New(layout.NewVBoxLayout(), labelContainer, body)
	return &NamedEntryContainer{c}
}

func NewLeftNamedContainer(text string, body fyne.CanvasObject) *NamedEntryContainer {
	label := widget.NewLabel(text)
	labelContainer := container.New(layout.NewHBoxLayout(), label, layout.NewSpacer(), layout.NewSpacer())
	c := container.New(layout.NewVBoxLayout(), labelContainer, body)
	return &NamedEntryContainer{c}
}

package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type DynamicTextLabel struct {
	*fyne.Container
	value *widget.Label
}

func NewDynamicTextLabel(title string) *DynamicTextLabel {
	l := &DynamicTextLabel{}
	titleLabel := widget.NewLabel(title)
	l.value = widget.NewLabel("")

	l.Container = container.New(layout.NewVBoxLayout(), titleLabel, l.value)
	return l
}

func (l *DynamicTextLabel) SetValue(value string) {
	l.value.SetText(value)
}

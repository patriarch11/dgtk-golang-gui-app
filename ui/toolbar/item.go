package toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Item struct {
	btn *widget.Button
}

func NewItem(btn *widget.Button) widget.ToolbarItem {
	return &Item{btn: btn}
}

func (i *Item) ToolbarObject() fyne.CanvasObject {
	return i.btn
}

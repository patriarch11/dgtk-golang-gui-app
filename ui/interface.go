package ui

import (
	"fyne.io/fyne/v2"
)

type AbstractContainer interface {
	GetCanvasObject() fyne.CanvasObject
	CreateRenderer() fyne.WidgetRenderer
	SetCurrentTab(key TabKey)
}

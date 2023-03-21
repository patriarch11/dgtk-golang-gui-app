package entries

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type ShowEntry struct {
	widget.Entry
}

func NewShowEntry() *ShowEntry {
	entry := &ShowEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *ShowEntry) MinSize() fyne.Size {

	return fyne.Size{Width: 150, Height: 70}
}

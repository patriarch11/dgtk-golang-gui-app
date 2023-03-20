package head

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/patriarch11/dgtk-golang-gui-app/ui"
	"github.com/patriarch11/dgtk-golang-gui-app/ui/toolbar"
)

type TabHandler func(c ui.AbstractContainer) *widget.Button

type TabItem struct {
	Key     ui.TabKey
	Handler TabHandler
}

func NewTabItem(key ui.TabKey, handle func(ui.AbstractContainer) *widget.Button) *TabItem {
	return &TabItem{Key: key, Handler: handle}
}

type Container struct {
	widget.BaseWidget
	T          *container.AppTabs
	Toolbar    *widget.Toolbar
	tabs       map[ui.TabKey]fyne.CanvasObject
	currentTab fyne.CanvasObject
}

func NewHeadContainer(initKey ui.TabKey, tabs map[ui.TabKey]fyne.CanvasObject, items []*TabItem, t *container.AppTabs) *Container {

	c := &Container{tabs: tabs}
	c.T = t
	toolbarItems := make([]widget.ToolbarItem, 0, 0)

	for _, item := range items {
		toolbarItems = append(toolbarItems,
			toolbar.NewItem(item.Handler(c)),
		)
	}

	c.Toolbar = widget.NewToolbar(toolbarItems...)
	c.currentTab = tabs[initKey]

	return c
}

func (hc *Container) GetCanvasObject() fyne.CanvasObject {
	return hc.currentTab
}

func (hc *Container) CreateRenderer() fyne.WidgetRenderer {
	return ui.NewRenderer(hc)
}

func (hc *Container) SetCurrentTab(key ui.TabKey) {
	hc.currentTab = hc.tabs[key]
	hc.currentTab.Refresh()
}

package main

import (
	"github.com/patriarch11/dgtk-golang-gui-app/ui/app"
)

func main() {
	application := app.NewUIApp("test")
	application.App.Run()
}

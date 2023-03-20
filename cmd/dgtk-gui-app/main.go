package main

import "github.com/patriarch11/dgtk-golang-gui-app/ui/app"

func main() {
	application := app.NewUIApp("Digital geodesic tool kit")
	application.App.Run()
}

package main

import (
	"log"
	"os"
	"pluralith-linux-installer/ui"

	"gioui.org/app"
	"gioui.org/unit"
)

func main() {
	go func() {
		// Instantiate Window
		window := app.NewWindow(
			app.Title("Pluralith Installer"),
			app.Size(unit.Dp(380), unit.Dp(450)),
		)

		// Draw UI
		err := ui.Draw(window)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}

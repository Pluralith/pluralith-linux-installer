package main

import (
	"embed"
	"log"
	"os"
	"pluralith-linux-installer/assets"
	"pluralith-linux-installer/backend"
	"pluralith-linux-installer/ui"

	_ "embed"

	"gioui.org/app"
	"gioui.org/unit"
)

// Embedding Assets

//go:embed assets/images/PluralithIcon.png
//go:embed assets/images/DownloadBadge.png
//go:embed assets/images/InstallBadge.png
//go:embed assets/images/CompleteBadge.png
var embedded embed.FS

func initAssets() {
	assets.ImageStore.PluralithIcon, _ = embedded.ReadFile("assets/images/PluralithIcon.png")
	assets.ImageStore.DownloadBadge, _ = embedded.ReadFile("assets/images/DownloadBadge.png")
	assets.ImageStore.InstallBadge, _ = embedded.ReadFile("assets/images/InstallBadge.png")
	assets.ImageStore.CompleteBadge, _ = embedded.ReadFile("assets/images/CompleteBadge.png")
}

func initWindow() {
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

func main() {
	initAssets()
	initWindow()

	backend.StateStore.InitState()
}

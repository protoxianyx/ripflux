package main

import (
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// DownloadService exposes methods directly to the Astro/React frontend
type DownloadService struct{}

func (d *DownloadService) Greet(name string) string {
	return "Hello from Wails 3, " + name
}

func main() {
	// 1. Initialize the Wails v3 Application
	app := application.New(application.Options{
		Name:        "Ripflux",
		Description: "Ripflux Media Downloader",
		Services: []application.Service{
			application.NewService(&DownloadService{}),
		},
	})

	// 2. Create the main application window pointing to Astro dev server
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Ripflux",
		Width:  950,
		Height: 650,
		URL:    "http://localhost:4321", // In dev mode, points to Astro
	})

	// 3. Run the desktop application
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

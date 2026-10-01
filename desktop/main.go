package main

import (
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// DownloadService exposes methods directly to the Astro/React frontend
type DownloadService struct{}

func (d *DownloadService) ChooseDownloadFolder() (string, error) {
	return application.Get().Dialog.OpenFile().
		CanChooseFiles(false).
		CanChooseDirectories(true).
		CanCreateDirectories(true).
		SetTitle("Choose Download folder").
		PromptForSingleSelection()
}

func main() {
	
	app := application.New(application.Options{
		Name:        "Ripflux",
		Description: "Ripflux Media Downloader",
		Services: []application.Service{
			application.NewService(&DownloadService{}),
		},
	})

	
	app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:  "Ripflux",
		Width:  950,
		Height: 650,
		URL:    "http://localhost:4321", // In dev mode, points to Astro
	})

	application.NewService(&DownloadService{})

	
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

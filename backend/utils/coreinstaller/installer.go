package coreinstaller

import (
	"encoding/json"
	"fmt"

	// "io"
	"net/http"
)

type Release struct {
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
}

func Install() {

	var ytdlp_release string = "https://api.github.com/repos/yt-dlp/yt-dlp/releases/latest"
    var downloadURL string

	resp, err := http.Get(ytdlp_release)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	release := Release{}

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return
	}


	for _, asset := range release.Assets {

        if asset.Name == "yt-dlp.exe" {
            downloadURL = asset.BrowserDownloadURL
            break
        }
	}

    fmt.Println(downloadURL)
}

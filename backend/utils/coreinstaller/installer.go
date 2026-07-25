package coreinstaller

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"ripflux/config"

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

    downloadURL := getLatestRelease()

    fmt.Println(downloadURL)
    downloadFile(downloadURL)

}

func getLatestRelease() string {

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
		panic(err)
	}


	for _, asset := range release.Assets {

        if asset.Name == config.YTDLP_EXE {
            downloadURL = asset.BrowserDownloadURL
            break
        }
	}

    return downloadURL
}

func downloadFile(downloadURL string) {

    // fmt.Println(downloadURL)

    resp, err := http.Get(downloadURL)
    	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

    ytdlp_path := filepath.Join(config.BIN_DIR, config.YTDLP_EXE)
    file, err := os.Create(ytdlp_path)
    if err != nil {
        return 
    }
    defer file.Close()

    bytesWritten, err := io.Copy(file, resp.Body)

    fmt.Println("Downloaded", bytesWritten, "bytes")
}

package packagemanager

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	// "ripflux/config"
	"ripflux/config"
	"ripflux/config/paths"
	"ripflux/utils/loggers"

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

type DownloadProps struct {
	downloadURL   string
	latestVersion string
}

func Install() string {

	info := getLatestReleaseInfo()

	fmt.Println(info.downloadURL)
	loggers.TaskLog(config.TEST_LOG_FILE_PATH, "Sucessfully reaching the INstall maker:\n")
	loggers.TaskLog(config.TEST_LOG_FILE_PATH, info.latestVersion)
	downloadFile(info.downloadURL)

	return info.latestVersion

}

func GetLatestVersionInfo() (string, error) {
	info := getLatestReleaseInfo()
	if info.latestVersion == " " {
		loggers.MTaskLog(config.ERROR_LOG_FILE_PATH, true, "Empty DownloadProps")
	}

	return info.latestVersion, nil
}

func getLatestReleaseInfo() DownloadProps {

	var ytdlp_release string = "https://api.github.com/repos/yt-dlp/yt-dlp/releases/latest"
	// var downloadURL string
	// var latestVersion string

	resp, err := http.Get(ytdlp_release)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	release := Release{}
	downloadProps := DownloadProps{}

	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		panic(err)
	}

	for _, asset := range release.Assets {

		if asset.Name == config.YTDLP_EXE {
			downloadProps.downloadURL = asset.BrowserDownloadURL
			downloadProps.latestVersion = release.TagName
			break
		}
	}

	// downloadProps.downloadURL = downloadURL

	return downloadProps
}

func CheckLatestVersion() bool {

	isLatestVersion := false
	return isLatestVersion

}

func downloadFile(downloadURL string) {

	// fmt.Println(downloadURL)

	resp, err := http.Get(downloadURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ytdlp_path := filepath.Join(paths.BIN_DIR, config.YTDLP_EXE)
	file, err := os.Create(ytdlp_path)
	if err != nil {
		return
	}
	defer file.Close()

	bytesWritten, err := io.Copy(file, resp.Body)

	fmt.Println("Downloaded", bytesWritten, "bytes")
}

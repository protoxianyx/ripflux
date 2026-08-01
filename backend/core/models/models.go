package models

type DownloadRequest struct {
    URL        string `json:"url"`
    Format     string `json:"format"`
    Resolution string `json:"resolution"`
}
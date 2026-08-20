package models

type DownloadRequestModel struct {
    URL        string `json:"url"`
    Format     string `json:"format"`
    Resolution string `json:"resolution"`
}
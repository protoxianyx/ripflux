package models

type DownloadRequest struct {
	URL        string `json:"url"`
	Type       string
	Resolution string
}
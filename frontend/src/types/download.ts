export interface DownloadRequest {
  url: string
  format: "Video" | "Audio" | "Format"
  resolution: string
}

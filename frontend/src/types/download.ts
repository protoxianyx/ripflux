export interface DownloadRequest {
  url: string
  format: "Video" | "Audio"
  resolution: string
}

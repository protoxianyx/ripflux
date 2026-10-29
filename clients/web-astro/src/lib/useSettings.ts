export interface SettingsState {
    downloadFolder: string
    defaultFormat: "Video" | "Audio"
    defaultResolution: string
    autoClearTmp: boolean
    concurrentDownloads: number
}
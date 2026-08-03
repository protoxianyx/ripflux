import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Button } from "./ui/button"

import { startDownload } from "@/api/download"
import { useState } from "react"
import { Input } from "./ui/input"

export default function DownloadForm() {
  const [url, setUrl] = useState(" ")
  const [format, setFormat] = useState<"Video" | "Audio">("Video")
  const [resolution, setResolution] = useState(" ")

  async function handleDownload() {


    await startDownload({
      url,
      format,
      resolution,
    })
  }

  return (
    <div className="w-full max-w-3xl space-y-6">
      <div>
        <Input
          value={url}
          onChange={(e) => setUrl(e.target.value)}
          placeholder="Paste a YouTube URL..."
          className="h-14 rounded-md px-5 text-lg shadow-lg"
        />
      </div>
      <div className="flex items-center gap-4">
        <Select value={format} onValueChange={(value) => value && setFormat(value)}>
          <SelectTrigger className="w-48">
            <SelectValue placeholder="Format" />
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="Video">Video</SelectItem>
            <SelectItem value="Audio">Audio</SelectItem>
          </SelectContent>
        </Select>

        <Select
          value={resolution}
          onValueChange={(value) => value && setResolution(value)}
        >
          <SelectTrigger className="w-48">
            <SelectValue placeholder="Resolution" />
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="2160">2160p</SelectItem>
            <SelectItem value="1440">1440p</SelectItem>
            <SelectItem value="1080">1080p</SelectItem>
            <SelectItem value="720">720p</SelectItem>
            <SelectItem value="480">480p</SelectItem>
          </SelectContent>
        </Select>

        <Button className="ml-auto px-8" onClick={handleDownload}>
          Download
        </Button>
      </div>
    </div>
  )
}

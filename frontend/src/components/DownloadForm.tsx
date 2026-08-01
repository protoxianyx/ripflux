import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select";
import { Button } from "./ui/button";

import { startDownload } from "@/api/download"


async function handleDownload() {
  await startDownload({
    url: "https://youtu.be/dQw4w9WgXcQ",
    format: "Video",
    resolution: "1080",
  })
}

export default function DownloadForm() {
    return (
      <div className="flex items-center gap-4">
        <Select>
          <SelectTrigger className="w-48">
            <SelectValue placeholder="Format" />
          </SelectTrigger>

          <SelectContent>
            <SelectItem value="Video">Video</SelectItem>
            <SelectItem value="Audio">Audio</SelectItem>
          </SelectContent>
        </Select>

        <Select>
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

        <Button className="ml-auto px-8" onClick={handleDownload}>Download</Button>
      </div>
    )
}
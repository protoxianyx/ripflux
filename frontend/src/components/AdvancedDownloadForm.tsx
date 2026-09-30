import { useState } from "react"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
  InputGroupText,
} from "./ui/input-group"
import { Button } from "./ui/button"

export default function AdvancedDownloadForm() {
  const [url, setUrl] = useState("")
  const [audioOnly, setAudioOnly] = useState(false)
  const [format, setFormat] = useState("mp3")
  const [resolution, setResolution] = useState("1080")

  function handleSubmit(event: React.SubmitEvent<HTMLFormElement>) {
    event.preventDefault()

    console.log({
      url,
      audioOnly,
      format,
      resolution,
    })
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <InputGroup className="h-12">
        <InputGroupAddon>
          <InputGroupText>URL</InputGroupText>
        </InputGroupAddon>

        <InputGroupInput
          value={url}
          onChange={(event) => setUrl(event.target.value)}
          placeholder="Paste a YouTube URL..."
        />
      </InputGroup>

      <InputGroup className="h-12">
        <InputGroupAddon>
          <InputGroupText>Resolution</InputGroupText>
        </InputGroupAddon>

        <InputGroupInput
          value={resolution}
          onChange={(event) => setResolution(event.target.value)}
        />
      </InputGroup>

      <InputGroup className="h-12">
        <InputGroupAddon>
          <InputGroupText>Audio format</InputGroupText>
        </InputGroupAddon>

        <InputGroupInput
          value={format}
          onChange={(event) => setFormat(event.target.value)}
        />
      </InputGroup>

      <label className="flex items-center gap-2">
        <input
          type="checkbox"
          checked={audioOnly}
          onChange={(event) => setAudioOnly(event.target.checked)}
        />
        Audio only
      </label>

      <Button type="submit">Preview options</Button>
    </form>
  )
}

import { useState } from "react"
import { Input } from "./ui/input"

export default function InputForm() {
  const [url, setUrl] = useState("")

  return (
    <Input
      value={url}
      onChange={(e) => setUrl(e.target.value)}
      placeholder="Paste a YouTube URL..."
      className="h-14 rounded-md px-5 text-lg shadow-lg"
    />
  )
}

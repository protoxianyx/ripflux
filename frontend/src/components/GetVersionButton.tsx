import { GetVerison } from "@/api/version"
import { useState } from "react"
import { Button } from "./ui/button"

export default function GetVersionButton() {
  const [version, setVersion] = useState<string | null>(null)

  async function handleClick() {
    try {
      const result = await GetVerison()
      setVersion(result.version)
    } catch {
      setVersion("Failed")
    }
  }

  return (
    <div className="absolute top-6 right-6 flex items-center gap-3">
      {version && <span className="text-sm">{version}</span>}
      <Button type="button" onClick={handleClick}>
        Get Version
      </Button>
    </div>
  )
}

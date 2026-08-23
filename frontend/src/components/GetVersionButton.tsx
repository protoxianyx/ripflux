import { getVerison, installLatestVersion } from "@/api/version"
// import { useState } from "react"
import { Button } from "./ui/button"
import { toast } from "./ui/toast"

export default function GetVersionButton() {
  // const [version, setVersion] = useState<string | null>(null)

  function versionToast(version: string, latestVerison: string) {
    const description: string = `Current Version: ${version} \n Latest Version: ${latestVerison}`

    function handleUpdater() {
      installLatestVersion()
    }

    toast.add({
      title: "Core Version",
      description: description,
      actionProps: {
        children: "Install Latest Verison",
        onClick() {
          handleUpdater()
        },
      },
    })
  }

  async function handleClick() {
    try {
      const result = await getVerison()
      versionToast(result.version, result.latestVersion)
      console.log(result)
    } catch (error) {
      console.log(error)
      versionToast("Failed to get version", "Failed to get latest version")
    }
  }

  return (
    <div className="absolute top-6 right-6 flex items-center gap-3">
      {/* {version && <span className="text-sm">{version}</span>} */}
      <Button
        type="button"
        onClick={() => {
          handleClick()
        }}
      >
        Get Version
      </Button>
    </div>
  )
}

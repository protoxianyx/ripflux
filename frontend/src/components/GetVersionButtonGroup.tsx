import {
  getLatestVersionInfo,
  getVerison,
  installLatestVersion,
} from "@/api/version"
// import { useState } from "react"
import { Button } from "./ui/button"
import { toast } from "./ui/toast"
import { ButtonGroup } from "./ui/button-group"

export default function GetVersionButtonGroup() {
  // const [version, setVersion] = useState<string | null>(null)

  function versionToast(version: string, latestVerison: string) {
    const description: string = `Current Version: ${version} \n Latest Version: ${latestVerison}`

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

  async function handleUpdater() {
    try {
      await installLatestVersion()
      console.log("Update request completed")
    } catch (error) {
      console.error("Update request failed:", error)
    }
  }

  async function showLatestVerison() {
    const latestVerison = await getLatestVersionInfo()

    toast.add({
      title: "Latest Version: ",
      description: latestVerison.latestVersionInfo,
    })
  }

  return (
    <div className="absolute top-6 right-6">
      {/* {version && <span className="text-sm">{version}</span>} */}
      <ButtonGroup>
        <Button type="button" onClick={handleClick}>
          Get Version
        </Button>
        <Button type="button" onClick={showLatestVerison}>
          {" "}
          Check For Latest Version
        </Button>
      </ButtonGroup>
    </div>
  )
}

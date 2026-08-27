import {
  getLatestVersionInfo,
  getVerison,
  installLatestVersion,
} from "@/api/version"

import { Button } from "./ui/button"
import { toast } from "./ui/toast"
import { ButtonGroup } from "./ui/button-group"
import { AlertBox } from "./AlertBox"
import { useState } from "react"

export default function GetVersionButtonGroup() {
  // const [version, setVersion] = useState("")
  const [alertTitle, setAlertTitle] = useState("")
  const [alertDescription, setAlertDescription] = useState<React.ReactNode>("")
  // const [latestVersion, setLatestVersion] = useState("")

  function versionToast(version: string, latestVerison: string) {
    // const description: string = `Current Version: ${version} \n Latest Version: ${latestVerison}`

    toast.add({
      title: "Core Version",
      description: (
        <span className="space-y-1">
          <span className="flex justify-between gap-6">
            <span>
              Current Version: <strong>{version}</strong>
            </span>
          </span>

          <span className="flex justify-between gap-6">
            <span>
              Latest version: <strong>{latestVerison}</strong>
            </span>
          </span>
        </span>
      ),
      actionProps: {
        children: (
          <span className="flex justify-between space-y-1">
            Install Latest Version
          </span>
        ),
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
    try {
      const latestVerisonResult = await getLatestVersionInfo()
      // setLatestVersion(latestVerisonResult.latestVersionInfo)
      setAlertTitle("Fetched Latest Verison")
      setAlertDescription(
        <div className="space-y-1">
          <div className="flex justify-between gap-6">
            <span>
              Current Version: <strong>{"Yet To be implimented"}</strong>
            </span>
          </div>

          <div className="flex justify-between gap-6">
            <span>
              Latest version:{" "}
              <strong>{latestVerisonResult.latestVersionInfo}</strong>
            </span>
          </div>
        </div>
      )
    } catch {
      setAlertTitle("Version Error")
      setAlertDescription("Could not fetch latest version")
    }

    // toast.add({
    //   title: "Latest Version: ",
    //   description: latestVerison.latestVersionInfo,
    // })

    // const alertTitle = `Fetched Latest Version`
    // const alertDescription = `Latest Version: ${latestVerison}`
  }

  return (
    <div>
      {/* {version && <span className="text-sm">{version}</span>} */}
      <ButtonGroup>
        <Button type="button" onClick={handleClick}>
          Get Version
        </Button>
        <Button type="button" onClick={showLatestVerison}>
          Check For Latest Version
        </Button>
      </ButtonGroup>
      {alertTitle && (
        <div className="fixed right-6 bottom-4 z-50 w-[min(24rem,calc(100vw-3rem))]">
          <AlertBox title={alertTitle} description={alertDescription} />
        </div>
      )}
    </div>
  )
}

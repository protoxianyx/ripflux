import { HttpMethod } from "@/types/httpMethods"
import { apiUrl } from "./client"
import { APIRoutes } from "./routes"


type ApiErrorResponse = {
  error?: string
}

export async function getVerison() {
  
  const response = await fetch(apiUrl(APIRoutes.version), {
    method: HttpMethod.GET,
  })

  if (!response.ok) {
    const errorBody = (await response
      .json()
      .catch(() => null)) as ApiErrorResponse | null
    throw new Error(errorBody?.error ?? "Failed to get version.")
  }

  return response.json() as Promise<{ version: string; latestVersion: string }>
}

export async function installLatestVersion() {
  
  const response = await fetch(apiUrl(APIRoutes.installLatestVersion), {
    method: HttpMethod.POST,
  })
  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to update version.")
  }
}

export async function getLatestVersionInfo() {
  // const route = apiRoutes.latestVersionInfo
  const response = await fetch(apiUrl(APIRoutes.latestVersionInfo), {
    method: HttpMethod.GET,
  })

  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "failed to get latest version")
  }

  return response.json() as Promise<{ latestVersionInfo: string }>
}

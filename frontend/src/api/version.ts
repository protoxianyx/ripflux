import { apiUrl } from "./client"
import { apiRoutes } from "./routes"

type ApiErrorResponse = {
  error?: string
}

export async function getVerison() {
  const route = apiRoutes.version
  const response = await fetch(apiUrl(route.path), {
    method: route.method,
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
  const route = apiRoutes.installLatestVersion
  const response = await fetch(apiUrl(route.path), {
    method: route.method,
  })
  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to update version.")
  }
}

export async function getLatestVersionInfo() {
  const route = apiRoutes.latestVersionInfo
  const response = await fetch(apiUrl(route.path), {
    method: route.method,
  })

  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "failed to get latest version")
  }

  return response.json() as Promise<{ latestVersionInfo: string }>
}

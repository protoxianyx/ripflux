import { HttpMethod } from "@/types/httpMethods"

export async function getVerison() {
  const response = await fetch("http://localhost:8080/version", {
    method: HttpMethod.GET
  })

  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to get version.")
  }

  return response.json() as Promise<{ version: string; latestVersion: string }>
}

export async function installLatestVersion() {
  const response = await fetch("http://localhost:8080/latestVersion", {
    method: HttpMethod.POST,
  })
  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to update version.")
  }
}

export async function getLatestVersionInfo() {
  const response = await fetch("http://localhost:8080/latestVersionInfo", {
    method: HttpMethod.GET,
  })

  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "failed to get latest version")
  }

  return response.json() as Promise<{latestVersionInfo: string}>
}

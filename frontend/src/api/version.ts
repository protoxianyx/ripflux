export async function getVerison() {
  const response = await fetch("http://localhost:8080/version")

  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to get version.")
  }

  return response.json() as Promise<{ version: string; latestVersion: string }>
}

export async function installLatestVersion() {
  const response = await fetch("http://localhost:8080/latestVersion", {
    method: "POST",
  })
  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to update version.")
  }
}

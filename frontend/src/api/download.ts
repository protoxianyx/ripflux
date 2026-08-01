import type { DownloadRequest } from "@/types/download"

export async function startDownload(request: DownloadRequest) {
  console.log(JSON.stringify(request, null, 2))

  const response = await fetch("http://localhost:8080/download", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(request),
  })

  if (!response.ok) {
    throw new Error("Failed to start download.")
  }

  return response.json()
}

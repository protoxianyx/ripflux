import type { DownloadRequest } from "@/types/download"
import { apiUrl } from "./client"
import { APIRoutes } from "./routes"
import { HttpMethod } from "@/types/httpMethods"

export async function startDownload(request: DownloadRequest) {
  console.log(JSON.stringify(request, null, 2))

  const response = await fetch(apiUrl(APIRoutes.download), {
    method: HttpMethod.POST,
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

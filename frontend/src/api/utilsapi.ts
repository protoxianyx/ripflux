import { HttpMethod } from "@/types/httpMethods"
import { apiUrl } from "./client"
import { APIRoutes } from "./routes"

export async function clearTmp() {
  const response = await fetch(apiUrl(APIRoutes.clearTmp), {
    method: HttpMethod.POST,
  })

  if (!response.ok) {
    const error = await response.json().catch(() => null)
    throw new Error(error?.error ?? "Failed to clear temporary files")
  }

  return response.json()
}

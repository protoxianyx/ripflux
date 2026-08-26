import { PUBLIC_API_BASE_URL } from "astro:env/client"

const API_BASE_URL = PUBLIC_API_BASE_URL ?? ""

export function apiUrl(path: string, params?: Record<string, string>): string {
  const url = new URL(path, API_BASE_URL || window.location.origin)

  if (params) {
    Object.entries(params).forEach(([key, value]) => {
      url.searchParams.set(key, value)
    })
  }
  return url.toString()
}

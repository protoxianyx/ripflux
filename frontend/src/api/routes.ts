import { HttpMethod } from "@/types/httpMethods"

export const apiRoutes = {
  version: {
    path: "/version",
    method: HttpMethod.GET,
  },
  installLatestVersion: {
    path: "/latestVersion",
    method: HttpMethod.POST,
  },
  latestVersionInfo: {
    path: "/latestVersionInfo",
    method: HttpMethod.GET,
  },
  download: {
    path: "/download",
    method: HttpMethod.POST,
  },
} as const

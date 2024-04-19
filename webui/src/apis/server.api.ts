import {
  CreateServerRequest,
  CreateServerResponse,
  DuplicateServerRequest,
  DuplicateServerResponse,
  GetServersResponse
} from '@/types/server.type'
import { http, httpPrivate } from '@/utils/http'

export const createServer = (req: CreateServerRequest) => http.post<CreateServerResponse>('/servers', req)
export const duplicateServer = (req: DuplicateServerRequest) =>
  http.post<DuplicateServerResponse>(`/servers/${req.server?.id}/duplicate`, req)
export const deleteServer = (id: number | string) => http.delete<object>(`/servers/${id}`)
export const getServer = (id: number | string) => http.delete<object>(`/servers/${id}`)
export const getServers = (limit: number | string, offset: number | string, signal?: AbortSignal) =>
  http.get<GetServersResponse>(`/servers`, {
    params: {
      limit,
      offset
    },
    signal
  })

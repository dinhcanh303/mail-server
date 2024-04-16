import { GetServersResponse } from '@/types/server.type'
import { http, httpPrivate } from '@/utils/http'

// export const createServer = (post: CreatePostRequest) => http.post<CreatePostResponse>('/servers', post)
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

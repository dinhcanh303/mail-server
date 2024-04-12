import { CreatePostRequest, CreatePostResponse } from '@/types/post.type'
import { http, httpPrivate } from '@/utils/http'

export const createServer = (post: CreatePostRequest) => http.post<CreatePostResponse>('/servers', post)
export const deleteServer = (id: number | string) => http.delete<object>(`/servers/${id}`)
export const getServer = (id: number | string) => http.delete<object>(`/servers/${id}`)
export const getServers = () => http.get<object>(`/servers`)

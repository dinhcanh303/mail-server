import {
  CreateClientRequest,
  CreateClientResponse,
  GetClientsResponse,
  UpdateClientRequest,
  UpdateClientResponse
} from '@/types/client.type'
import { http, httpPrivate } from '@/utils/http'

export const createClient = (client: CreateClientRequest) => http.post<CreateClientResponse>('/clients', client)
export const updateClient = (req: UpdateClientRequest) =>
  http.put<UpdateClientResponse>(`/clients/${req.client.id}`, req)
export const deleteClient = (id: number | string) => http.delete<object>(`/clients/${id}`)
export const getClient = (id: number | string) => http.delete<object>(`/clients/${id}`)
export const getClients = (limit: number | string, offset: number | string, signal?: AbortSignal) =>
  http.get<GetClientsResponse>(`/clients`, {
    params: {
      limit,
      offset
    },
    signal
  })

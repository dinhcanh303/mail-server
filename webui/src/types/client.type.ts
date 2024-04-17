import { Client } from '@/models/Client'

export interface CreateClientRequest {
  name: string
  html: string
}
export interface UpdateClientRequest {
  client: Client
}
export type CreateClientResponse = Client
export type UpdateClientResponse = Client
export interface GetClientsResponse {
  clients: Client[]
}

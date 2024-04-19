import { Client } from '@/models/Client'

export interface CreateClientRequest {
  client: Client
}
export type UpdateClientRequest = CreateClientRequest
export type DuplicateClientRequest = CreateClientRequest
export type CreateClientResponse = CreateClientRequest
export type UpdateClientResponse = CreateClientRequest
export type DuplicateClientResponse = CreateClientRequest
export interface GetClientsResponse {
  clients: Client[]
}

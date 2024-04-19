import { Server } from '@/models/Server'

export interface CreateServerRequest {
  server: Server
}
export type CreateServerResponse = CreateServerRequest
export type DuplicateServerRequest = CreateServerRequest
export type DuplicateServerResponse = CreateServerRequest
export interface GetServersResponse {
  servers: Server[]
}

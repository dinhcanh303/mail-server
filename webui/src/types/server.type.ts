import { Server } from '@/models/Server'

export interface CreateServerRequest {
  name: string
  host: string
  port: number | string
  username: string
  password: string
}
export type CreateServerResponse = Server
export interface GetServersResponse {
  servers: Server[]
}

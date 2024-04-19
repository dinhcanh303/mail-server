export interface Server {
  id?: string
  name: string
  host: string
  port: string
  authProtocol: string
  username: string
  password: string
  fromName: string
  fromAddress: string
  tlsType: string
  tlsSkipVerify: boolean
  maxConnections: number
  idleTimeout: number
  retries: number
  waitTimeout: number
  isDefault?: boolean
  createdAt?: string
  updatedAt?: string
}

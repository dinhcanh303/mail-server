export interface Attachment {
  id: string
  name: string
  host: string
  port: number | string
  username: string
  password: string
  tls: string
  skipTls: boolean
  maxConnections: number | string
  idleTimeout: number | string
  retries: number | string
  waitTimeout: number | string
  createdAt: string
  updatedAt: string
}

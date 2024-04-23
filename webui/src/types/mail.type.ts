export interface TestSendMailConnection {
  host: string
  port: string
  authProtocol: string
  username: string
  password: string
  tlsType: string
  fromName: string
  fromAddress: string
  maxConnections: number
  idleTimeout: number
  retries: number
  waitTimeout: number
  to: string
}

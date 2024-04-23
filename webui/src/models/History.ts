export interface History {
  id?: string
  apiKey: string
  to: string
  subject: string
  cc: string
  bcc: string
  content: object
  status: string
  createdAt?: string
  updatedAt?: string
}

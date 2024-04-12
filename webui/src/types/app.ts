import type { AxiosError } from 'axios'

export type ApiResponse = {
  code: string
  message: string
  detail?: [string]
}

export type ApiError = AxiosError<ApiResponse>

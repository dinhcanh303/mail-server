import { GetHistoriesResponse } from '@/types/history.type'
import { http, httpPrivate } from '@/utils/http'

export const getHistories = (limit: number | string, offset: number | string, signal?: AbortSignal) =>
  http.get<GetHistoriesResponse>(`/histories`, {
    params: {
      limit,
      offset
    },
    signal
  })

import {
  CreateTemplateRequest,
  CreateTemplateResponse,
  GetTemplatesResponse,
  UpdateTemplateRequest,
  UpdateTemplateResponse
} from '@/types/template.type'
import { http, httpPrivate } from '@/utils/http'

export const createTemplate = (template: CreateTemplateRequest) =>
  http.post<CreateTemplateResponse>('/templates', template)
export const updateTemplate = (req: UpdateTemplateRequest) =>
  http.put<UpdateTemplateResponse>(`/templates/${req.template.id}`, req)
export const deleteTemplate = (id: number | string) => http.delete<object>(`/templates/${id}`)
export const getTemplate = (id: number | string) => http.delete<object>(`/templates/${id}`)
export const getTemplatesActive = () => http.get<GetTemplatesResponse>(`/templates/active`)
export const getTemplates = (limit: number | string, offset: number | string, signal?: AbortSignal) =>
  http.get<GetTemplatesResponse>(`/templates`, {
    params: {
      limit,
      offset
    },
    signal
  })

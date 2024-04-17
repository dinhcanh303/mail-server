import { Template } from '@/models/Template'

export interface CreateTemplateRequest {
  name: string
  html: string
}
export interface UpdateTemplateRequest {
  template: Template
}
export type CreateTemplateResponse = Template
export type UpdateTemplateResponse = Template
export interface GetTemplatesResponse {
  templates: Template[]
}

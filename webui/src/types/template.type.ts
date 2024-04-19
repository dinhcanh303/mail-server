import { Template } from '@/models/Template'

export interface CreateTemplateRequest {
  template: Template
}
export interface UpdateTemplateRequest {
  template: Template
}
export interface CreateTemplateResponse {
  template: Template
}
export interface UpdateTemplateResponse {
  template: Template
}
export interface GetTemplatesResponse {
  templates: Template[]
}

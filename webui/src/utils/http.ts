import axios, { AxiosInstance } from 'axios'
import { getTokenAuth } from './auth'
const domainApi = import.meta.env.VITE_DOMAIN_API ?? "localhost"

class Http {
  instance: AxiosInstance
  constructor() {
    this.instance = axios.create({
      baseURL: `https://${domainApi}/api/v1`,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json'
      }
    })
  }
}
const http = new Http().instance

class HttpPrivate {
  instance: AxiosInstance
  constructor() {
    const { accessToken, clientId } = getTokenAuth()
    this.instance = axios.create({
      baseURL: `https://${domainApi}/api/v1`,
      timeout: 10000,
      headers: {
        'Content-Type': 'application/json',
        Authorization: accessToken,
        'Grpc-Metadata-X-Client-Id': clientId
      }
    })
  }
}
const httpPrivate = new HttpPrivate().instance

class HttpPrivateUpload {
  instance: AxiosInstance
  constructor() {
    const { accessToken, clientId } = getTokenAuth()
    this.instance = axios.create({
      baseURL: `https://${domainApi}/api/v1`,
      timeout: 10000,
      headers: {
        'Content-Type': 'multipart/form-data',
        Authorization: accessToken,
        'Grpc-Metadata-X-Client-Id': clientId
      }
    })
  }
}
const httpPrivateUpload = new HttpPrivateUpload().instance
export { http, httpPrivate, httpPrivateUpload }

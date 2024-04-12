import axios, { AxiosError } from 'axios'
import moment from 'moment'
import { useSearchParams } from 'react-router-dom'

export const useParamsString = () => {
  const [searchParams] = useSearchParams()
  const searchParamsObject = Object.fromEntries([...searchParams])
  return searchParamsObject
}
export function isAxiosError<T>(error: unknown): error is AxiosError<T> {
  return axios.isAxiosError(error)
}

export function capitalizeWord(word: string) {
  return word.charAt(0).toUpperCase() + word.slice(1)
}

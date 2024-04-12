import { SignInRequest, SignInResponse } from '@/types/auth.type'
import { http, httpPrivate } from '@/utils/http'

export const signIn = (signIn: SignInRequest) => http.post<SignInResponse>('/auth/signin', signIn)

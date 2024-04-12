import Cookies from 'js-cookie'

const COOKIE_ACCESS_TOKEN_KEY = 'access_token'

const optionsCookies = {
  expires: 30,
  path: '/',
  domain: import.meta.env.VITE_DOMAIN_APP
}
export const saveTokenAuth = (accessToken: string, refreshToken: string, clientId: string) => {
  if (accessToken && refreshToken && clientId) {
    Cookies.set(COOKIE_ACCESS_TOKEN_KEY, accessToken, {
      ...optionsCookies
    })
  }
}
export const removeTokenAuth = () => {
  Cookies.remove(COOKIE_ACCESS_TOKEN_KEY)
}
export const getTokenAuth = () => {
  const accessToken = Cookies.get(COOKIE_ACCESS_TOKEN_KEY)
  return {
    accessToken
  }
}

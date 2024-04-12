import { User } from '@/models/User'
import { create } from 'zustand'
import { devtools } from 'zustand/middleware'

type AuthStore = {
  accessToken: string | undefined
  isAuthenticated: boolean
  isInitialized: boolean
  setToken: (accessToken: string) => void
  logout: () => void
}
export const useAuthStore = create<AuthStore>()(
  devtools(
    (set) => ({
      accessToken: undefined,
      isAuthenticated: false,
      isInitialized: false,
      setToken: (accessToken) => set({ accessToken: accessToken }),
      logout: () => set({ accessToken: undefined })
    }),
    {
      enabled: true,
      name: 'auth-store'
    }
  )
)

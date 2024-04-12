import { User } from '@/models/User'

export interface AuthState {
  isAuthenticated?: boolean
  isInitialized?: boolean
  user: User | null
}

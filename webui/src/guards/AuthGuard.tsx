import Loader from '@/components/loader/Loader'
import { useAuth } from '@/hooks/useAuth'
import { Navigate } from 'react-router-dom'

interface AuthGuardProps {
  children?: any
}
const AuthGuard: React.FC<AuthGuardProps> = ({ children }) => {
  const { isInitialized, isAuthenticated } = useAuth()
  if (!isInitialized) return <Loader />
  if (!isAuthenticated) return <Navigate to='/login' />
  return <>{children}</>
}
export default AuthGuard

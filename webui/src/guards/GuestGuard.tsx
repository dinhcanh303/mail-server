import Loader from '@/components/loader/Loader'
import { useAuth } from '@/hooks/useAuth'
import { Navigate } from 'react-router-dom'

interface GuestGuardProps {
  children?: any
}
const GuestGuard: React.FC<GuestGuardProps> = ({ children }) => {
  const { isInitialized, isAuthenticated } = useAuth()
  if (!isInitialized) return <Loader />
  if (isAuthenticated) return <Navigate to='/guest' />
  return <>{children}</>
}
export default GuestGuard

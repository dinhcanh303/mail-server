import { lazy } from 'react'

const Dashboard = lazy(() => import('@/pages/dashboard/Dashboard'))
const Server = lazy(() => import('@/pages/servers/Server'))

const coreRoutes = [
  {
    path: '/',
    title: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/servers',
    title: 'Servers',
    component: Server
  },
  {
    path: '/templates',
    title: 'Templates',
    component: Dashboard
  }
]

const routes = [...coreRoutes]
export default routes

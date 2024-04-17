import { lazy } from 'react'

const Dashboard = lazy(() => import('@/pages/dashboard/Dashboard'))
const Server = lazy(() => import('@/pages/servers/Server'))
const Template = lazy(() => import('@/pages/templates/Template'))
const Client = lazy(() => import('@/pages/clients/Client'))

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
    component: Template
  },
  {
    path: '/clients',
    title: 'Clients',
    component: Client
  }
]

const routes = [...coreRoutes]
export default routes

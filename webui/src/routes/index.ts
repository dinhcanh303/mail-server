import { lazy } from 'react'

const Dashboard = lazy(() => import('@/pages/dashboard/Dashboard'))
const Server = lazy(() => import('@/pages/servers/Server'))
const Template = lazy(() => import('@/pages/templates/Template'))
const Client = lazy(() => import('@/pages/clients/Client'))
const History = lazy(() => import('@/pages/histories/History'))

const coreRoutes = [
  {
    path: '/',
    title: 'Dashboard',
    component: Dashboard
  },
  {
    path: '/clients',
    title: 'Clients',
    component: Client
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
    path: '/histories',
    title: 'Histories',
    component: History
  }
]

const routes = [...coreRoutes]
export default routes

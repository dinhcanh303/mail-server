import { Suspense, lazy, useEffect, useState } from 'react'
import { Route, Routes } from 'react-router-dom'
import { Toaster } from 'react-hot-toast'

import SignIn from '@/pages/auth/SignIn'
import Loader from '@/components/loader/Loader'
import routes from '@/routes'
import AuthGuard from './guards/AuthGuard'

const DefaultLayout = lazy(() => import('./layout/DefaultLayout'))

function App() {
  const [loading, setLoading] = useState<boolean>(true)

  useEffect(() => {
    setTimeout(() => setLoading(false), 1000)
  }, [])
  return loading ? (
    <Loader />
  ) : (
    <>
      <Toaster position='top-right' reverseOrder={false} containerClassName='overflow-auto' />
      <Routes>
        <Route path='/login' element={<SignIn />} />
        <Route element={<DefaultLayout />}>
          {routes.map((routes, index) => {
            const { path, component: Component } = routes
            return (
              <Route
                key={index}
                path={path}
                element={
                  <Suspense fallback={<Loader />}>
                    {/* <AuthGuard> */}
                    <Component />
                    {/* </AuthGuard> */}
                  </Suspense>
                }
              />
            )
          })}
        </Route>
      </Routes>
    </>
  )
}

export default App

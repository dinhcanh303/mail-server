import React, { useState } from 'react'
import { Outlet } from 'react-router-dom'
import { motion } from 'framer-motion'
import Sidebar from '@/components/sidebar/SideBar'
import Header from '@/components/header/Header'

// eslint-disable-next-line @typescript-eslint/no-empty-interface
interface DefaultLayoutProps {}
const DefaultLayout: React.FC<DefaultLayoutProps> = () => {
  const [sidebarOpen, setSidebarOpen] = useState(false)
  return (
    <main>
      <motion.div
        initial={{ opacity: 0 }}
        whileInView={{ opacity: 1 }}
        viewport={{ once: true }}
        className='bg-gray-100 dark:bg-gray-800'
      >
        <div className='flex h-screen overflow-hidden'>
          <Sidebar sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />

          <div className='relative flex flex-1 flex-col overflow-y-auto overflow-x-hidden'>
            <Header sidebarOpen={sidebarOpen} setSidebarOpen={setSidebarOpen} />
            <main>
              <Outlet />
            </main>
          </div>
        </div>
      </motion.div>
    </main>
  )
}

export default DefaultLayout

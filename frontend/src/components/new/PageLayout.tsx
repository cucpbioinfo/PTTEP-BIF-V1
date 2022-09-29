import React from 'react'
import { Navbar } from './Navbar'

export const PageLayout = ({ children }) => {
  return (
    <div>
      <Navbar />
      <main className="md:px-16 md:py-16 px-8 py-8">{children}</main>
    </div>
  )
}

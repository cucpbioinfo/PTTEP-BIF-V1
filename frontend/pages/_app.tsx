import '../styles/index.css'

import type { AppProps } from 'next/app'
import React from 'react'
import { AuthProvider } from 'contexts/AuthContext'

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <AuthProvider>
      <Component {...pageProps} />
    </AuthProvider>
  )
}

export default App

import { useAuth } from 'contexts/AuthContext'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { Spin } from 'antd'

export const PrivateRoute = (Component) => {
  return function (props) {
    const { isLogin, loading } = useAuth()
    const router = useRouter()

    useEffect(() => {
      if (!isLogin && !loading) {
        router.push('/login')
      }
    }, [isLogin, loading])

    if (!isLogin) {
      return (
        <div className="flex justify-center">
          <Spin spinning delay={500} size="large" />
        </div>
      )
    }
    return <Component {...props} />
  }
}

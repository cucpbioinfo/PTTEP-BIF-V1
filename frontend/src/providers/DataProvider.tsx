import { getMe } from 'api/auth/getMe'
import { eDNAAxios } from 'api/const'
import { useRouter } from 'next/router'
import { useEffect } from 'react'
import { useAuth } from '../stores/AuthStore'

export const DataProvider = ({ children }) => {
  const { currentAuth, setIsLogin, setUser, setLoading } = useAuth()

  const loadUser = async () => {
    const authToken = localStorage.getItem('authToken')
    console.log(authToken)
    if (authToken) {
      eDNAAxios.defaults.headers.Authorization = `${authToken}`
      try {
        await setLoading(true)
        const userProfile = await getMe()
        console.log(userProfile)
        console.log('here setIsLogin true')
        await setIsLogin(true)
        await setUser(userProfile)
        console.log(currentAuth)
      } catch (err) {}
    }
    setLoading(false)
  }

  useEffect(() => {
    loadUser()
  }, [])

  return children
}

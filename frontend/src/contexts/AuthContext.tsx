import { message } from 'antd'
import { getMe } from 'api/auth/getMe'
import { login } from 'api/auth/login'
import { eDNAAxios } from 'api/const'
import { LOGIN_PAGE_TEXT } from 'features/login/locale'
import { useRouter } from 'next/router'
import { createContext, useEffect, useState, useContext } from 'react'
import { IUserProfile } from 'types/user'

interface IAuthContext {
  isLogin: boolean
  user: IUserProfile
  loading: boolean
  login: Function
  logout: Function
}

const AuthContext = createContext<IAuthContext>({
  isLogin: false,
  user: null,
  loading: true,
  login: () => {},
  logout: () => {},
})

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState<IUserProfile>(null)
  const [loading, setLoading] = useState<boolean>(true)
  const router = useRouter()

  const loadUser = async () => {
    const authToken = localStorage.getItem('authToken')
    if (authToken) {
      eDNAAxios.defaults.headers.Authorization = `${authToken}`
      try {
        const userProfile = await getMe()
        setUser(userProfile)
      } catch (err) {
        console.log(err)
      }
    } else {
      setUser(null)
    }
    setLoading(false)
  }

  useEffect(() => {
    loadUser()
  }, [])

  const loginUser = async ({
    email,
    password,
  }: {
    email: string
    password: string
  }) => {
    try {
      const { data } = await login({ email, password })
      localStorage.setItem('authToken', data.authToken)
      eDNAAxios.defaults.headers.Authorization = `${data.authToken}`
      await loadUser()
      message.success(LOGIN_PAGE_TEXT[router.locale].loginSuccessMessage)
      router.push('/internal-data')
    } catch (err) {
      message.error(LOGIN_PAGE_TEXT[router.locale].loginFailMessage)
    }
  }

  const logoutUser = async () => {
    // call logout api
    localStorage.removeItem('authToken')
    await loadUser()
    router.push('/')
  }

  return (
    <AuthContext.Provider
      value={{
        isLogin: !!user,
        user,
        loading,
        login: ({ email, password }) => loginUser({ email, password }),
        logout: logoutUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  return useContext(AuthContext)
}

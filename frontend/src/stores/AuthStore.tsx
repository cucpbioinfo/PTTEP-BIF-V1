import { atom, useAtom } from 'jotai'
import { IUserProfile } from 'types/user'

interface IAuthSession {
  isLogin: boolean
  user: IUserProfile
  loading: boolean
  logout: Function
}

const initialAuthSession: IAuthSession = {
  isLogin: false,
  user: {
    userId: '',
    email: '',
  },
  loading: false,
  logout: () => {},
}
const authAtom = atom<IAuthSession>(initialAuthSession)

export const useAuth = () => {
  const [auth, setAuth] = useAtom(authAtom)
  const setIsLogin = (isLogin: boolean) => setAuth({ ...auth, isLogin })
  const setUser = (user: IUserProfile) => setAuth({ ...auth, user })
  const setLoading = (loading: boolean) => setAuth({ ...auth, loading })
  console.log(auth)
  return {
    currentAuth: auth,
    setIsLogin,
    setUser,
    setLoading,
  }
}

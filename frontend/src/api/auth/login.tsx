import { eDNAAxios } from 'api/const'
import { LoginBody } from 'types/login'

export const login = async (loginBody: LoginBody) => {
  const response = await (await eDNAAxios.post('/login', loginBody)).data
  return response
}

import { eDNAAxios } from 'api/const'
import { IUserProfile } from 'types/user'

export const getMe = async (): Promise<IUserProfile> => {
  const response = await (await eDNAAxios.get('/me')).data.data
  return response
}

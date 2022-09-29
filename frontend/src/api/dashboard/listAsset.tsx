import { eDNAAxios } from 'api/const'

export const listAsset = async () => {
  const response = await (await eDNAAxios.get('/asset')).data
  return response
}

import { eDNAAxios } from 'api/const'

export const listKingdom = async () => {
  const response = await (await eDNAAxios.get('/kingdoms')).data
  return response
}

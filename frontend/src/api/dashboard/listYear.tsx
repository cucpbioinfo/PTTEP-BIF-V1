import { eDNAAxios } from 'api/const'

export const listEvennessYear = async () => {
  const response = await (await eDNAAxios.get('/evennessyear')).data
  return response
}

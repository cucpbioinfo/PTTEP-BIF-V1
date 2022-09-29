import { eDNAAxios } from 'api/const'

export const listMajorGroup = async () => {
  const response = await (await eDNAAxios.get('/major-groups')).data
  return response
}

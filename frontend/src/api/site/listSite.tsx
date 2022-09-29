import { eDNAAxios } from 'api/const'
import { Site } from 'types/site'

export const listSite = async (): Promise<Site[]> => {
  const res = await (await eDNAAxios.get('/sites')).data.data
  return res
}

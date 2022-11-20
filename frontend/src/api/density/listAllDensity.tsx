import { eDNAAxios } from 'api/const'
import { ListDenQuery } from 'types/species'

export const listAllDensity = async (params: ListDenQuery) => {
  const { data } = await (
    await eDNAAxios.get('/alldensity', {
      params: { ...params },
    })
  ).data
  return { data }
}

import { eDNAAxios } from 'api/const'
import { ListADenQuery } from 'types/species'

export const listADen = async (params: ListADenQuery) => {
  const { data } = await (
    await eDNAAxios.get('/assetdensity', {
      params: { ...params },
    })
  ).data
  return { data }
}

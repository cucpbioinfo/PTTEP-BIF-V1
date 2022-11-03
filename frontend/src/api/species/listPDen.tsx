import { eDNAAxios } from 'api/const'
import { ListPDenQuery } from 'types/species'

export const listPDen = async (params: ListPDenQuery) => {
  const { data } = await (
    await eDNAAxios.get('/platformdensity', {
      params: { ...params },
    })
  ).data
  return { data }
}

import { eDNAAxios } from 'api/const'
import { ListDenQuery } from 'types/species'

export const listDen = async (params: ListDenQuery) => {
  const { data } = await (
    await eDNAAxios.get('/density', {
      params: { ...params },
    })
  ).data
  return { data }
}

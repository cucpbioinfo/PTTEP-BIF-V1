import { eDNAAxios } from 'api/const'
import { ListSumQuery } from 'types/species'

export const listSum = async (params: ListSumQuery) => {
  const { data } = await (
    await eDNAAxios.get('/assetsummary', {
      params: { ...params },
    })
  ).data
  return { data }
}

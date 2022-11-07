import { eDNAAxios } from 'api/const'
import { ListSumQuery } from 'types/species'

export const listAllSummary = async (params: ListSumQuery) => {
  const { data } = await (
    await eDNAAxios.get('/allsummary', {
      params: { ...params },
    })
  ).data
  return { data }
}

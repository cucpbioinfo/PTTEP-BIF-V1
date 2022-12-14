import { eDNAAxios } from 'api/const'
import { ListSumStationQuery } from 'types/species'

export const listAllSummary = async (params: ListSumStationQuery) => {
  const { data } = await (
    await eDNAAxios.get('/allsummary', {
      params: { ...params },
    })
  ).data
  return { data }
}

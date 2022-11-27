import { eDNAAxios } from 'api/const'
import { ListSumStationQuery } from 'types/species'

export const listSumStation = async (params: ListSumStationQuery) => {
  const { data } = await (
    await eDNAAxios.get('/summary', {
      params: { ...params },
    })
  ).data
  return { data }
}

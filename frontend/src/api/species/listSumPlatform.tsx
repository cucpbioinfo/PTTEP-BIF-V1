import { eDNAAxios } from 'api/const'
import { ListSumPlatformQuery } from 'types/species'

export const listSumPlatform = async (params: ListSumPlatformQuery) => {
  const { data } = await (
    await eDNAAxios.get('/platformsummary', {
      params: { ...params },
    })
  ).data
  return { data }
}

import { eDNAAxios } from 'api/const'
import { ListSumAssetQuery } from 'types/species'

export const listSumAsset = async (params: ListSumAssetQuery) => {
  const { data } = await (
    await eDNAAxios.get('/assetsummary', {
      params: { ...params },
    })
  ).data
  return { data }
}

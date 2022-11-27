import { eDNAAxios } from 'api/const'

export const listYearAsset = async (assetId?: string) => {
  const response = await (
    await eDNAAxios.get('/yearassetdensity', {
      params: {assetId},
    })
  ).data
  return response
}

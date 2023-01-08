import { eDNAAxios } from 'api/const'

export const listRefPlatform = async (assetId?: string) => {
  const response = await (
    await eDNAAxios.get('/refplatform', {
      params: { assetId },
    })
  ).data
  return response
}

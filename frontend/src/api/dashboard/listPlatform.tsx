import { eDNAAxios } from 'api/const'

export const listPlatform = async (assetId?: string) => {
  const response = await (
    await eDNAAxios.get('/platform', {
      params: { assetId },
    })
  ).data
  return response
}

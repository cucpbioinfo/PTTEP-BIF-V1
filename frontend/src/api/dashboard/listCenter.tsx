import { eDNAAxios } from 'api/const'

export const listCenter = async (assetId?: string) => {
  const response = await (
    await eDNAAxios.get('/centerlocation', {
      params: { assetId },
    })
  ).data
  return response
}

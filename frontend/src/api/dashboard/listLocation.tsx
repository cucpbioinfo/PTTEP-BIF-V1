import { eDNAAxios } from 'api/const'

export const listLocation = async (assetId?: string) => {
  const response = await (
    await eDNAAxios.get('/assetlocation', {
      params: { assetId },
    })
  ).data
  return response
}

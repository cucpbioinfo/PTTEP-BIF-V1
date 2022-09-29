import { eDNAAxios } from 'api/const'

export const listFamily = async (orderId?: string) => {
  const response = await (
    await eDNAAxios.get('/families', {
      params: { orderId },
    })
  ).data
  return response
}

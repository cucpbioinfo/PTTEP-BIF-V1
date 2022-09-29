import { eDNAAxios } from 'api/const'

export const listDiversity = async (stationId?: string) => {
  const response = await (
    await eDNAAxios.get('/diversity', {
      params: { stationId },
    })
  ).data
  return response
}

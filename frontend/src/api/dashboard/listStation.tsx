import { eDNAAxios } from 'api/const'

export const listStation = async (platformId?: string) => {
  const response = await (
    await eDNAAxios.get('/station', {
      params: { platformId },
    })
  ).data
  return response
}

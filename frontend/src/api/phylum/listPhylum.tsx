import { eDNAAxios } from 'api/const'

export const listPhylum = async (kingdomId?: string) => {
  const response = await (
    await eDNAAxios.get('/phylums', {
      params: { kingdomId },
    })
  ).data
  return response
}

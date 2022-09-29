import { eDNAAxios } from 'api/const'

export const listGenus = async (familyId?: string) => {
  const response = await (
    await eDNAAxios.get('/genus', {
      params: { familyId },
    })
  ).data
  return response
}

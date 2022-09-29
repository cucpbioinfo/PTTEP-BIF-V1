import { eDNAAxios } from 'api/const'

export const listOrder = async (classId?: string) => {
  const response = await (
    await eDNAAxios.get('/orders', {
      params: { classId },
    })
  ).data
  return response
}

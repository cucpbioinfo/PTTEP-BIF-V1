import { eDNAAxios } from 'api/const'

export const listClass = async (phylumId?: string) => {
  const response = await (
    await eDNAAxios.get('/classes', {
      params: { phylumId },
    })
  ).data
  return response
}

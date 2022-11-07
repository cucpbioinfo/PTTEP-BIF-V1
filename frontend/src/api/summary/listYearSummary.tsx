import { eDNAAxios } from 'api/const'

export const listYearSummary = async (platformId?: string) => {
  const response = await (
    await eDNAAxios.get('/yearsummary', {
      params: { platformId },
    })
  ).data
  return response
}

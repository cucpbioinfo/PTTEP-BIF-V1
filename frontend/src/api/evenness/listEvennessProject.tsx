import { eDNAAxios } from 'api/const'

export const listEvennessProject = async (year?: string) => {
  const response = await (
    await eDNAAxios.get('/evenness/project', {
      params: { year },
    })
  ).data
  return response
}

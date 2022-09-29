import { eDNAAxios } from 'api/const'
import { ListDensityQuery } from 'types/density'

export const listDensities = async (params: ListDensityQuery) => {
  const { data, meta } = await (
    await eDNAAxios.get('/densities', {
      params: { ...params },
    })
  ).data
  return { data, totalDensities: meta.total }
}

import { eDNAAxios } from 'api/const'
import { ListSpeciesQuery } from 'types/species'

export const listSpecies = async (params: ListSpeciesQuery) => {
  const { data, meta } = await (
    await eDNAAxios.get('/species', {
      params: { ...params },
    })
  ).data
  return { data, totalSpecies: meta.total }
}

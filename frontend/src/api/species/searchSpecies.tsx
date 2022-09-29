import { eDNAAxios } from 'api/const'
import { SearchSpeciesQuery } from 'types/species'

export const searchSpecies = async (params: SearchSpeciesQuery) => {
  const response = await (
    await eDNAAxios.get('/species/search', {
      params: { ...params, limit: 5 },
    })
  ).data.data
  return response
}

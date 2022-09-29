import { eDNAAxios } from 'api/const'
import { SpeciesDetailsResponse } from 'types/species'

export const getSpeciesDetails = async (
  speciesId: string,
): Promise<SpeciesDetailsResponse> => {
  const res = await (
    await eDNAAxios.get(`/species/${speciesId}/details`)
  ).data.data
  return res
}

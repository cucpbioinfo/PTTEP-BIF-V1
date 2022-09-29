import { eDNAAxios } from 'api/const'

export const listAllSpecies = async (SpeciesId?: string) => {
  const response = await (
    await eDNAAxios.get('/allspecies', {
      params: { SpeciesId },
    })
  ).data
  return response
}

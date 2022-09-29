import { eDNAAxios } from 'api/const'
import { SpeciesCount } from 'types/species'

export const getSpeciesCounts = async (): Promise<SpeciesCount[]> => {
  const res: SpeciesCount[] = await (
    await eDNAAxios.get('/species/count')
  ).data.data
  return res
}

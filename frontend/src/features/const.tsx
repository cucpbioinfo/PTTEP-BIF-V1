import { SpeciesCount, SpeciesType } from 'types/species'

export const DefaultSpeciesCounts: SpeciesCount[] = [
  { speciesType: SpeciesType.Phytoplankton, count: 0 },
  { speciesType: SpeciesType.Zooplankton, count: 0 },
  { speciesType: SpeciesType.Nekton, count: 0 },
  { speciesType: SpeciesType.Benthos, count: 0 },
  { speciesType: SpeciesType.Bacteria, count: 0 },
]

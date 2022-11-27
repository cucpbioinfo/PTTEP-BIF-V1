import { PaginationQuery } from './pagination'

export enum SpeciesType {
  Phytoplankton = 'phytoplankton',
  Zooplankton = 'zooplankton',
  Nekton = 'nekton',
  Benthos = 'benthos',
  Bacteria = 'bacteria',
}

export class ListSpeciesQuery extends PaginationQuery {
  majorGroupId?: string
  kingdomId?: string
  phylumId?: string
  classId?: string
  orderId?: string
  familyId?: string
  genusId?: string
  keyword?: string
  pageNumber: number
  pageSize: number
}

export class ListADenQuery {
  speciesId?: string
  assetId?: string
  year?: string
  // pageNumber: number
  // pageSize: number
}

export class ListPDenQuery {
  speciesId?: string
  assetId?: string
  platformId?: string
  year?: string
  // pageNumber: number
  // pageSize: number
}

export class ListDenQuery {
  speciesId?: string
  assetId?: string
  platformId?: string
  stationId?: string
  year?: string
  // pageNumber: number
  // pageSize: number
}

export class ListSumAssetQuery {
  majorGroupId?: string
  assetId?: string
  platformId?: string
  stationId?: string
  year?: string
  // pageNumber: number
  // pageSize: number
}

export class ListSumPlatformQuery {
  majorGroupId?: string
  assetId?: string
  platformId?: string
  stationId?: string
  year?: string
  // pageNumber: number
  // pageSize: number
}

export class ListSumStationQuery {
  majorGroupId?: string
  assetId?: string
  platformId?: string
  stationId?: string
  year?: string
  // pageNumber: number
  // pageSize: number
}

export class SpeciesListMeta {
  pageNumber?: number
  pageSize?: number
  total?: number
}

export class ADenListMeta {
  pageNumber?: number
  pageSize?: number
  total?: number
}

export class SpeciesCount {
  speciesType: SpeciesType
  count: number
}

export class SearchSpeciesQuery {
  keyword: string
  limit: number
}

export class SpeciesDetailsResponse {
  speciesId: string
  speciesName: string
  majorGroupId : string
  majorGroupName: string
  kingdomName: string
  phylumName: string
  className: string
  orderName: string
  familyName: string
  genusName: string
  createAt: Date
  updatedAt: Date
  kingdomId: string //26June, Quicklink
  phylumId: string
  classId: string
  orderId: string
  familyId: string
  genusId: string
  identificationName: string //29Sep
  methodName: string
}

export class ADenDetailsResponse {
  densityId: string
  year: string
  assetId : string
  assetName: string
  speciesId: string
  speciesName: string
  surface: string
  euphotic_zone: string
}
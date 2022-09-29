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

export class SpeciesListMeta {
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

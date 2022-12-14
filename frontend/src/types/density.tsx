import { PaginationQuery } from './pagination'

export class ListDensityQuery extends PaginationQuery {
  // majorGroupId?: string
  // kingdomId?: string
  // phylumId?: string
  // classId?: string
  // orderId?: string
  // familyId?: string
  // genusId?: string
  speciesId?: string
  keyword?: string
  pageNumber: number
  pageSize: number
  assetId?: string
  platformId?: string
  stationId?: string
}

export class DensityListMeta {
  pageNumber?: number
  pageSize?: number
  total?: number
}

// export class DensityCount {
//   speciesType: SpeciesType
//   count: number
// }

export class SearchDensityQuery {
  keyword: string
  limit: number
}

export class DensityDetailsResponse {
  densityId: string
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
}


export class ListAssetDensityQuery {
  assetId?: string
  year?: string
}

export class AssetDensityDetailsResponse {
  densityId: string
  year: string
  assetId: string
  assetName: string
  speciesId: string
  speciesName: string
  surface: string
  euphotic_zone: string
}
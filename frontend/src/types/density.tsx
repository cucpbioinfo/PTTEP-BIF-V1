import { PaginationQuery } from './pagination'

export class ListDensityQuery extends PaginationQuery {
  majorGroupId?: string
  kingdomId?: string
  phylumId?: string
  classId?: string
  orderId?: string
  familyId?: string
  genusId?: string
  speciesId?: string
  keyword?: string
  pageNumber: number
  pageSize: number
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

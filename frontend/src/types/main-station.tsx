import { FieldLocation } from './field-location'
import { Location } from './location'

export class MainStation {
  mainStationId: string
  mainStationName: string
  radiusFromSite: string
  easting: number
  northing: number
  location: Location
  fieldLocations: Array<FieldLocation>
  createdAt: Date
  updatedAt: Date
}

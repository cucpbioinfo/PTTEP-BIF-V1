import { MainStation } from './main-station'

export class Site {
  siteId: string
  siteName: string
  mainStations: Array<MainStation>
  createdAt: Date
  updatedAt: Date
}

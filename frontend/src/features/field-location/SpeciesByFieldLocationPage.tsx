import { Form, Select } from 'antd/'
import { listMainStation } from 'api/main-station/listMainStation'
import { listSite } from 'api/site/listSite'
import { PageLayout } from 'components/PageLayout'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { MainStation } from 'types/main-station'
import { Site } from 'types/site'
import { SpeciesType } from 'types/species'
import { SPECIES_BY_FIELD_LOCATION_TEXT } from './locale'
import { SpeciesByFieldLocationTable } from './SpeciesByFieldLocationTable'

const { Option } = Select

export const SpeciesByFieldLocationPage = () => {
  const { locale } = useRouter()
  const [sites, setSites] = useState<Site[]>([])
  const [mainStations, setMainStations] = useState<MainStation[]>([])

  const fetchSites = async () => {
    const data = await listSite()
    console.log(data)
    setSites(data)
  }
  const fetchMainStations = async () => {
    const data = await listMainStation()
    console.log(data)
    setMainStations(data)
  }

  useEffect(() => {
    fetchSites()
    fetchMainStations()
  }, [])

  return (
    <PageLayout>
      <h1 className="text-2xl mb-8 font-bold">
        {SPECIES_BY_FIELD_LOCATION_TEXT[locale].pageTitle}
      </h1>
      <div className="grid md:grid-cols-6 md:gap-6 grid-cols-2 gap-4">
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_BY_FIELD_LOCATION_TEXT[locale].selectSiteLabel}
          </div>
          <Select placeholder="site">
            {sites.map((site: Site) => (
              <Option key={site.siteId} value={site.siteId}>
                {site.siteName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_BY_FIELD_LOCATION_TEXT[locale].selectStationLabel}
          </div>
          <Select placeholder="station">
            {mainStations.map((site: MainStation) => (
              <Option key={site.mainStationId} value={site.mainStationId}>
                {site.mainStationName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_BY_FIELD_LOCATION_TEXT[locale].selectSpeciesTypeLabel}
          </div>
          <Select placeholder="species type">
            {Object.keys(SpeciesType).map((k: string) => (
              <Option key={SpeciesType[k]} value={SpeciesType[k]}>
                {SpeciesType[k]}
              </Option>
            ))}
          </Select>
        </div>
      </div>
      <div className="w-full mt-8">
        <SpeciesByFieldLocationTable />
      </div>
    </PageLayout>
  )
}

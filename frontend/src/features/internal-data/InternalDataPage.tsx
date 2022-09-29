import { Divider, Select } from 'antd'
import { listMainStation } from 'api/main-station/listMainStation'
import { listMajorGroup } from 'api/major-group/listMajorGroup'
import { listSite } from 'api/site/listSite'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { MainStation } from 'types/main-station'
import { Site } from 'types/site'
import { MOCK_INTERNAL_DATA } from 'utils/mockData'
import { InternalDataItem } from './InternalDataItem'
import { INTERNAL_DATA_FILTER_TEXT } from './locale'

const { Option } = Select

export const InternalDataPage = () => {
  const router = useRouter()
  const { locale } = router
  const [sites, setSites] = useState<Site[]>([])
  const [mainStations, setMainStations] = useState<MainStation[]>([])
  const [years, setYears] = useState<Array<number>>([])
  const [majorGroups, setMajorGroups] = useState([])

  const onFilterChange = (key: string, value?: string) => {
    const url = {
      pathname: router.pathname,
      query: {
        ...router.query,
        [key]: value,
      },
    }
    router.push(url, url, { locale })
  }

  const fetchSites = async () => {
    const data = await listSite()
    setSites(data)
  }

  const fetchMainStations = async (siteId?: string) => {
    const data = await listMainStation(siteId)
    setMainStations(data)
  }

  const fetchYears = async () => {
    const now = new Date().getFullYear()
    setYears([now, now - 1, now - 2, now - 3])
  }

  const fetchMajorGroups = async () => {
    const { data } = await listMajorGroup()
    setMajorGroups(data)
  }

  useEffect(() => {
    fetchSites()
    fetchYears()
    fetchMajorGroups()
  }, [])

  useEffect(() => {
    fetchMainStations(router.query.siteId as string)
  }, [router.query.siteId])

  return (
    <>
      <div className="grid md:grid-cols-6 grid-cols-2 md:gap-8 gap-4">
        <div className="w-full">
          <div className="text-sm">
            {INTERNAL_DATA_FILTER_TEXT[router.locale].selectSiteLabel}
          </div>
          <Select
            placeholder={
              INTERNAL_DATA_FILTER_TEXT[router.locale].selectSitePlaceholder
            }
            onChange={(value: string | undefined) => {
              onFilterChange('siteId', value)
            }}
          >
            {sites.map((site) => (
              <Option value={site?.siteId} key={site?.siteId}>
                {site?.siteName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {INTERNAL_DATA_FILTER_TEXT[router.locale].selectStationLabel}
          </div>
          <Select
            placeholder={
              INTERNAL_DATA_FILTER_TEXT[router.locale].selectStationPlaceholder
            }
            onChange={(value: string | undefined) => {
              onFilterChange('mainStationId', value)
            }}
          >
            {mainStations.map((mainStation) => (
              <Option
                value={mainStation?.mainStationId}
                key={mainStation?.mainStationId}
              >
                {mainStation?.mainStationName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {INTERNAL_DATA_FILTER_TEXT[router.locale].selectYearLabel}
          </div>
          <Select
            placeholder={
              INTERNAL_DATA_FILTER_TEXT[router.locale].selectYearPlaceholder
            }
            onChange={(value: string | undefined) => {
              onFilterChange('year', value)
            }}
          >
            {years.map((year) => (
              <Option value={year} key={year}>
                {year}
              </Option>
            ))}
          </Select>
        </div>
        {/* <div className="w-full col-span-2">
          <div className="text-sm">
            {INTERNAL_DATA_FILTER_TEXT[router.locale].selectMajorGroupLabel}
          </div>
          <Select
            placeholder={
              INTERNAL_DATA_FILTER_TEXT[locale].selectMajorGroupPlaceholder
            }
            onChange={(value: string | undefined) => {
              onFilterChange('majorGroupId', value)
            }}
          >
            {majorGroups.map((majorGroup) => (
              <Option
                value={majorGroup?.majorGroupId}
                key={majorGroup?.majorGroupId}
              >
                {majorGroup?.majorGroupName}
              </Option>
            ))}
          </Select>
        </div> */}
      </div>
      <Divider />
      <div className="grid grid-cols-1 gap-y-8 px-8 md:px-16 mt-8 md:mt-16">
        {MOCK_INTERNAL_DATA.map((data) => (
          <InternalDataItem
            title={data.title}
            href={data.href}
            updatedAt={data.updatedAt}
            key={data.title}
          />
        ))}
      </div>
      <Divider />

    </>
  )
}

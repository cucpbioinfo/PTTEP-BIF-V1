import { Select } from 'antd'
import { listYearStation } from 'api/species/listYearStation'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const { Option } = Select
export const SummaryFilterYear = () => {
  const router = useRouter()
  const { locale } = router
  const [stationyear, setStationYear] = useState([])
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
  const fetchStationYear = async (stationId?: string) => {
      const { data } = await listYearStation(stationId)
      setStationYear(data)
    }

  useEffect(() => {
    fetchStationYear(router.query.stationId as string)
  }, [router.query.stationId])

  return (
    <>
      <div className="text-sm">
        Year
      </div>
      <div className="md:w-1/6 w-full">
        <Select
            placeholder="year"
            onChange={(value: string | undefined) => {
              onFilterChange('year', value)
            }}
          >
            {stationyear.map((year) => (
              <Option value={year?.year} key={year?.year}>
                {year?.year}
              </Option>
            ))}
        </Select>
      </div>
    </>
  )
}

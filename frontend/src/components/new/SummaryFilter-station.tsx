import { Select } from 'antd'
import { listStation } from 'api/dashboard/listStation'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const { Option } = Select
export const SummaryFilterStation = () => {
  const router = useRouter()
  const { locale } = router

  const [station, setStation] = useState([])

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
  
  const fetchStation = async (platformId?: string) => {
    const { data } = await listStation(platformId)
    setStation(data)
  }
 
  useEffect(() => {
    fetchStation(router.query.platformId as string)
  }, [router.query.platformId])
  

  return (
    <>
      {/* <div className="text-sm">
        Station
      </div> */}
      <div className="md:w-1/6 w-full mr-2">
        <Select
          placeholder="Station"
          onChange={(value: string | undefined) => {
            onFilterChange('stationId', value)
          }}
        >
          {station.map((station) => (
            <Option
              value={station?.stationId}
              key={station?.stationId}
            >
              {station?.stationName.toUpperCase()}
            </Option>
          ))}
        </Select>
      </div>
    </>
  )
}

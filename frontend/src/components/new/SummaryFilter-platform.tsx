import { Select } from 'antd'
import { listPlatform } from 'api/dashboard/listPlatform'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const { Option } = Select
export const SummaryFilterPlatform = () => {
  const router = useRouter()
  const { locale } = router

  const [platform, setPlatform] = useState([])

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
 
  const fetchPlatform = async (assetId?: string) => {
    const { data } = await listPlatform(assetId)
    setPlatform(data)
  }
  
  useEffect(() => {
    fetchPlatform(router.query.assetId as string)
  }, [router.query.assetId])
  // useEffect(() => {
  //   fetchPlatformYear(router.query.platformId as string)
  // }, [router.query.platformId])
  
  return (
    <>
      {/* <div className="text-sm">
        Platform
      </div> */}
      <div className="md:w-1/6 w-full mr-2">
        <Select
          placeholder="Platform"
          onChange={(value: string | undefined) => {
            onFilterChange('platformId', value)
          }}
        >
          {platform.map((platform) => (
            <Option
              value={platform?.platformId}
              key={platform?.platformId}
            >
              {platform?.platformName.toUpperCase()}
            </Option>
          ))}
        </Select>
      </div>
    </>
  )
}

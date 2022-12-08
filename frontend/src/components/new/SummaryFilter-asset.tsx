import { Select, Button } from 'antd'
import { listAsset } from 'api/dashboard/listAsset'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const { Option } = Select
export const SummaryFilterAsset = () => {
  const router = useRouter()
  const { locale } = router

  const [asset, setAsset] = useState([])
  
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

  const fetchAsset = async () => {
    const { data } = await listAsset()
    setAsset(data)
  }
  
  useEffect(() => {
    fetchAsset()
  }, [])

  return (
    <>
      <div className="md:w-1/3 w-full mr-2">
        <Select
          placeholder="Asset"
          onChange={(value: string | undefined) => {
            onFilterChange('assetId', value)
          }}
        >
          {asset.map((asset) => (
            <Option
              value={asset?.assetId}
              key={asset?.assetId}
            >
              {asset?.assetName.toUpperCase()}
            </Option>
          ))}
        </Select>
      </div>
    </>
  )
}

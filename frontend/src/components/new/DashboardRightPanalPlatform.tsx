import { Select } from 'antd'
import { listPlatform } from 'api/dashboard/listPlatform'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const DashboardRightPanalPlatform = () => {
  const router = useRouter()

  const [platform, setPlatform] = useState([])

  
  const fetchPlatform = async (assetId?: string) => {
    const { data } = await listPlatform(assetId)
    setPlatform(data)
  }
 
  useEffect(() => {
    fetchPlatform(router.query.assetId as string)
  }, [router.query.assetId])
  

  return (
    <>
      <div className="font-bold">Platform</div>
      <div className="w-full p-1 rounded-md">
        {platform.map((platform) => (
          <div className="flex">
            <div className="w-3/4 m-auto bg-sky-500 hover:bg-sky-700">
              <a onClick={() => router.push(`/?platformId=${platform?.platformId}`)}>
                <div className="font-bold">{platform?.platformName.toUpperCase()}</div>
              </a>
            </div>
          </div >
        ))}
      </div>
    </>
  )
}

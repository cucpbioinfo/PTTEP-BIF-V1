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
      <div className="font-bold text-base">Location</div>
      {/* <div className="w-full p-1 rounded-md">
        {platform.map((platform) => (
          <div className="flex">
            <div className="w-3/4 m-auto bg-sky-500 hover:bg-sky-700">
              <a onClick={() => router.push(`/?platformId=${platform?.platformId}`)}>
                <div className="font-bold">{platform?.platformName.toUpperCase()}</div>
              </a>
            </div>
          </div >
        ))}
      </div> */}
    
    
      <ul role="list" className="divide-y divide-slate-200 list-none hover:list-disc ">

      {platform.map((platform) => (
      <div className="border-1">
      <li className="flex h-auto w-auto m-auto border-1 box-border items-center justify-between">
        <div className="flex box-border items-start w-full m-auto p-auto">
          <div className="m-auto p-auto w-full box-border">
            <h4 className="m-auto p-auto box-border outline-none">
              <a onClick={() => router.push(`/?platformId=${platform?.platformId}&view=platform`)}>
                <div className="hover:font-bold">{platform?.platformName.toUpperCase()}</div>
              </a>
            </h4>
            {/* <div className="m-auto p-auto box-border outline-none">celestin.louis@example.com</div> */}
          </div>
        </div>
      </li>
      </div>

      ))}
            
    </ul>
    </>
  )
}

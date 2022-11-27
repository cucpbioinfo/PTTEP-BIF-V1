import { Select } from 'antd'
import { listStation } from 'api/dashboard/listStation'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const DashboardRightPanalStation = () => {
  const router = useRouter()

  const [station, setStation] = useState([])

  
  const fetchStation = async (platformId?: string) => {
    const { data } = await listStation(platformId)
    setStation(data)
  }
 
  useEffect(() => {
    fetchStation(router.query.platformId as string)
  }, [router.query.platformId])
  

  return (
    <>
      <div className="font-bold">Station</div>
       <div className="w-full p-1 rounded-md">
         {/* <div className="flex"></div>
         <div className="w-3/4 m-auto">
           <a onClick={() => router.push(`/?assetId=${router.query.assetId}&platformId=${router.query.platformId}&stationId=`)}>
             <div className="font-bold">All</div>
           </a>
         </div> */}
         {station.map((station) => (
        <div className="flex">
          <div className="w-3/4 m-auto">
            <a onClick={() => router.push(`/?assetId=${router.query.assetId}&platformId=${router.query.platformId}&stationId=${station?.stationId}`)}>
              <div className="font-bold">{station?.stationName.toUpperCase()}</div>
            </a>
          </div>
        </div >
        ))}
      </div>    
    </>
  )
}

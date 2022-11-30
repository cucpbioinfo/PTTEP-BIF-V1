import { Divider, Empty ,Button } from 'antd'
import { listDen } from 'api/species/listDen'
import { DenBar } from 'features/echart/DenBar'
import { DensityFilter } from 'components/new/DensityFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const DensityBarFilter = (specId) => {
  const router = useRouter()
  const [den, setDen] = useState([])
  let sid = ""
  const fetchDen = async () => {
    const {
      speciesId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listDen({
      speciesId: speciesId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setDen(data)
  }

  useEffect(() => {
    fetchDen()
  }, [router.query])

  den.map(({ speciesId }) => (
    sid = speciesId
  ))

  return (
    <div>
      
      {/* <DensityFilter/> */}
      
      <Divider />
      {/* <a href="/density?speciesId=e1dced3d-653b-4f83-9ae8-b341c4d85523">View all</a> */}

      {/* <Button onClick={() => router.push(`/density?speciesId=${sid}`)}>See all</Button> */}

      <div>{den.length === 0 && <Empty />}</div>
      <div>
        {den.length !== 0 &&
          den.map(({ densityId ,speciesName,stationName, year, surface ,euphotic_zone }) => (
            <div>
              {/* <div>{densityId}</div>
              <div>{speciesName}</div>
              <div>{stationName}</div>
              <div>{year}</div>
              <div>{surface}</div>
              <div>{euphotic_zone}</div> */}
            <DenBar densityId={densityId} speciesName={speciesName} name={stationName} year={year} surface={surface} zone={euphotic_zone} />
            </div>
          ))}
      </div>
  
      </div>
  )
}

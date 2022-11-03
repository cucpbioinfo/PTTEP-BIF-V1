import { Divider, Empty } from 'antd'
import { listDen } from 'api/species/listDen'
import { DenBar } from 'features/echart/DenBar'
import { DensityFilter } from 'components/new/DensityFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const DensityBarFilter = (specId) => {
  const router = useRouter()
  const [den, setDen] = useState([])
  
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

  return (
    <div>
      
      <DensityFilter/>
      
      <Divider />

      <div>{den.length === 0 && <Empty />}</div>
      <div>
        {den.length !== 0 &&
          den.map(({ densityId ,speciesName,stationName, year, surface ,euphotic_zone }) => (
            <div>
            <DenBar densityId={densityId} speciesName={speciesName} name={stationName} year={year} surface={surface} zone={euphotic_zone} />
            </div>
          ))}
      </div>
  
      </div>
  )
}
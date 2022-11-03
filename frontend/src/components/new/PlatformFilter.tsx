import { Divider, Empty } from 'antd'
import { listPDen } from 'api/species/listPDen'
import { DenBar } from 'features/echart/DenBar'
import { PlatformDensityFilter } from 'components/new/PlatformDensityFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const PlatformDensityBarFilter = (specId) => {
  const router = useRouter()
  const [pden, setPDen] = useState([])
  
  const fetchPDen = async () => {
    const {
      speciesId,
      assetId,
      platformId,
      year,
      
    } = router.query
    const { data } = await listPDen({
      speciesId: speciesId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      year: year as string,
      
    })
    setPDen(data)
  }

  useEffect(() => {
    fetchPDen()
  }, [router.query])

  return (
    <div>
      
      <PlatformDensityFilter/>
      
      <Divider />

      <div>{pden.length === 0 && <Empty />}</div>
      <div>
        {pden.length !== 0 &&
          pden.map(({ densityId ,speciesName,platformName, year, surface ,euphotic_zone }) => (
            <div>
            <DenBar densityId={densityId} speciesName={speciesName} name={platformName} year={year} surface={surface} zone={euphotic_zone} />
            </div>
          ))}
      </div>
  
      </div>
  )
}

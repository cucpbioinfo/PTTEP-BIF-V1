import { Divider, Empty, Pagination, Select } from 'antd'
import { listADen } from 'api/species/listADen'
import { DenBar } from 'features/echart/DenBar'
import { AssetDensityFilter } from 'components/new/AssetDensityFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'


export const AssetDensityBarFilter = (specId) => {
  const router = useRouter()
  const [aden, setADen] = useState([])
  
  const fetchADen = async () => {
    const {
      speciesId,
      assetId,
      year,
      
    } = router.query
    const { data } = await listADen({
      speciesId: speciesId as string,
      assetId: assetId as string,
      year: year as string,
      
    })
    setADen(data)
  }

  useEffect(() => {
    fetchADen()
  }, [router.query])

  return (
    <div>
      
      <AssetDensityFilter/>
      
      <Divider />
      
      <div>{aden.length === 0 && <Empty />}</div>
      <div>
        {aden.length !== 0 &&
          aden.map(({ densityId ,speciesId,speciesName,assetName, year, surface ,euphotic_zone }) => (
            <div>
            <DenBar densityId={densityId} speciesId={speciesId} speciesName={speciesName} assetName={assetName} year={year} surface={surface} zone={euphotic_zone} />
            </div>
          ))}
      </div>
      
      </div>
  )
}

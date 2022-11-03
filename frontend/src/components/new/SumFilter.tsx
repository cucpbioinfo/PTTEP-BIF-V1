import { Divider, Empty } from 'antd'
import { listSum } from 'api/species/listSum'
import { DenBar } from 'features/echart/DenBar'
import { SummaryFilter } from 'components/new/SummaryFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const SummaryBarFilter = (specId) => {
  const router = useRouter()
  const [sum, setSum] = useState([])
  
  const fetchSum = async () => {
    const {
      speciesId,
      majorGroupId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listSum({
      speciesId: speciesId as string,
      majorGroupId: majorGroupId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setSum(data)
  }

  useEffect(() => {
    fetchSum()
  }, [router.query])

  return (
    <div>
      
      {/* <SummaryFilter/>
      
      <Divider /> */}

      <div>{sum.length === 0 && <Empty />}</div>
      <div>
        {sum.length !== 0 &&
          sum.map(({ summaryId,year,majorGroupName,assetName,platformName,stationName,surfaceEvenness,euphoticzoneEvenness,surfaceShannon,euphoticzoneShannon,surfaceNumber,euphoticzoneNumber}) => (
            <div>
              {/* <div>id: {summaryId}</div>
              <div>year: {year}</div>
              <div>majorGroupName: {majorGroupName}</div>
              <div>assetName: {assetName}</div>
              <div>platformName: {platformName}</div>
              <div>stationName: {stationName}</div> */}

              <div className="flex space-x-4 mt-6 ml-4 mr-4 mb-8">
                <div className="w-1/3 border p-4 shadow rounded-md">
                  <div className="w-full h-60 mb-1 border text-left">Evenness</div>
                  <div>
                    <DenBar densityId={summaryId} speciesName="Evenness" name={stationName} year={year} surface={surfaceEvenness} zone={euphoticzoneEvenness} />
                  </div>
                </div>
                <div className="w-1/3 border p-4 shadow rounded-md">
                  <div className="w-full h-60 mb-1 border text-left">Diversity</div>
                  <div>
                    <DenBar densityId={summaryId} speciesName="Diversity" name={stationName} year={year} surface={surfaceShannon} zone={euphoticzoneShannon} />
                  </div>
                </div>
                <div className="w-1/3 border p-4 shadow rounded-md">
                  <div className="w-full h-60 mb-1 border text-left">Number</div>
                  <div>
                    <DenBar densityId={summaryId} speciesName="Number Of Species" name={stationName} year={year} surface={surfaceNumber} zone={euphoticzoneNumber} />
                  </div>
                </div>

              </div>

    
            </div>
          ))}
      </div>
  
      </div>
  )
}

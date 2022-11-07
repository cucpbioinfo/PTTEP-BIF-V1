import { PageLayout } from 'components/new/PageLayout'
import { SummaryFilter } from 'components/new/SummaryFilter'
import { SummaryBarFilter } from 'components/new/SumFilter'
import { SummaryBarStationYear } from 'features/echart/summarybarstationyear'
import { AssetDensityBarFilter } from 'components/new/AssetFilter'
//import { SummaryFilterGraph } from 'features/echart/testfilter'
import { useRouter } from 'next/router'
import { SummaryChartStation } from 'components/new/summarychartStation'
import React, { useEffect, useState } from 'react'
const dataimport = [
  {
    stationName: 'station_a',
    year: '2019',
    surfaceShannon: '3.09',
    surfaceNumber: '47',
    surfaceMax: '3.85',
    surfaceEvenness: '0.8',
    euphoticzoneShannon: '2.83',
    euphoticzoneNumber: '53',
    euphoticzoneMax: '3.97',
    euphoticzoneEvenness: '0.71',
  },
  {
    stationName: 'station_a',
    year: '2022',
    surfaceShannon: '3.01',
    surfaceNumber: '47',
    surfaceMax: '3.85',
    surfaceEvenness: '0.7',
    euphoticzoneShannon: '2.73',
    euphoticzoneNumber: '53',
    euphoticzoneMax: '3.97',
    euphoticzoneEvenness: '0.75',
  },
  {
    stationName: 'station_b',
    year: '2019',
    surfaceShannon: '2.09',
    surfaceNumber: '47',
    surfaceMax: '3.85',
    surfaceEvenness: '0.8',
    euphoticzoneShannon: '3.83',
    euphoticzoneNumber: '53',
    euphoticzoneMax: '3.97',
    euphoticzoneEvenness: '0.71',
  },
  {
    stationName: 'station_b',
    year: '2022',
    surfaceShannon: '1.51',
    surfaceNumber: '47',
    surfaceMax: '3.85',
    surfaceEvenness: '0.7',
    euphoticzoneShannon: '1.73',
    euphoticzoneNumber: '53',
    euphoticzoneMax: '3.97',
    euphoticzoneEvenness: '0.75',
  },
] 
const assetbar = () => {
  const router = useRouter()
  return (
    <PageLayout>
      {/* BarGraph
      <AssetDensityBarFilter onClick={() => router.push(`/species/${'a906251d-6916-4a2f-b035-129f6cb3d59f'}`)}/> */}
      <SummaryFilter/>
      {/* <SummaryBarStationYear dataimport={dataimport}/>
      <Testcheckbox/> */}
      <SummaryChartStation/>
    </PageLayout>
  )
}

export default assetbar

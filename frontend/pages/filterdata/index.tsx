import { PageLayout } from 'components/new/PageLayout'
import { SummaryFilter } from 'components/new/SummaryFilter'
import { SummaryBarFilter } from 'components/new/SumFilter'
import { SummaryBarStationYear } from 'features/echart/summarybarstationyear'
import { AssetDensityBarFilter } from 'components/new/AssetFilter'
//import { SummaryFilterGraph } from 'features/echart/testfilter'
import { useRouter } from 'next/router'
import { SummaryChartStation } from 'components/new/summarychartStation'
import React, { useEffect, useState } from 'react'

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

import { Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import { useRouter } from 'next/router'
import { SummaryChartStationYear } from 'components/new/summarychartStationYear'
import { SummaryChartStation } from 'components/new/summarychartStation'
import { SummaryFilterAsset } from 'components/new/SummaryFilter-asset'
import { SummaryFilterPlatform } from 'components/new/SummaryFilter-platform' 
import { SummaryFilterStation } from 'components/new/SummaryFilter-station'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
import { SummaryFilterGroup } from 'components/new/SummaryFilter-group'
import React, { useEffect, useState } from 'react'

const index = () => {

  return (
    <PageLayout>
      <SummaryFilterGroup/>
      <SummaryFilterAsset/>
      <SummaryFilterPlatform/>
      <SummaryFilterStation/>
      {/* <SummaryFilterYear/> */}
      <SummaryChartStationYear />
    </PageLayout>
  )
}

export default index

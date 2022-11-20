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

const { TabPane } = Tabs

const assetbar = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <Tabs defaultActiveKey="1">
        <TabPane tab="SummaryChartStationYear" key="1" >
          <SummaryFilterGroup/>
          <SummaryFilterAsset/>
          <SummaryFilterPlatform/>
          <SummaryFilterStation/>
          {/* <SummaryFilterYear/> */}
          <SummaryChartStationYear />
        </TabPane>
        <TabPane tab="SummaryChartStation" key="2">
          <SummaryFilterGroup/>
          <SummaryFilterAsset/>
          <SummaryFilterPlatform/>
          {/* <SummaryFilterStation/> */}
          <SummaryFilterYear/>
          <SummaryChartStation/>
        </TabPane>
      </Tabs>
    </PageLayout>
  )
}

export default assetbar

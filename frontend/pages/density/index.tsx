import { Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import { useRouter } from 'next/router'
import { DensityStationYear } from 'components/new/densityStationYear'
import { DensityStation } from 'components/new/densityStation'
import { SummaryFilterAsset } from 'components/new/SummaryFilter-asset'
import { SummaryFilterPlatform } from 'components/new/SummaryFilter-platform' 
import { SummaryFilterStation } from 'components/new/SummaryFilter-station'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
import { SummaryFilterGroup } from 'components/new/SummaryFilter-group'
import React, { useEffect, useState } from 'react'

const { TabPane } = Tabs

const density = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <Tabs defaultActiveKey="1">
        <TabPane tab="DensityStation" key="1">
          <SummaryFilterAsset/>
          <SummaryFilterPlatform/>
          {/* <SummaryFilterGroup/> */}
          <SummaryFilterYear/>
          <DensityStation/>
        </TabPane>
        <TabPane tab="DensityStationYear" key="2">
          <SummaryFilterAsset/>
          <SummaryFilterPlatform/>
          <SummaryFilterStation/>
          {/* <SummaryFilterGroup/> */}
          <DensityStationYear/>
        </TabPane>
      </Tabs>
    </PageLayout>
  )
}

export default density

import { Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import { SummaryFilter } from 'components/new/SummaryFilter'
import { SummaryBarFilter } from 'components/new/SumFilter'
import { SummaryBarStationYear } from 'features/echart/summarybarstationyear'
import { AssetDensityBarFilter } from 'components/new/AssetFilter'
//import { SummaryFilterGraph } from 'features/echart/testfilter'
import { useRouter } from 'next/router'
import { SummaryChartStationYear } from 'components/new/summarychartStationYear'
import { SummaryChartStation } from 'components/new/summarychartStation'
import { DensityStationYear } from 'components/new/densityStationYear'
import { DensityStation } from 'components/new/densityStation'
import React, { useEffect, useState } from 'react'

const { TabPane } = Tabs

const assetbar = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <Tabs defaultActiveKey="1">
                <TabPane tab="SummaryChartStationYear" key="1" >
                  <SummaryChartStationYear />
                </TabPane>
                <TabPane tab="SummaryChartStation" key="2">
                  <SummaryChartStation/>
                </TabPane>
                <TabPane tab="DensityStationYear" key="3">
                  <DensityStationYear/>
                </TabPane>
                <TabPane tab="DensityStation" key="4">
                  <DensityStation/>
                </TabPane>
                

              </Tabs>
      
    </PageLayout>
  )
}

export default assetbar

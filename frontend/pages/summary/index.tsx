import { Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import { useRouter } from 'next/router'
import { SummaryChartStationYear } from 'components/new/summarychartStationYear'
import { SummaryChartStation } from 'components/new/summarychartStation'
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
      </Tabs>
    </PageLayout>
  )
}

export default assetbar

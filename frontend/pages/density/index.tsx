import { Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import { useRouter } from 'next/router'
import { DensityStationYear } from 'components/new/densityStationYear'
import { DensityStation } from 'components/new/densityStation'
import React, { useEffect, useState } from 'react'

const { TabPane } = Tabs

const density = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <Tabs defaultActiveKey="1">
        <TabPane tab="DensityStation" key="1">
          <DensityStation/>
        </TabPane>
        <TabPane tab="DensityStationYear" key="2">
          <DensityStationYear/>
        </TabPane>
      </Tabs>
    </PageLayout>
  )
}

export default density

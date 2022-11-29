import { Tabs ,Space,Button,Divider } from 'antd'
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
      <div>
      <Divider orientation="left" plain>
        Sort by
      </Divider>
      <Space direction="vertical">
        <Space wrap>
        <div className="mr-2"><Button href="/summarystation" >Station</Button></div>
        <div className="mr-2"><Button type="primary" href="/summarystationyear" >Year</Button></div>
        </Space>
      </Space>
      <Divider orientation="left" plain/>
      <div>
      <Space direction="vertical">
        <Space wrap>
          <SummaryFilterAsset/>
          <SummaryFilterPlatform/>
          <SummaryFilterStation/>
        </Space>
      </Space>
      </div>
      <div className="mt-4">
        <Space direction="vertical">
          <Space wrap>
            <SummaryFilterGroup/>
          </Space>
        </Space>
      </div>
      <Divider orientation="left" plain>
        Station
      </Divider>
      <SummaryChartStationYear />
    </div>
    </PageLayout>
  )
}

export default index

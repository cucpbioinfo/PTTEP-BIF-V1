import { Divider, Empty,Button,Tooltip,Tabs } from 'antd'
import { listSumAsset } from 'api/species/listSumAsset'
import { listSumPlatform } from 'api/species/listSumPlatform'
import { listSumStation } from 'api/species/listSumStation'
import { SummaryBarDashboardAsset } from 'components/new/SumBarDashboardAsset';
import { SummaryBarDashboardPlatform } from 'components/new/SumBarDashboardPlatform';
import { SummaryBarDashboardStation } from 'components/new/SumBarDashboardStation';
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

export const SummaryBarFilter = () => {
  const router = useRouter()
  const [suma, setSuma] = useState([])
  const [sump, setSump] = useState([])
  const [sums, setSums] = useState([])
  const { TabPane } = Tabs
  const fetchSuma = async () => {
    const {
      majorGroupId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listSumAsset({
      majorGroupId: majorGroupId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setSuma(data)
  }
  const fetchSump = async () => {
    const {
      majorGroupId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listSumPlatform({
      majorGroupId: majorGroupId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setSump(data)
  }
  const fetchSums = async () => {
    const {
      majorGroupId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listSumStation({
      majorGroupId: majorGroupId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setSums(data)
  }

  useEffect(() => {
    fetchSuma()
  }, [router.query])
  useEffect(() => {
    fetchSump()
  }, [router.query])
  useEffect(() => {
    fetchSums()
  }, [router.query])
let DefaultSelected = "1"
  const SepType = () => {
    
    return (
      <>
        <div className="mt-1 ml-4 mr-4 mb-8">
          <Tabs defaultActiveKey={DefaultSelected}>
            <TabPane tab="Asset" key="1">
              <SummaryBarDashboardAsset/>
            </TabPane>
            <TabPane tab="Platform" key="2">
              <SummaryBarDashboardPlatform/>
            </TabPane>
            {/* <TabPane tab="Station" key="3">
              <SummaryBarDashboardStation/>
            </TabPane> */}
          </Tabs>
        </div>
      </>
    )
  }

  return (
    <>
      <SepType/>
    </>
  )
}

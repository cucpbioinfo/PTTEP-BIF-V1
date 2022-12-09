import { Divider, Empty,Button,Tooltip,Tabs } from 'antd'
import { InfoCircleOutlined } from '@ant-design/icons';
import { listSumAsset } from 'api/species/listSumAsset'
import { listSumPlatform } from 'api/species/listSumPlatform'
import { listSumStation } from 'api/species/listSumStation'
import { DenBar } from 'features/echart/DenBar'
import { SumBar } from 'features/echart/SumBar'
import { SummaryFilter } from 'components/new/SummaryFilter'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
import { SummaryFilterGroup } from 'components/new/SummaryFilter-group'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const DiversityInfo = <span>Definition of Shannon-Weiner Species Diversity Index.</span>;
const EvennessInfo = <span>Definition of about Evenness Index.</span>;
const NumberInfo = <span>Definition of about Number of Species Index.</span>;

export const SummaryBarDashboardPlatform = () => {
  const router = useRouter()
  const [sump, setSump] = useState([])
  const { TabPane } = Tabs

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
  
  useEffect(() => {
    fetchSump()
  }, [router.query])

  const SepType = () => {
    
    return (
      <>
      <div>
      <div>{sump.length === 0 && <Empty 
      description={
        <span>
          Please select platform for sorting data.
        </span>
      }
      />}</div>
      <div>
        {sump.length !== 0 &&
          sump.map(({ summaryId,year,majorGroupName,assetName,platformName,stationName,surfaceEvenness,euphoticzoneEvenness,surfaceShannon,euphoticzoneShannon,surfaceNumber,euphoticzoneNumber}) => (
            <div>
              <div className="flex w-full items-center border p-4 shadow rounded-md mb-1">
                <div className="flex w-full items-center divide-x mr-4">
                  
                  <div>
                    <div className="mr-4 ml-2">Asset : {assetName.toUpperCase()}</div>
                  </div>
                  <div>
                    <div className="mr-4 ml-2"><b>Location : {platformName.toUpperCase()}</b></div>
                  </div>
                  
                  <div className="flex">
                    <div className="ml-5"><SummaryFilterGroup/></div>
                    <div className="ml-2"><SummaryFilterYear/></div>
                  </div>
                </div>
                <div>
                  <div className="mr-2"><Button type="primary" href="/summarystation" >View All</Button></div>
                </div>
            </div>
                <div className="flex space-x-4">
                  {/* ------------------------------------------------------------- */}
                  <div className="w-1/3 border p-3 shadow rounded-md">
                    <div className="flex w-full items-center">
                      
                      <div className="flex w-full items-center divide-x mr-4">
                        <div>
                          <div className="mr-4 ml-2"></div>
                        </div>
                      </div>
                      <div>
                        <div className="mr-2">
                          <Tooltip placement="top" title="def"><InfoCircleOutlined/></Tooltip>
                        </div>
                      </div>

                    </div>

                    <div className="w-full text-lg text-left">
                      <b>Shannon-Weiner Species Diversity Index</b>
                    </div>
                    <br/>
                    <div>
                      <SumBar densityId={summaryId} speciesName="Diversity" name={stationName} year={year} surface={surfaceShannon} zone={euphoticzoneShannon} />
                    </div>
                  </div>
                  {/* ------------------------------------------------------------- */}
                  <div className="w-1/3 border p-3 shadow rounded-md">
                    <div className="flex w-full items-center">
                      
                      <div className="flex w-full items-center divide-x mr-4">
                        <div>
                          <div className="mr-4 ml-2"></div>
                        </div>
                      </div>
                      <div>
                        <div className="mr-2">
                          <Tooltip placement="top" title={EvennessInfo}><InfoCircleOutlined/></Tooltip>
                        </div>
                      </div>

                    </div>

                    <div className="w-full text-lg text-left">
                      <b>Evenness Index</b>
                    </div>
                    <br/>
                    <div>
                      <SumBar densityId={summaryId} speciesName="Evenness" name={stationName} year={year} surface={surfaceEvenness} zone={euphoticzoneEvenness} />
                    </div>
                  </div>
                  {/* ------------------------------------------------------------- */}
                  <div className="w-1/3 border p-3 shadow rounded-md">
                    <div className="flex w-full items-center">
                      
                      <div className="flex w-full items-center divide-x mr-4">
                        <div>
                          <div className="mr-4 ml-2"></div>
                        </div>
                      </div>
                      <div>
                        <div className="mr-2">
                          <Tooltip placement="top" title={NumberInfo}><InfoCircleOutlined/></Tooltip>
                        </div>
                      </div>

                    </div>

                    <div className="w-full text-lg text-left">
                      <b>Number of Species</b>
                    </div>
                    <br/>
                    <div>
                      <SumBar densityId={summaryId} speciesName="No." name={stationName} year={year} surface={surfaceNumber} zone={euphoticzoneNumber} />
                    </div>
                  </div>
                  {/* ------------------------------------------------------------- */}
                  

                  {/* <div className="w-1/3 border p-4 shadow rounded-md">
                    <div className="w-full h-60 mb-1 border text-left">
                      Evenness Index <Tooltip placement="top" title={EvennessInfo}><InfoCircleOutlined/></Tooltip>
                    </div>
                    <div>
                      <SumBar densityId={summaryId} speciesName="Evenness" name={stationName} year={year} surface={surfaceEvenness} zone={euphoticzoneEvenness} />
                    </div>
                  </div> */}


                  {/* <div className="w-1/3 border p-4 shadow rounded-md">
                    <div className="w-full h-60 mb-1 border text-left">
                      Number of Species <Tooltip placement="top" title={NumberInfo}><InfoCircleOutlined/></Tooltip>
                    </div>
                    <div>
                      <SumBar densityId={summaryId} speciesName="No." name={stationName} year={year} surface={surfaceNumber} zone={euphoticzoneNumber} />
                    </div>
                  </div> */}
                  
                </div>
          </div>

          ))}
      </div>
  
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

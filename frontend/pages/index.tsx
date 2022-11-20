import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { Select } from 'antd'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
/////
// import { listAsset } from 'api/dashboard/listAsset'
// import { listPlatform } from 'api/dashboard/listPlatform'
import { listStation } from 'api/dashboard/listStation'
/////
// import { listMajorGroup } from 'api/major-group/listMajorGroup'
// import { listEvenness } from 'api/dashboard/listEvenness'
// import { listDiversity } from 'api/dashboard/listDiversity'
////////////////////////////////////
import { SummaryFilter } from 'components/new/SummaryFilter'
import { SummaryBarFilter } from 'components/new/SumFilter'
/////
import { SummaryFilterAsset } from 'components/new/SummaryFilter-asset'
import { SummaryFilterPlatform } from 'components/new/SummaryFilter-platform' 
import { SummaryFilterStation } from 'components/new/SummaryFilter-station'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
import { SummaryFilterGroup } from 'components/new/SummaryFilter-group'
////
import GMap from 'components/map/mapComp'

const { Option } = Select
const index = () => {

  const router = useRouter()
  const { locale } = router

  const [asset, setAsset] = useState([])
  const [platform, setPlatform] = useState([])
  const [station, setStation] = useState([])

  const [majorGroups, setMajorGroups] = useState([])
  const [evenness, setEvenness] = useState([])
  const [diversity, setDiversity] = useState([])

  ////////////////////////////////////////
  // const [evennesyear, setEvennesyear] = useState([])
  // const [evennesproject, setEvennesproject] = useState([])

  // const onFilterChange = (key: string, value?: string) => {
  //   const url = {
  //     pathname: router.pathname,
  //     query: {
  //       ...router.query,
  //       [key]: value,
  //     },
  //   }
  //   router.push(url, url, { locale })
  // }

  // /////
  // const fetchAsset = async () => {
  //   const { data } = await listAsset()
  //   setAsset(data)
  // }
  // const fetchPlatform = async (assetId?: string) => {
  //   const { data } = await listPlatform(assetId)
  //   setPlatform(data)
  // }
  const fetchStation = async (platfomId?: string) => {
    const { data } = await listStation(platfomId)
    setStation(data)
  }
  // /////
  // const fetchMajorGroups = async () => {
  //   const { data } = await listMajorGroup()
  //   setMajorGroups(data)
  // }
  // const fetchEvenness = async (stationId?: string) => {
  //   const { data } = await listEvenness(stationId)
  //   setEvenness(data)
  // }
  // const fetchDiversity = async (stationId?: string) => {
  //   const { data } = await listDiversity(stationId)
  //   setDiversity(data)
  // }
  //////////////////////////////////////
  // const fetchEvennessYear = async () => {
  //   const { data } = await listEvennessYear()
  //   setEvennesyear(data)
  // }
  // const fetchEvennesproject = async (year?: string) => {
  //   const { data } = await listEvennessProject(year)
  //   setEvennesproject(data)
  // }

  // useEffect(() => {
  //   fetchMajorGroups()
  //   fetchAsset()
  //   //fetchEvennessYear()
  // }, [])
  // // useEffect(() => {
  // //   fetchEvennesproject(router.query.year as string)
  // // }, [router.query.year])
  // useEffect(() => {
  //   fetchPlatform(router.query.assetId as string)
  // }, [router.query.assetId])
  useEffect(() => {
    fetchStation(router.query.platfomId as string)
  }, [router.query.platfomId])
  // useEffect(() => {
  //   fetchEvenness(router.query.stationId as string)
  // }, [router.query.stationId])
  // useEffect(() => {
  //   fetchDiversity(router.query.stationId as string)
  // }, [router.query.stationId])
  

  return (
    <PageLayout>
      <div className="w-3/4 h-full m-auto">
        <div className="flex space-x-4 mt-6 ml-4 mr-4 mb-4">
          {/* <div className="flex w-full h-24 border p-4 shadow rounded-md items-center">
            <SummaryFilter/>
          </div> */}
          <div className="flex w-full h-24 border p-4 shadow rounded-md items-center">
            <SummaryFilterAsset/>
            <SummaryFilterPlatform/>
            <SummaryFilterStation/>
            <SummaryFilterYear/>
          </div>
        </div>

      
        <div className="flex space-x-4 h-full mt-6 ml-4 mr-4 mb-4">
          <div className="flex w-2/3 border p-4 shadow rounded-md items-center">
            <div className="w-full text-left">
              {/* <div className="font-bold">Content left header</div>
              <div>Content left</div> */}
              <div className="h-auto w-full">
                <GMap/>
              </div>
            </div>
          </div>
          <div className="w-1/3 border p-1 shadow rounded-md items-center">
      
          <div className="w-full h-60 mb-1 border text-left">
            <div className="md:w-full w-full">
              <SummaryFilterGroup/>
              {/* <Select
                placeholder="Group"
                onChange={(value: string | undefined) => {
                  onFilterChange('majorGroupId', value)
                }}
              >
                {majorGroups.map((majorGroup) => (
                  <Option
                    value={majorGroup?.majorGroupId}
                    key={majorGroup?.majorGroupId}
                  >
                    {majorGroup?.majorGroupName}
                  </Option>
                ))}
              </Select> */}

              

              <div className="font-bold">Station</div>
                <div className="w-full h-128 p-1 overflow-y-auto rounded-md">
                      

                    {station.map((station) => (
                      <div className="flex">
                        {/* <div className="w-1/4 mt-2 mb-2 ml-2 mr-8 border shadow">x</div>  */}
                        <div className="w-3/4 m-auto">
                          <a onClick={() => router.push(`/?stationId=${station?.stationId}`)}>
                            <div className="font-bold">{station?.stationName.toUpperCase()}</div>
                          </a>
                        </div>
                      </div >
                    ))}
              
                </div>

            </div>
          </div>

          </div>
        </div>

        <SummaryBarFilter/>

      </div>
    </PageLayout>
  )
}

export default PrivateRoute(index)

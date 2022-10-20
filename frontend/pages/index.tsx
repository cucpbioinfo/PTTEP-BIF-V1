import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { Select } from 'antd'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
/////
import { listAsset } from 'api/dashboard/listAsset'
import { listPlatform } from 'api/dashboard/listPlatform'
import { listStation } from 'api/dashboard/listStation'
/////
import { listMajorGroup } from 'api/major-group/listMajorGroup'
import { listEvenness } from 'api/dashboard/listEvenness'
import { listDiversity } from 'api/dashboard/listDiversity'
////////////////////////////////////
//import { BarDE } from 'features/echart/barde'
////
// import { listEvennessYear } from 'api/evenness/listEvennessYear'
// import { listEvennessProject } from 'api/evenness/listEvennessProject'

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

  const onFilterChange = (key: string, value?: string) => {
    const url = {
      pathname: router.pathname,
      query: {
        ...router.query,
        [key]: value,
      },
    }
    router.push(url, url, { locale })
  }

  /////
  const fetchAsset = async () => {
    const { data } = await listAsset()
    setAsset(data)
  }
  const fetchPlatform = async (assetId?: string) => {
    const { data } = await listPlatform(assetId)
    setPlatform(data)
  }
  const fetchStation = async (platfomId?: string) => {
    const { data } = await listStation(platfomId)
    setStation(data)
  }
  /////
  const fetchMajorGroups = async () => {
    const { data } = await listMajorGroup()
    setMajorGroups(data)
  }
  const fetchEvenness = async (stationId?: string) => {
    const { data } = await listEvenness(stationId)
    setEvenness(data)
  }
  const fetchDiversity = async (stationId?: string) => {
    const { data } = await listDiversity(stationId)
    setDiversity(data)
  }
  //////////////////////////////////////
  // const fetchEvennessYear = async () => {
  //   const { data } = await listEvennessYear()
  //   setEvennesyear(data)
  // }
  // const fetchEvennesproject = async (year?: string) => {
  //   const { data } = await listEvennessProject(year)
  //   setEvennesproject(data)
  // }

  useEffect(() => {
    fetchMajorGroups()
    fetchAsset()
    //fetchEvennessYear()
  }, [])
  // useEffect(() => {
  //   fetchEvennesproject(router.query.year as string)
  // }, [router.query.year])
  useEffect(() => {
    fetchPlatform(router.query.assetId as string)
  }, [router.query.assetId])
  useEffect(() => {
    fetchStation(router.query.platfomId as string)
  }, [router.query.platfomId])
  useEffect(() => {
    fetchEvenness(router.query.stationId as string)
  }, [router.query.stationId])
  useEffect(() => {
    fetchDiversity(router.query.stationId as string)
  }, [router.query.stationId])
  

  return (
    <PageLayout>

    <div className="w-3/4 h-full m-auto">
    <div className="flex space-x-4 mt-6 ml-4 mr-4 mb-4">
      <div className="flex w-full h-24 border p-4 shadow rounded-md items-center">
        {/* <div className="font-bold">Year *ไม่สามารถทำ group ได้ถ้าใช้ฟังชั่น relation ในการ join table</div>
        <div className="w-1/3 text-left">
          
          <div className="md:w-1/3 w-full">
            
            <Select
              placeholder="year"
              onChange={(value: string | undefined) => {
                onFilterChange('year', value)
              }}
            >
              {evennesyear.map((year) => (
                <Option
                  value={year?.year}
                  key={year?.year}
                >
                  {year?.year}
                </Option>
              ))}
            </Select>
          </div>
        </div> */}
        <div className="w-1/5 text-left">
          <div className="font-bold">Asset</div>
          <div className="md:w-1/3 w-full">
            <Select
              placeholder="asset"
              onChange={(value: string | undefined) => {
                onFilterChange('assetId', value)
              }}
            >
              {asset.map((asset) => (
                <Option
                  value={asset?.assetId}
                  key={asset?.assetId}
                >
                  {asset?.assetName}
                </Option>
              ))}
            </Select>
          </div>
        </div>

        <div className="w-1/5 text-left">
          <div className="font-bold">Platform</div>
          <div className="md:w-1/3 w-full">
            <Select
              placeholder="platfom"
              onChange={(value: string | undefined) => {
                onFilterChange('platformId', value)
              }}
            >
              {platform.map((platform) => (
                <Option
                  value={platform?.platformId}
                  key={platform?.platformId}
                >
                  {platform?.platformName}
                </Option>
              ))}
            </Select>
          </div>
        </div>

        <div className="w-1/5 text-left">
          <div className="font-bold"></div>
          <div className="md:w-1/3 w-full">
          </div>
        </div>

        <div className="w-1/5 text-left">
          <div className="font-bold"></div>
          <div className="md:w-1/3 w-full">
          </div>
        </div>

        <div className="w-1/5 text-Right text-red-600">
          <div className="font-bold">Year Filter Here</div>
          <div className="md:w-1/3 w-full">
          </div>
        </div>

        <div className="w-1/5 text-left">
          <div className="font-bold">Station</div>
          <div className="md:w-1/3 w-full">
            <Select
              placeholder="station"
              onChange={(value: string | undefined) => {
                onFilterChange('stationId', value)
              }}
            >
              {station.map((station) => (
                <Option
                  value={station?.stationId}
                  key={station?.stationId}
                >
                  {station?.stationName}
                </Option>
              ))}
            </Select>
          </div>
        </div>


      </div>
    </div>

    
    <div className="flex space-x-4 h-full mt-6 ml-4 mr-4 mb-4">
      <div className="flex w-2/3 border p-4 shadow rounded-md items-center">
        <div className="w-2/3 text-left">
          <div className="font-bold">Content left header</div>
          <div>Content left</div>
          <div className="bg-black h-auto w-full">
            map
          </div>
        </div>
        <div className="w-1/3 text-left">
          <div className="font-bold">Station</div>
          <div className="w-full h-128 p-1 overflow-y-auto rounded-md">
                

              {station.map((station) => (
                <div className="flex h-24 ">
                  {/* <div className="w-1/4 mt-2 mb-2 ml-2 mr-8 border shadow">x</div>  */}
                  <div className="w-3/4 m-auto">
                    <a onClick={() => router.push(`/dashboard?stationId=${station?.stationId}`)}>
                      <div className="font-bold">{station?.stationName}</div>
                      <div className="font-bold text-red-600">API location Here.</div>
                    </a>
                  </div>
                </div >
              ))}


                

                {/* <div className="flex h-24 ">
                  <div className="w-1/4 mt-2 mb-2 ml-2 mr-8 border shadow">x</div> 
                  <div className="w-3/4 m-auto">
                    <div className="font-bold">platform 1</div>
                    <div>platform 1</div>
                    <div>platform 1</div>
                  </div>
                </div >

                <div className="flex h-24 ">
                  <div className="w-1/4 mt-2 mb-2 ml-2 mr-8 border shadow">x</div> 
                  <div className="w-3/4 m-auto">
                    <div className="font-bold">platform 2</div>
                    <div>platform 2</div>
                    <div>platform 2</div>
                  </div>
                </div > */}

        
          </div>

        </div>
      </div>
      <div className="w-1/3 border p-1 shadow rounded-md items-center">
  
      <div className="w-full h-60 mb-1 border text-left">
        <div className="md:w-1/3 w-full">
          <Select
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
          </Select>
        </div>
      </div>

      <div className="w-full h-60 mb-1 border text-left">Evenness </div>
      <div className="w-full h-60 mb-1 border text-left">
        {/* <BarDE barname='Evenness' stationId='demo'/> */}
      </div>
      {/* number figure 
      Shannon & Evenness (bar)
      number of species
      */}
      <div className="w-full h-60 mb-1 border text-left">Diversity</div>
      <div className="w-full h-60 mb-1 border text-left">
        {/* <BarDE barname='Diversity' stationId='demo'/> */}
      </div>
       {/* Bar Chart year>project>platform */}
      {/* <div className="w-full h-60 mb-1 border text-left">idea ใช้ project เป็น query ในการดึง api มาทำ line กราฟ เปรียบเทียบค่าแต่ละปี โดยใช้ค่า average</div> */}
      
      </div>
    </div>




    <div className="flex space-x-4 mt-6 ml-4 mr-4 mb-8">
    
      <div className="flex w-2/4 h-40 border p-4 shadow rounded-md items-center text-red-600">
        <div className="w-1/2 text-left">
          <div className="font-bold">Graph Algorithm Idea</div>
          <div>Fetch Data As Loop. </div>
        </div>
        <div className="w-1/2 text-right pr-4">Declare Array For kept data from fetch. push data to array each loop "map".</div>
        <div className="w-1/2 text-right pr-4">use array form as data for graph.</div>
        
      </div>

      
      <div className="w-1/4 border p-4 shadow rounded-md">
        {evenness.map((evenness) => (
          <div className="flex h-24 ">
            <div className="w-1/4 mt-2 mb-2 ml-2 mr-8 border shadow">x</div> 
            <div className="w-3/4 m-auto">
              <div className="font-bold">Year : {evenness?.year}</div>
              <div>{evenness?.assetName} / {evenness?.platformName} / {evenness?.stationName}</div>
              <div>surface : {evenness?.surface}</div>
              <div>euphotic_zone : {evenness?.euphotic_zone}</div>
          </div>
        </div >
        ))}
      </div>

      <div className="w-1/4 border p-4 shadow rounded-md">
        {diversity.map((diversity) => (
          <div className="flex h-24 ">
            <div className="w-1/4 mt-2 mb-2 ml-2 mr-8 border shadow">x</div> 
            <div className="w-3/4 m-auto">
              <div className="font-bold">Year : {diversity?.year}</div>
              <div>{diversity?.assetName} / {diversity?.platformName} / {diversity?.stationName}</div>
              <div>surface : {diversity?.surface}</div>
              <div>euphotic_zone : {diversity?.euphotic_zone}</div>
          </div>
        </div >
        ))}
      </div>

    </div>

    {/* <div>
      <div>project api</div>
      {project.map((project) => (
                <li key={project?.projectId}>
                  {project?.projectName}
                </li>
                ))}
    </div>

    <div>
    <div>platform api</div>
      {platform.map((platform) => (
                <li key={platform?.platformId}>
                  {platform?.platformName}/
                  {platform?.projectName}/
                  {platform?.Latitude}/
                  {platform?.Longitude}/
                </li>
                ))}
    </div>


    <div>
    <div>evenness api</div>
      {evenness.map((evenness) => (
                <li key={evenness?.evennessId}>
                  {evenness?.year}/
                  {evenness?.projectName}/
                  {evenness?.PlatformName}/
                  {evenness?.surface}/
                  {evenness?.euphotic_zone}/
                </li>
                ))}
    </div>

    <div>
    <div>platform api</div>
      {platform.map((platform) => (
                <li key={platform?.platformId}>
                  {platform?.platformName}
                </li>
                ))}
    </div>
 */}



      
    
  </div>


  







    </PageLayout>
  )
}

export default PrivateRoute(index)

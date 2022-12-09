import { Divider, Empty ,Button,Tooltip } from 'antd'
import { listDen } from 'api/species/listDen'
import { DenBar } from 'features/echart/DenBar'
import { DensityFilter } from 'components/new/DensityFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { InfoCircleOutlined } from '@ant-design/icons';
import { Diversityinfo } from './DefInfo'

export const DensityBarFilter = () => {
  const router = useRouter()
  const [den, setDen] = useState([])
  let sid = ""
  const fetchDen = async () => {
    const {
      speciesId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listDen({
      speciesId: speciesId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setDen(data)
  }

  useEffect(() => {
    fetchDen()
  }, [router.query])

  den.map(({ speciesId }) => (
    sid = speciesId
  ))

  function SpeciesSplice(string:String) {
    console.log(string)
    if(string.search(/sp./i)==-1){
      let touppercase = string.charAt(0).toUpperCase() + string.slice(1);
      return <>{touppercase}</>;
    }
    else
    {
      let position = string.search(/sp./i);
      let speciesName = string.substring(0, position);
      let upperspeciesName = speciesName.charAt(0).toUpperCase() + speciesName.slice(1);
      let SpecialText = string.slice(position);
      return {upperspeciesName}+" "+{SpecialText};
    }
  }
  
  return (
    <div>
      
      {/* <DensityFilter/> */}
      <Divider />
      
      {/* <a href="/density?speciesId=e1dced3d-653b-4f83-9ae8-b341c4d85523">View all</a> */}

      {/* <Button onClick={() => router.push(`/density?speciesId=${sid}`)}>See all</Button> */}

      <div>{den.length === 0 && <Empty />}</div>
      <div>
        {den.length !== 0 &&
          den.map(({ densityId ,speciesName,assetName,platformName,stationName, year, surface ,euphotic_zone }) => (
            <div>
              <div className="flex w-full items-center divide-x mr-4">
                  
                  <div>
                    <div className="mr-4 ml-4"><b>Asset</b> : {assetName.toUpperCase()}</div>
                  </div>
                  
                  <div>
                    <div className="mr-4 ml-4"><b>Platform</b> : {platformName.toUpperCase()}</div>
                  </div>

                  <div>
                    <div className="mr-4 ml-4"><b>Station</b> : {stationName.toUpperCase()}</div>
                  </div>

                  {/* <div>
                    <div className="mr-4 ml-4"><b>Species Name</b> : {speciesName}</div>
                  </div> */}

              </div>
            <Divider />
            {/* <div className="flex w-full items-center">       
              <div className="flex w-full items-center divide-x mr-4">
                <div>
                  <div className="mr-4 ml-2"></div>
                </div>
              </div>
              <div>
                <div className="mr-2">
                  <Diversityinfo/>
                  <Tooltip placement="top" title="def"><InfoCircleOutlined/></Tooltip>
                </div>
              </div>
            </div> */}

            <div className="w-full text-lg text-left">
              <b>Density</b>
            </div>
            {/* {console.log(SpeciesSplice(speciesName))}
            <div>{SpeciesSplice(speciesName)}</div> */}
            {/* <div>{SpeciesSplice(JSON.stringify(speciesName))}</div> */}
            <DenBar densityId={densityId} speciesName={speciesName} name={stationName} year={year} surface={surface} zone={euphotic_zone} />
            </div>
          ))}
      </div>
  
      </div>
  )
}

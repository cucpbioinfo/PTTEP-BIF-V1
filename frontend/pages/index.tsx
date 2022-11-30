import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { Divider,Button } from 'antd'
import { useRouter } from 'next/router'
import React from 'react'
/////
import { SummaryBarFilter } from 'components/new/SumFilter'
/////
import { SummaryFilterAsset } from 'components/new/SummaryFilter-asset'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
////
import GMap from 'components/map/mapComp'
// import { DashboardFilter } from 'components/new/Dashboard-Filter'
import { DashboardRightPanalStation } from 'components/new/DashboardRightPanalStation'
import { DashboardRightPanalPlatform } from 'components/new/DashboardRightPanalPlatform'

const index = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <div className="w-3/4 h-full m-auto">
        <div className="flex space-x-4 mt-6 ml-4 mr-4 mb-4">
          <div className="flex justify-between w-full p-4 border shadow rounded-md">  
              <SummaryFilterAsset/>
              {/* <DashboardFilter/> */}
              {/* <SummaryFilterYear/>  */}
              <div>
                  <div className="mr-2"><Button type="primary" href="/" >Clear</Button></div>
                </div>
              {/* <div className="mr-2">
                <Button onClick={() => router.push(`/`)}>clear</Button>
              </div>  */}
              
          </div>
        </div>
        <div className="flex space-x-4 h-full mt-6 ml-4 mr-4 mb-4">
          <div className="flex w-2/3 border p-4 shadow rounded-md items-center">
            <div className="w-full text-left">
              <div className="h-auto w-full">
                <GMap/>
              </div>
            </div>
          </div>
          <div className="w-1/3 border p-1 shadow rounded-md items-center h-96 overflow-y-auto">
            <div className="w-full h-auto mb-1 text-left ">
              <div className="md:w-full w-full"></div> 
                <DashboardRightPanalPlatform/>
                {/* <Divider/> */}
                {/* <DashboardRightPanalStation/> */}
            </div>
          </div>
        </div>
        <SummaryBarFilter/>
      </div>
    </PageLayout>
  )
}

//export default PrivateRoute(index)//Bypass
export default index

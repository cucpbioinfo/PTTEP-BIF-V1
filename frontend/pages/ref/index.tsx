import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { Divider,Button } from 'antd'
import { useRouter } from 'next/router'
import React from 'react'
import { DefDiversity,DefEvenness,DefNumber,DefSam,DefDensity } from 'components/new/DefInfo'
/////

const index = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <div className="w-3/4 h-full m-auto">
        <div className="font-bold text-base">References</div>
        <br/>
        <DefDiversity/>
        <Divider/>
        <DefEvenness/>
        <Divider/>
        <DefNumber/>
        <Divider/>
        <DefDensity/>
        <Divider/>
        <DefSam/>
      </div>
    </PageLayout>
  )
}

//export default PrivateRoute(index)//Bypass
export default index

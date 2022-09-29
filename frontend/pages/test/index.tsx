import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import React from 'react'
import { SunburstChart } from 'features/sunburstd3/SunbrustChart'
const index = () => {
  return (
    <PageLayout>
      <h2>Sunburst Chart by d3 comp</h2>
      <SunburstChart/>
    </PageLayout>
  )
}

export default PrivateRoute(index)

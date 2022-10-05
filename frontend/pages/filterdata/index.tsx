import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { FilterData } from 'features/filter/filterdata'
import React from 'react'

const index = () => {
  return (
    <PageLayout>
      <FilterData />
    </PageLayout>
  )
}

export default PrivateRoute(index)

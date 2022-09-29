import { Table } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { LocationTable } from 'features/new/LocationTable'
import React from 'react'

const index = () => {
  return (
    <PageLayout>
      <LocationTable />
    </PageLayout>
  )
}

export default PrivateRoute(index)

import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { InternalDataPage } from 'features/internal-data/InternalDataPage'
import React from 'react'

const index = () => {
  return (
    <PageLayout>
      <InternalDataPage />
    </PageLayout>
  )
}

export default PrivateRoute(index)

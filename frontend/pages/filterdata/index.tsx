import { PageLayout } from 'components/new/PageLayout'
import { AssetDensityBarFilter } from 'components/new/AssetFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const assetbar = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <AssetDensityBarFilter onClick={() => router.push(`/species/${'a906251d-6916-4a2f-b035-129f6cb3d59f'}`)}/>
    </PageLayout>
  )
}

export default assetbar

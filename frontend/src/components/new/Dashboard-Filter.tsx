import { Select, Button } from 'antd'
import React, { useEffect, useState } from 'react'
import { useRouter } from 'next/router'

export const DashboardFilter = () => {
  const router = useRouter()
  const { locale } = router
  const IsEmp = (Data:any) => {
    if (Data!="") {
      return (
        <>
          {Data}
        </>
      )
    }
  }

  return (
    <>
      <Button onClick={() => router.push(`/?platformId=${router.query.platformId}&stationId=${router.query.stationId}`)}>{IsEmp(router.query.assetId)}</Button>
    </>
  )
}

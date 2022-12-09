import { PageLayout } from 'components/new/PageLayout'
import { PrivateRoute } from 'components/PrivateRoute'
import { Divider,Button } from 'antd'
import { useRouter } from 'next/router'
import React from 'react'
/////

const index = () => {
  const router = useRouter()
  return (
    <PageLayout>
      <div className="w-3/4 h-full m-auto">
      <ul className="list-disc">
        <li>Now this is a story all about how, my life got flipped-turned upside down</li>
      </ul>

      <ol className="list-decimal">
        <li>Now this is a story all about how, my life got flipped-turned upside down</li>
      </ol>

      <ul className="list-none">
        <li>Now this is a story all about how, my life got flipped-turned upside down</li>
      </ul>
      </div>
    </PageLayout>
  )
}

//export default PrivateRoute(index)//Bypass
export default index

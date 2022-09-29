/** @jsxRuntime classic */
/** @jsx jsx */
import { css, jsx } from '@emotion/react'
import { Collapse, Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import React, { useEffect, useMemo, useState } from 'react'
import { SpeciesOverviewTab } from 'features/species/SpeciesOverviewTab'
import { SpeciesDetail } from './SpeciesDetail'
import { SPECIES_DETAIL } from 'utils/mockData'
import Carousel from 'react-multi-carousel'
import 'react-multi-carousel/lib/styles.css'
import { getSpeciesDetails } from 'api/species/getSpeciesDetails'
import { SpeciesDetailsResponse } from 'types/species'
import { LoadingOutlined } from '@ant-design/icons'
import { Bar2 } from 'features/echart/bar2' 
import { Pie } from 'features/echart/pie' 

const { Panel } = Collapse

const { TabPane } = Tabs

export const SpeciesDetailPage = ({ speciesId }) => {
  const [species, setSpecies] = useState<SpeciesDetailsResponse | undefined>()
  const fetchSpeciesDetails = async (speciesId: string) => {
    const data = await getSpeciesDetails(speciesId)
    setSpecies(data)
  }

  useEffect(() => {
    if (speciesId) {
      fetchSpeciesDetails(speciesId)
    }
  }, [speciesId])

  if (!species) {
    return <LoadingOutlined />
  }

  return (
    <PageLayout>
      <div className="text-2xl font-bold">
        <i>{species && species?.speciesName}</i>
      </div>
      <div className="block md:grid md:grid-cols-2 mt-2 md:mt-4 md:gap-16">
        <Collapse bordered={false} ghost defaultActiveKey={['1']}>
          <Panel key="1" header={`Species details`}>
            <SpeciesDetail speciesDetail={species} />
          </Panel>
        </Collapse>
        <div className="mt-8">
          <Tabs defaultActiveKey="1">
            <TabPane tab="Overview" key="1">
              <SpeciesOverviewTab speciesDetail={species}/>
            </TabPane>
            <TabPane tab="Metrics" key="2">
              <Pie SpeciesName={species?.speciesName}/>
              <Bar2 SpeciesName={species?.speciesName}/>
            </TabPane>
            {/* 26 new eDNA tab */}
            {/* <TabPane tab="eDNA" key="3">
              similarlity index
            </TabPane> */}
          </Tabs>
        </div>
      </div>
    </PageLayout>
  )
}

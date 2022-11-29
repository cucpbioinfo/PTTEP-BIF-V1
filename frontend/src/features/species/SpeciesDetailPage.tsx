/** @jsxRuntime classic */
/** @jsx jsx */
import { css, jsx } from '@emotion/react'
import { Collapse, Tabs } from 'antd'
import { PageLayout } from 'components/new/PageLayout'
import React, { useEffect, useState } from 'react'
import { SpeciesOverviewTab } from 'features/species/SpeciesOverviewTab'
import { SpeciesDetail } from './SpeciesDetail'
import 'react-multi-carousel/lib/styles.css'
import { getSpeciesDetails } from 'api/species/getSpeciesDetails'
import { SpeciesDetailsResponse } from 'types/species'
import { LoadingOutlined } from '@ant-design/icons'
import { DensityBarFilter } from 'components/new/Filter'
import { SummaryFilterAsset } from 'components/new/SummaryFilter-asset'
import { SummaryFilterPlatform } from 'components/new/SummaryFilter-platform' 
import { SummaryFilterStation } from 'components/new/SummaryFilter-station'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
import { SummaryFilterGroup } from 'components/new/SummaryFilter-group'

const { Panel } = Collapse
const { TabPane } = Tabs

function SpeciesSplice(string:string) {
  if(string.search(/sp./i)==-1){
    let touppercase = string.charAt(0).toUpperCase() + string.slice(1);
    return <><i>{touppercase}</i></>;
  }
  else
  {
    let position = string.search(/sp./i);
    let speciesName = string.substring(0, position);
    let upperspeciesName = speciesName.charAt(0).toUpperCase() + speciesName.slice(1);
    let SpecialText = string.slice(position);
    return <><i>{upperspeciesName}</i> {SpecialText}</>;
  }
  }
  
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
        {SpeciesSplice(species?.speciesName)}
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
              <Tabs defaultActiveKey="3">
                {/* <TabPane tab="Asset" key="1" >
                  <AssetDensityBarFilter />
                </TabPane>
                <TabPane tab="Platform" key="2">
                  <PlatformDensityBarFilter/>
                </TabPane> */}
                <TabPane tab="Station" key="3">
                  <SummaryFilterAsset/>
                  <SummaryFilterPlatform/>
                  <SummaryFilterStation/>
                  <SummaryFilterGroup/>
                  <SummaryFilterYear/>
                  <DensityBarFilter />
                </TabPane>

              </Tabs>
              
            </TabPane>

          </Tabs>
        </div>
      </div>
    </PageLayout>
  )
}

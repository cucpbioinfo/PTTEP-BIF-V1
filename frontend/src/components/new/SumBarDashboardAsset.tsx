import { Divider, Empty,Button,Tooltip,Tabs } from 'antd'
import { InfoCircleOutlined } from '@ant-design/icons';
import { listSumAsset } from 'api/species/listSumAsset'
import { listSumPlatform } from 'api/species/listSumPlatform'
import { listSumStation } from 'api/species/listSumStation'
import { DenBar } from 'features/echart/DenBar'
import { SumBar } from 'features/echart/SumBar'
import { SummaryFilter } from 'components/new/SummaryFilter'
import { SummaryFilterYear } from 'components/new/SummaryFilter-year' 
import { SummaryFilterGroup } from 'components/new/SummaryFilter-group'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { DefDiversity } from "components/new/DefInfo";

const DiversityInfo = <span>The Shannon-Weiner Species Diversity Index is an index that is commonly used to characterize species diversity in a community. 
This index accounts for both abundance and evenness of the species present. A large number of species can increase diversity. 
Similarly, increasing the uniformity of individual distribution among species will also increase diversity. If each individual 
belongs to a different species, the diversity index is the largest. In contrast, if there is/are dominant species (1 or more species), 
the diversity index will be decreased.

The Shannon-Weiner Species Diversity Index is calculated by taking the number of each species, the proportion each species 
is of the total number of individuals, and sums the proportion times the natural log of the proportion for each species. 
The formula is as follows

Where H’ is the species diversity index, 
s is the number of species
pi is the proportion of individuals of each species belonging to the ith species of the total number of individuals.

References
https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index
Nolan, K.A. and J.E. Callahan. 2006. Beachcomber biology: The Shannon-Weiner Species Diversity Index. 
Pages 334-338, in Tested Studies for Laboratory Teaching, Volume 27 
(M.A. O'Donnell, Editor). Proceedings of the 27th Workshop/Conference of the Association for Biology Laboratory Education (ABLE), 383 pages.<a>Link</a></span>;

const EvennessInfo = <span>Definition of about Evenness Index.</span>;
const NumberInfo = <span>Definition of about Number of Species Index.</span>;

export const SummaryBarDashboardAsset = () => {
  const router = useRouter()
  const [suma, setSuma] = useState([])
  const { TabPane } = Tabs
  const fetchSuma = async () => {
    const {
      majorGroupId,
      assetId,
      platformId,
      stationId,
      year,
      
    } = router.query
    const { data } = await listSumAsset({
      majorGroupId: majorGroupId as string,
      assetId: assetId as string,
      platformId: platformId as string,
      stationId: stationId as string,
      year: year as string,
      
    })
    setSuma(data)
  }

  useEffect(() => {
    fetchSuma()
  }, [router.query])

  const SepType = () => {
    
    return (
      <>
      <div>
      <div>{suma.length === 0 && <Empty 
      description={
        <span>
          Please select asset for sorting data.
        </span>
      }
      />}</div>
      <div>
        {suma.length !== 0 &&
          suma.map(({ summaryId,year,majorGroupName,assetName,platformName,stationName,surfaceEvenness,euphoticzoneEvenness,surfaceShannon,euphoticzoneShannon,surfaceNumber,euphoticzoneNumber}) => (
            <div>
              <div className="flex w-full items-center border p-4 shadow rounded-md mb-1">
                <div className="flex w-full items-center divide-x mr-4">
                  
                  <div>
                    <div className="mr-4 ml-2"><b>Asset : {assetName.toUpperCase()}</b></div>
                  </div>
                  
                  <div className="flex">
                    <div className="ml-5"><SummaryFilterGroup/></div>
                    <div className="ml-2"><SummaryFilterYear/></div>
                  </div>
                </div>
                <div>
                  <div className="mr-2"><Button type="primary" href="/summarystation" >View All</Button></div>
                </div>
            </div>
                <div className="flex space-x-4">
                  {/* ------------------------------------------------------------- */}
                  <div className="w-1/3 border p-3 shadow rounded-md">
                    <div className="flex w-full items-center">
                      
                      <div className="flex w-full items-center divide-x mr-4">
                        <div>
                          <div className="mr-4 ml-2"></div>
                        </div>
                      </div>
                      <div>
                        <div className="mr-2">
                          <Tooltip placement="top" title="def"><InfoCircleOutlined/></Tooltip>
                        </div>
                      </div>

                    </div>

                    <div className="w-full text-lg text-left">
                      <b>Shannon-Weiner Species Diversity Index</b>
                    </div>
                    <br/>
                    <div>
                      <SumBar densityId={summaryId} speciesName="Diversity" name={stationName} year={year} surface={surfaceShannon} zone={euphoticzoneShannon} />
                    </div>
                  </div>
                  {/* ------------------------------------------------------------- */}
                  <div className="w-1/3 border p-3 shadow rounded-md">
                    <div className="flex w-full items-center">
                      
                      <div className="flex w-full items-center divide-x mr-4">
                        <div>
                          <div className="mr-4 ml-2"></div>
                        </div>
                      </div>
                      <div>
                        <div className="mr-2">
                          <Tooltip placement="top" title={EvennessInfo}><InfoCircleOutlined/></Tooltip>
                        </div>
                      </div>

                    </div>

                    <div className="w-full text-lg text-left">
                      <b>Evenness Index</b>
                    </div>
                    <br/>
                    <div>
                      <SumBar densityId={summaryId} speciesName="Evenness" name={stationName} year={year} surface={surfaceEvenness} zone={euphoticzoneEvenness} />
                    </div>
                  </div>
                  {/* ------------------------------------------------------------- */}
                  <div className="w-1/3 border p-3 shadow rounded-md">
                    <div className="flex w-full items-center">
                      
                      <div className="flex w-full items-center divide-x mr-4">
                        <div>
                          <div className="mr-4 ml-2"></div>
                        </div>
                      </div>
                      <div>
                        <div className="mr-2">
                          <Tooltip placement="top" title={NumberInfo}><InfoCircleOutlined/></Tooltip>
                        </div>
                      </div>

                    </div>

                    <div className="w-full text-lg text-left">
                      <b>Number of Species</b>
                    </div>
                    <br/>
                    <div>
                      <SumBar densityId={summaryId} speciesName="No." name={stationName} year={year} surface={surfaceNumber} zone={euphoticzoneNumber} />
                    </div>
                  </div>
                  {/* ------------------------------------------------------------- */}
                  

                  {/* <div className="w-1/3 border p-4 shadow rounded-md">
                    <div className="w-full h-60 mb-1 border text-left">
                      Evenness Index <Tooltip placement="top" title={EvennessInfo}><InfoCircleOutlined/></Tooltip>
                    </div>
                    <div>
                      <SumBar densityId={summaryId} speciesName="Evenness" name={stationName} year={year} surface={surfaceEvenness} zone={euphoticzoneEvenness} />
                    </div>
                  </div> */}


                  {/* <div className="w-1/3 border p-4 shadow rounded-md">
                    <div className="w-full h-60 mb-1 border text-left">
                      Number of Species <Tooltip placement="top" title={NumberInfo}><InfoCircleOutlined/></Tooltip>
                    </div>
                    <div>
                      <SumBar densityId={summaryId} speciesName="No." name={stationName} year={year} surface={surfaceNumber} zone={euphoticzoneNumber} />
                    </div>
                  </div> */}

                </div>
          </div>
          ))}
      </div>
  
      </div>
      </>
    )
  }

  return (
    <>
      <SepType/>
    </>
  )
}

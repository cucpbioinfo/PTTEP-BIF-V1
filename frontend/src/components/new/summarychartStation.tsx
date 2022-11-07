//to fetch data and make it to array
import { useRouter } from 'next/router'
import React, { useCallback,useEffect, useState } from 'react'
//fetch year from db to have it as filter
import { listYearSummary } from 'api/summary/listYearSummary'
import { listAllSummary } from 'api/summary/listAllSummary'
import  {Testcheckbox} from 'features/echart/testfilter'
export const SummaryChartStation = () => {
  const router = useRouter()
  //yearfilter
  const [yearfiter, setYearFilter] = useState([])
  //summary array
  const [summary, setSummary] = useState([])

  interface IProps {}
  interface IState {}

  //make array of Yearfilter
   const Yearfilter = []
  const fetchYearFilter = async (platformId?: string) => {
    const { data } = await listYearSummary(platformId)
    setYearFilter(data)
  }
  useEffect(() => {
    fetchYearFilter(router.query.platformId as string)
  }, [router.query.platformId])
   yearfiter.map((year) => (Yearfilter.push(year.year)))
//    console.log(Yearfilter)
  //make array of summarylist
//   const summaryyear = []
//   const summarymajorGroupName = []
//   const summaryidentificationName = []
//   const summaryassetName = []
//   const summaryplatformName = []
//   const summarystationName = []
//   const summarysurfaceShannon = []
//   const summarysurfaceNumber = []
//   const summarysurfaceMax = []
//   const summarysurfaceEvenness = []
//   const summaryeuphoticzoneShannon = []
//   const summaryeuphoticzoneNumber = []
//   const summaryeuphoticzoneMax = []
//   const summaryeuphoticzoneEvenness = []
  const SummaryList = []
  const fetchSummary = async () => {
    const {
        majorGroupId,
        assetId,
        platformId,
        stationId,
        year,
      } = router.query
      const { data } = await listAllSummary({
        majorGroupId: majorGroupId as string,
        assetId: assetId as string,
        platformId: platformId as string,
        stationId: stationId as string,
        year: year as string,
      })
      setSummary(data)
    }
    useEffect(() => {
        fetchSummary()
      }, [router.query])
    
      summary.map((sum) => (
        // summaryyear.push(sum.year),
        // summarymajorGroupName.push(sum.majorGroupName),
        // summaryidentificationName.push(sum.identificationName),
        // summaryassetName.push(sum.assetName),
        // summaryplatformName.push(sum.platformName),
        // summarystationName.push(sum.stationName),
        // summarysurfaceShannon.push(sum.surfaceShannon),
        // summarysurfaceNumber.push(sum.surfaceNumber),
        // summarysurfaceMax.push(sum.surfaceMax),
        // summarysurfaceEvenness.push(sum.surfaceEvenness),
        // summaryeuphoticzoneShannon.push(sum.euphoticzoneShannon),
        // summaryeuphoticzoneNumber.push(sum.euphoticzoneNumber),
        // summaryeuphoticzoneMax.push(sum.euphoticzoneMax),
        // summaryeuphoticzoneEvenness.push(sum.euphoticzoneEvenness),
        SummaryList.push(
            {
                "summaryId": sum.summaryId,
                "year": sum.year,
                "majorGroupName": sum.majorGroupName,
                "identificationName": sum.identificationName,
                "assetName": sum.assetName,
                "platformName": sum.platformName,
                "stationName": sum.stationName,
                "surfaceShannon": sum.surfaceShannon,
                "surfaceNumber": sum.surfaceNumber,
                "surfaceMax": sum.surfaceMax,
                "surfaceEvenness": sum.surfaceEvenness,
                "euphoticzoneShannon": sum.euphoticzoneShannon,
                "euphoticzoneNumber": sum.euphoticzoneNumber,
                "euphoticzoneMax": sum.euphoticzoneMax,
                "euphoticzoneEvenness": sum.euphoticzoneEvenness
            }
        )
    ))
    // if(SummaryList.length != 0){
    //     console.log(SummaryList);
    //   }
    //console.log(SummaryList)
    
  return (
    <>
      <div>SummaryChartStation - test</div>
      {/* <div>{Yearfilter[0]}</div>
      <div>{summaryyear.length}</div>
      <div>{summaryplatformName.length}</div>
      <div>{summaryeuphoticzoneShannon.length}</div>
      <div>{summaryeuphoticzoneEvenness.length}</div> */}
      {/* <div>{SummaryList.length}</div> */}
      {/* <SummaryFilterGraph data={testarray} datafilter={Yearfilter} /> */}
      {/* <<< END OF TEST DATA */}
      <Testcheckbox filter={Yearfilter} datax={SummaryList} />
      {/* <SummaryFilterGraph /> */}
      {/* <YearFilters 
          categories={Yearfilter}
          onFilterChange={handleFilterChange}/>
        <ProductsList products={state.products} /> */}

      
    </>
  )
}

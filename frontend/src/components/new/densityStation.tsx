import { Select,Button } from 'antd'
//to fetch data and make it to array
import { useRouter } from 'next/router'
import React, { useCallback,useEffect, useState } from 'react'
//Dropdown Filter API
//import { SummaryFilter } from 'components/new/SummaryFilter'
//fetch year from db to have it as filter
import { listYearSummary } from 'api/summary/listYearSummary'
//import { listAllSummary } from 'api/summary/listAllSummary'

import { listAllDensity } from 'api/density/listAllDensity'
//
import { listAsset } from 'api/dashboard/listAsset'
import { listPlatform } from 'api/dashboard/listPlatform'
import { listStation } from 'api/dashboard/listStation'
//
import { DensityBar } from 'features/echart/densitybar'
//import  {Testcheckbox} from 'features/echart/testfilter'
const { Option } = Select
export const DensityStation = () => {
  const router = useRouter()
  const { locale } = router
  //yearfilter
  const [yearfiter, setYearFilter] = useState([])
  //density array
  const [density, setDensity] = useState([])
  //
  const [asset, setAsset] = useState([])
  const [platform, setPlatform] = useState([])
  const [station, setStation] = useState([])
  //
  const onFilterChange = (key: string, value?: string) => {
    const url = {
      pathname: router.pathname,
      query: {
        ...router.query,
        [key]: value,
      },
    }
    router.push(url, url, { locale })
  }

  const fetchAsset = async () => {
    const { data } = await listAsset()
    setAsset(data)
  }
  const fetchPlatform = async (assetId?: string) => {
    const { data } = await listPlatform(assetId)
    setPlatform(data)
  }
  const fetchStation = async (platformId?: string) => {
    const { data } = await listStation(platformId)
    setStation(data)
  }
  //
  function DropdownFilter(){
    return (
      <>
      <div className="text-sm">
        Asset
      </div>
      <div className="md:w-1/6 w-full">
        <Select
          placeholder="Asset"
          onChange={(value: string | undefined) => {
            onFilterChange('assetId', value)
          }}
        >
          {asset.map((asset) => (
            <Option
              value={asset?.assetId}
              key={asset?.assetId}
            >
              {asset?.assetName.toUpperCase()}
            </Option>
          ))}
        </Select>
      </div>

      <div className="text-sm">
        Platform
      </div>
      <div className="md:w-1/6 w-full">
        <Select
          placeholder="Platform"
          onChange={(value: string | undefined) => {
            onFilterChange('platformId', value)
          }}
        >
          {platform.map((platform) => (
            <Option
              value={platform?.platformId}
              key={platform?.platformId}
            >
              {platform?.platformName.toUpperCase()}
            </Option>
          ))}
        </Select>
      </div>

      <div className="text-sm">
        Station
      </div>
      <div className="md:w-1/5 w-full">
        <Select
          placeholder="Station"
          onChange={(value: string | undefined) => {
            onFilterChange('stationId', value)
          }}
        >
          {station.map((station) => (
            <Option
              value={station?.stationId}
              key={station?.stationId}
            >
              {station?.stationName.toUpperCase()}
            </Option>
          ))}
        </Select>
      </div>
      </>
    )
  }
  //
  useEffect(() => {
    fetchAsset()
  }, [])
  useEffect(() => {
    fetchPlatform(router.query.assetId as string)
  }, [router.query.assetId])
  useEffect(() => {
    fetchStation(router.query.platformId as string)
  }, [router.query.platformId])
  const Stationfilter = []
  station.map((s) => (Stationfilter.push(s.stationName)))
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
  const DensityList = []
  const fetchDensity = async () => {
    const {
        speciesId,
        assetId,
        platformId,
        stationId,
        year,
      } = router.query
      const { data } = await listAllDensity({
        speciesId: speciesId as string,
        assetId: assetId as string,
        platformId: platformId as string,
        stationId: stationId as string,
        year: year as string,
      })
      setDensity(data)
    }
    useEffect(() => {
        fetchDensity()
      }, [router.query])
    
      density.map((den) => (
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
        DensityList.push(
            {
                "densityId": den.summaryId,
                "year": den.year,
                "assetName": den.assetName,
                "platformName": den.platformName,
                "stationId": den.stationId,
                "stationName": den.stationName,
                "speciesId": den.speciesId,
                "speciesName": den.speciesName,
                "surface": den.surface,
                "euphotic_zone": den.euphotic_zone  
            }
        )
    ))

    const Datatest = [
      {
        "densityId": "densityId1",
        "year": "2020",
        "assetName": "assetName1",
        "platformName": "platformName1",
        "stationId": "stationId1",
        "stationName": "stationName1",
        "speciesId": "speciesId1",
        "speciesName": "speciesName",
        "surface": "1.5",
        "euphotic_zone": "3.5"  
    },
    {
      "densityId": "densityId2",
      "year": "2020",
      "assetName": "assetName2",
      "platformName": "platformName2",
      "stationId": "stationId2",
      "stationName": "stationName2",
      "speciesId": "speciesId2",
      "speciesName": "speciesName2",
      "surface": "3.5",
      "euphotic_zone": "1.5"  
    }
    ]
    // if(SummaryList.length != 0){
    //     console.log(SummaryList);
    //   }
    //console.log(SummaryList)
    /////////////////////////////////////////////////////
    function ProductFilters(props) {
        const { 
          categories,
          onFilterChange,
        } = props
        
        return (
          <section 
            className="filters"
            aria-labelledby="filters-header">
            <header id="filters-header">
              {'Filters'}
            </header>
            
            <ul className="flex m-2">
              {categories.map(category => (
                <li className=" m-2" key={category}>
                  <label>
                    <input 
                      onChange={onFilterChange}
                      type="checkbox"
                      value={category} />
                    {" "+category.toUpperCase()}
                  </label>
                </li>
              ))}
            </ul>
          </section>
        )
      }
      
      function Product(props) {
        const { product } = props
        
        return (
          <li
            key={product.SummaryId}
            className="product">
            <div className="product-details">
              <header>{product.stationName}</header>
              <div className="category">{product.year}</div>
            </div>
          </li>
        )
      }
      
      function ProductsList(props) {
        const { products } = props
        
        return (
          <ul className="products">
            {products.map(product => (
              <Product product={product} />
            ))}
          </ul>
        )
      }



      function Testcheckbox() {
        const [state, setState] = useState({
          products: DensityList,
          filters: new Set(),
        })
        
        const handleFilterChange = useCallback(event => {
          setState(previousState => {
            let filters = new Set(previousState.filters)
            let products = DensityList
            console.log("products first")
            console.log(products)
            if (event.target.checked) {
              filters.add(event.target.value)

            } else {
              filters.delete(event.target.value)
            }
            
            if (filters.size) {
              products = products.filter(product => {
                return filters.has(product.stationName)
              })
            }
            console.log("filter")
            console.log(filters)

            console.log("products filtered")
            console.log(products)
            return {
              filters,
              products,
              
            }
          })
        }, [setState])
        
        return (
          <main>
            {/* <DropdownFilter/> */}
            <ProductFilters 
              categories={Stationfilter}
              onFilterChange={handleFilterChange}/>
              <div>
                <div className="mr-2"><Button type="primary" href="/density" >Clear</Button></div>
              </div>
            {/* <ProductsList products={state.products} /> */}
            <DensityBar dataimport={state.products}/>
          </main>
        )
      }

    
  return (
    <>
      <Testcheckbox/>
    </>
  )
}

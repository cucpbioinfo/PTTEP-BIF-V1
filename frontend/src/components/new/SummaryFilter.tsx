import { Select } from 'antd'
import { listAsset } from 'api/dashboard/listAsset'
//import { listYearAsset } from 'api/species/listYearAsset'
import { listPlatform } from 'api/dashboard/listPlatform'
import { listYearPlatform } from 'api/species/listYearPlatform'
import { listStation } from 'api/dashboard/listStation'
import { listYearStation } from 'api/species/listYearStation'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { SPECIES_FILTER_TEXT } from './locale'

const { Option } = Select
export const SummaryFilter = () => {
  const router = useRouter()
  const { locale } = router

  const [asset, setAsset] = useState([])
  //const [assetyear, setAssetYear] = useState([])
  const [platform, setPlatform] = useState([])
  // const [platformyear, setPlatformYear] = useState([])
  const [station, setStation] = useState([])
  const [stationyear, setStationYear] = useState([])
  // const [majorGroups, setMajorGroups] = useState([])
  // const [kingdoms, setKingdoms] = useState([])
  // const [phylums, setPhylums] = useState([])
  // const [orders, setOrders] = useState([])
  // const [classes, setClasses] = useState([])
  // const [families, setFamilies] = useState([])
  // const [genus, setGenus] = useState([])
  // const [AllSpecies,setAllSpecies] = useState([]) ///26june,AllSpecies

  ////26june,AllSpecies-------------------------------------------------
  // const onFilterChangepush = (key: string, value?: string) => {
  //   const url = {
  //     pathname: '/species/[valueid]',
  //     query: {
  //       //...router.query,
  //       valueid: value,
  //     },
  //   }
  //   router.push(url)
  // }
  //-------------------------------------------------------------------


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
  // const fetchAssetDensity = async (speciesId?: string) => {
  //   const { data } = await listAssetDensity(speciesId)
  //   setAssetDensity(data)
  // }
  const fetchAsset = async () => {
    const { data } = await listAsset()
    setAsset(data)
  }
  // const fetchAssetYear = async (assetId?: string) => {
  //     const { data } = await listYearAsset(assetId)
  //     setAssetYear(data)
  //   }
  const fetchPlatform = async (assetId?: string) => {
    const { data } = await listPlatform(assetId)
    setPlatform(data)
  }
  // const fetchPlatformYear = async (platformId?: string) => {
  //     const { data } = await listYearPlatform(platformId)
  //     setPlatformYear(data)
  //   }
  const fetchStation = async (platformId?: string) => {
    const { data } = await listStation(platformId)
    setStation(data)
  }
  const fetchStationYear = async (stationId?: string) => {
      const { data } = await listYearStation(stationId)
      setStationYear(data)
    }

  // const fetchKingdoms = async () => {
  //   const { data } = await listKingdom()
  //   setKingdoms(data)
  // }
  // const fetchPhylums = async (kingdomId?: string) => {
  //   const { data } = await listPhylum(kingdomId)
  //   setPhylums(data)
  // }
  // const fetchClasses = async (phylumId?: string) => {
  //   const { data } = await listClass(phylumId)
  //   setClasses(data)
  // }
  // const fetchOrders = async (classId?: string) => {
  //   const { data } = await listOrder(classId)
  //   setOrders(data)
  // }
  // const fetchFamilies = async (orderId?: string) => {
  //   const { data } = await listFamily(orderId)
  //   setFamilies(data)
  // }
  // const fetchGenus = async (familyId?: string) => {
  //   const { data } = await listGenus(familyId)
  //   setGenus(data)
  // }
  // //26june,AllSpecies
  // const fetchAllSpecies = async (SpeciesId?: string) => {
  //   const { data } = await listAllSpecies(SpeciesId)
  //   setAllSpecies(data)
  // }

  //   fetching filter data
  useEffect(() => {
    fetchAsset()
    // fetchMajorGroups()
    // fetchKingdoms()
  }, [])

  // useEffect(() => {
  //     fetchAssetYear(router.query.assetId as string)
  //   }, [router.query.assetId])

  useEffect(() => {
    fetchPlatform(router.query.assetId as string)
  }, [router.query.assetId])
  // useEffect(() => {
  //   fetchPlatformYear(router.query.platformId as string)
  // }, [router.query.platformId])
  useEffect(() => {
    fetchStation(router.query.platformId as string)
  }, [router.query.platformId])
  useEffect(() => {
    fetchStationYear(router.query.stationId as string)
  }, [router.query.stationId])
  // useEffect(() => {
  //   fetchPhylums(router.query.kingdomId as string)
  // }, [router.query.kingdomId])

  // useEffect(() => {
  //   fetchClasses(router.query.phylumId as string)
  // }, [router.query.phylumId])

  // useEffect(() => {
  //   fetchOrders(router.query.classId as string)
  // }, [router.query.classId])

  // useEffect(() => {
  //   fetchFamilies(router.query.orderId as string)
  // }, [router.query.orderId])

  // useEffect(() => {
  //   fetchGenus(router.query.familyId as string)
  // }, [router.query.familyId])

    // //26june,AllSpecies
    // useEffect(() => {
    //   fetchAllSpecies(router.query.SpeciesId as string)
    // }, [router.query.SpeciesId])

  return (
    <>
      {/* <h2 className="text-2xl font-bold">Filter Header</h2> */}
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
              {asset?.assetName}
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
              {platform?.platformName}
            </Option>
          ))}
        </Select>
      </div>

      <div className="text-sm">
        Year
      </div>
      <div className="md:w-1/6 w-full">
        <Select
            placeholder="year"
            onChange={(value: string | undefined) => {
              onFilterChange('year', value)
            }}
          >
            {stationyear.map((year) => (
              <Option value={year?.year} key={year?.year}>
                {year?.year}
              </Option>
            ))}
        </Select>
      </div>

      {/* <div className="text-sm">
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
              {station?.stationName}
            </Option>
          ))}
        </Select>
      </div> */}

      {/* <div className="mt-4 grid md:grid-cols-7 md:gap-7 grid-cols-2 gap-4">
        <div className="w-full">
          <div className="text-sm">
            Year
          </div>
          <Select
            placeholder="year"
            onChange={(value: string | undefined) => {
              onFilterChange('year', value)
            }}
          >
            {stationyear.map((year) => (
              <Option value={year?.year} key={year?.year}>
                {year?.year}
              </Option>
            ))}
          </Select>
        </div>
      </div> */}
    </>
  )
}

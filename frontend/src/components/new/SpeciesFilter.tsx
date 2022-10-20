import { Select } from 'antd'
import { listClass } from 'api/class/listClass'
import { listFamily } from 'api/family/listFamily'
import { listGenus } from 'api/genus/listGenus'
import { listKingdom } from 'api/kingdom/listKingdom'
import { listMajorGroup } from 'api/major-group/listMajorGroup'
import { listOrder } from 'api/order/listOrder'
import { listPhylum } from 'api/phylum/listPhylum'
import { listAllSpecies } from 'api/species/ListAllSpecies'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { SPECIES_FILTER_TEXT } from './locale'

const { Option } = Select

export const SpeciesFilter = () => {
  const router = useRouter()
  const { locale } = router

  const [majorGroups, setMajorGroups] = useState([])
  const [kingdoms, setKingdoms] = useState([])
  const [phylums, setPhylums] = useState([])
  const [orders, setOrders] = useState([])
  const [classes, setClasses] = useState([])
  const [families, setFamilies] = useState([])
  const [genus, setGenus] = useState([])
  // const [AllSpecies,setAllSpecies] = useState([]) ///26june,AllSpecies

  ////26june,AllSpecies-------------------------------------------------
  const onFilterChangepush = (key: string, value?: string) => {
    const url = {
      pathname: '/species/[valueid]',
      query: {
        //...router.query,
        valueid: value,
      },
    }
    router.push(url)
  }
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
  const fetchMajorGroups = async () => {
    const { data } = await listMajorGroup()
    setMajorGroups(data)
  }
  const fetchKingdoms = async () => {
    const { data } = await listKingdom()
    setKingdoms(data)
  }
  const fetchPhylums = async (kingdomId?: string) => {
    const { data } = await listPhylum(kingdomId)
    setPhylums(data)
  }
  const fetchClasses = async (phylumId?: string) => {
    const { data } = await listClass(phylumId)
    setClasses(data)
  }
  const fetchOrders = async (classId?: string) => {
    const { data } = await listOrder(classId)
    setOrders(data)
  }
  const fetchFamilies = async (orderId?: string) => {
    const { data } = await listFamily(orderId)
    setFamilies(data)
  }
  const fetchGenus = async (familyId?: string) => {
    const { data } = await listGenus(familyId)
    setGenus(data)
  }
  // //26june,AllSpecies
  // const fetchAllSpecies = async (SpeciesId?: string) => {
  //   const { data } = await listAllSpecies(SpeciesId)
  //   setAllSpecies(data)
  // }

  //   fetching filter data
  useEffect(() => {
    fetchMajorGroups()
    fetchKingdoms()
  }, [])

  useEffect(() => {
    fetchPhylums(router.query.kingdomId as string)
  }, [router.query.kingdomId])

  useEffect(() => {
    fetchClasses(router.query.phylumId as string)
  }, [router.query.phylumId])

  useEffect(() => {
    fetchOrders(router.query.classId as string)
  }, [router.query.classId])

  useEffect(() => {
    fetchFamilies(router.query.orderId as string)
  }, [router.query.orderId])

  useEffect(() => {
    fetchGenus(router.query.familyId as string)
  }, [router.query.familyId])

    // //26june,AllSpecies
    // useEffect(() => {
    //   fetchAllSpecies(router.query.SpeciesId as string)
    // }, [router.query.SpeciesId])

  return (
    <>
      <h2 className="text-2xl font-bold">Major group classification</h2>
      <div className="text-sm">
        {SPECIES_FILTER_TEXT[locale].selectMajorGroupLabel}
      </div>
      <div className="md:w-1/3 w-full">
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectMajorGroupPlaceholder}
          onChange={(value: string | undefined) => {
            onFilterChange('majorGroupId', value)
          }}
        >
          {majorGroups.map((majorGroup) => (
            <Option
              value={majorGroup?.majorGroupId}
              key={majorGroup?.majorGroupId}
            >
              {majorGroup?.majorGroupName}
            </Option>
          ))}
        </Select>
      </div>
      <div className="mt-4 grid md:grid-cols-7 md:gap-7 grid-cols-2 gap-4">
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectKingdomLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectKingdomPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChange('kingdomId', value)
            }}
          >
            {kingdoms.map((kingdom) => (
              <Option value={kingdom?.kingdomId} key={kingdom?.kingdomId}>
                {kingdom?.kingdomName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectPhylumLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectPhylumPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChange('phylumId', value)
            }}
          >
            {phylums.map((phylum) => (
              <Option value={phylum?.phylumId} key={phylum?.phylumId}>
                {phylum?.phylumName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectClassLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectClassPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChange('classId', value)
            }}
          >
            {classes.map((_class) => (
              <Option value={_class?.classId} key={_class?.classId}>
                {_class?.className}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectOrderLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectOrderPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChange('orderId', value)
            }}
          >
            {orders.map((order) => (
              <Option value={order?.orderId} key={order?.orderId}>
                {order?.orderName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectFamilyLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectFamilyPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChange('familyId', value)
            }}
          >
            {families.map((family) => (
              <Option value={family?.familyId} key={family?.familyId}>
                {family?.familyName}
              </Option>
            ))}
          </Select>
        </div>
        <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectGenusLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectGenusPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChange('genusId', value)
            }}
          >
            {genus.map((_genus) => (
              <Option value={_genus?.genusId} key={_genus?.genusId}>
                {_genus?.genusName}
              </Option>
            ))}
          </Select>
        </div>
        {/* //26june,AllSpecies */}
        {/* <div className="w-full">
          <div className="text-sm">
            {SPECIES_FILTER_TEXT[locale].selectAllSpeciesLabel}
          </div>
          <Select
            placeholder={SPECIES_FILTER_TEXT[locale].selectAllSpeciesPlaceholder}
            onChange={(value: string | undefined) => {
              onFilterChangepush('SpeciesId', value)
            }}
          >
            {AllSpecies.map((AllSpecies) => (
              <Option value={AllSpecies?.speciesId} key={AllSpecies?.speciesId}>
                {AllSpecies?.speciesName}
              </Option>
            ))}
          </Select>
        </div> */}
      </div>
    </>
  )
}

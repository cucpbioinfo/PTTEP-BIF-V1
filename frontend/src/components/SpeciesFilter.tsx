import { Select } from 'antd'
import { listClass } from 'api/class/listClass'
import { listFamily } from 'api/family/listFamily'
import { listGenus } from 'api/genus/listGenus'
import { listKingdom } from 'api/kingdom/listKingdom'
import { listOrder } from 'api/order/listOrder'
import { listPhylum } from 'api/phylum/listPhylum'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { SPECIES_FILTER_TEXT } from './locale'
const { Option } = Select

interface SpeciesFilterProps {
  onFilter?: (selectedFilter?: ISelectedFilter) => void
}

export interface ISelectedFilter {
  kingdomId?: string
  phylumId?: string
  orderId?: string
  classId?: string
  familyId?: string
  genusId?: string
}

export const SpeciesFilter = ({ onFilter }: SpeciesFilterProps) => {
  const { locale } = useRouter()
  const [kingdoms, setKingdoms] = useState([])
  const [phylums, setPhylums] = useState([])
  const [orders, setOrders] = useState([])
  const [classes, setClasses] = useState([])
  const [families, setFamilies] = useState([])
  const [genus, setGenus] = useState([])
  const [selectedFilter, setSelectedFilter] = useState<ISelectedFilter>({})

  const fetchKingdoms = async () => {
    const { data } = await listKingdom()
    setKingdoms(data)
  }
  const fetchPhylums = async () => {
    const { data } = await listPhylum(selectedFilter.kingdomId)
    setPhylums(data)
  }
  const fetchClasses = async () => {
    const { data } = await listClass(selectedFilter.phylumId)
    setClasses(data)
  }
  const fetchOrders = async () => {
    const { data } = await listOrder(selectedFilter.classId)
    setOrders(data)
  }
  const fetchFamilies = async () => {
    const { data } = await listFamily(selectedFilter.orderId)
    setFamilies(data)
  }
  const fetchGenus = async () => {
    const { data } = await listGenus(selectedFilter.familyId)
    setGenus(data)
  }
  useEffect(() => {
    onFilter(selectedFilter)
  }, [selectedFilter])

  useEffect(() => {
    fetchKingdoms()
  }, [])

  useEffect(() => {
    fetchPhylums()
  }, [selectedFilter.kingdomId])
  useEffect(() => {
    fetchClasses()
  }, [selectedFilter.phylumId])
  useEffect(() => {
    fetchOrders()
  }, [selectedFilter.classId])
  useEffect(() => {
    fetchFamilies()
  }, [selectedFilter.orderId])
  useEffect(() => {
    fetchGenus()
  }, [selectedFilter.familyId])

  return (
    <div className="grid md:grid-cols-6 md:gap-6 grid-cols-2 gap-4">
      <div className="w-full">
        <div className="text-sm">
          {SPECIES_FILTER_TEXT[locale].selectKingdomLabel}
        </div>
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectKingdomPlaceholder}
          onChange={(value) => {
            setSelectedFilter((prev) => ({
              ...prev,
              kingdomId: value as string,
            }))
          }}
        >
          {kingdoms.map((kingdom) => (
            <Option value={kingdom?.kingdomId} key={kingdom?.kingdomId}>
              {kingdom?.kingdomName}
            </Option>
          ))}
        </Select>
      </div>
      <div>
        <div className="text-sm">
          {SPECIES_FILTER_TEXT[locale].selectPhylumLabel}
        </div>
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectPhylumPlaceholder}
          onChange={(value) => {
            setSelectedFilter((prev) => ({
              ...prev,
              phylumId: value as string,
            }))
          }}
        >
          {phylums.map((phylum) => (
            <Option value={phylum?.phylumId} key={phylum?.phylumId}>
              {phylum?.phylumName}
            </Option>
          ))}
        </Select>
      </div>
      <div>
        <div className="text-sm">
          {SPECIES_FILTER_TEXT[locale].selectClassLabel}
        </div>
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectClassPlaceholder}
          onChange={(value) => {
            setSelectedFilter((prev) => ({
              ...prev,
              classId: value as string,
            }))
          }}
        >
          {classes.map((_class) => (
            <Option value={_class?.classId} key={_class?.classId}>
              {_class?.className}
            </Option>
          ))}
        </Select>
      </div>
      <div>
        <div className="text-sm">
          {SPECIES_FILTER_TEXT[locale].selectOrderLabel}
        </div>
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectOrderPlaceholder}
          onChange={(value) => {
            setSelectedFilter((prev) => ({
              ...prev,
              orderId: value as string,
            }))
          }}
        >
          {orders.map((order) => (
            <Option value={order?.orderId} key={order?.orderId}>
              {order?.orderName}
            </Option>
          ))}
        </Select>
      </div>
      <div>
        <div className="text-sm">
          {SPECIES_FILTER_TEXT[locale].selectFamilyLabel}
        </div>
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectFamilyPlaceholder}
          onChange={(value) => {
            setSelectedFilter((prev) => ({
              ...prev,
              familyId: value as string,
            }))
          }}
        >
          {families.map((family) => (
            <Option value={family?.familyId} key={family?.familyId}>
              {family?.familyName}
            </Option>
          ))}
        </Select>
      </div>
      <div>
        <div className="text-sm">
          {SPECIES_FILTER_TEXT[locale].selectGenusLabel}
        </div>
        <Select
          placeholder={SPECIES_FILTER_TEXT[locale].selectGenusPlaceholder}
          onChange={(value) => {
            setSelectedFilter((prev) => ({
              ...prev,
              genusId: value as string,
            }))
          }}
        >
          {genus.map((_genus) => (
            <Option value={_genus?.genusId} key={_genus?.genusId}>
              {_genus?.genusName}
            </Option>
          ))}
        </Select>
      </div>
    </div>
  )
}

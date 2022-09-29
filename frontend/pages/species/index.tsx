import { Divider, Empty, Pagination, Select } from 'antd'
import { listSpecies } from 'api/species/listSpecies'
import { listDensities } from 'api/density/listDensities'
import { PageLayout } from 'components/new/PageLayout'
import { SpeciesCard } from 'components/new/SpeciesCard'
import { SpeciesFilter } from 'components/new/SpeciesFilter'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import { MOCK_SPECIES } from 'utils/mockData'

const species = () => {
  const router = useRouter()
  const [species, setSpecies] = useState([])
  const [total, setTotal] = useState(0)

  const [pageNumber, setPageNumber] = useState(1)
  const [pageSize, setPageSize] = useState(10)

  const onPaginationChanged = (page: number, size: number) => {
    setPageNumber(page)
    setPageSize(size)
  }

  const fetchSpecies = async () => {
    const {
      majorGroupId,
      kingdomId,
      phylumId,
      classId,
      orderId,
      familyId,
      genusId,
      keyword,
    } = router.query
    const { data, totalSpecies } = await listSpecies({
      majorGroupId: majorGroupId as string,
      kingdomId: kingdomId as string,
      phylumId: phylumId as string,
      classId: classId as string,
      orderId: orderId as string,
      familyId: familyId as string,
      genusId: genusId as string,
      keyword: keyword as string,
      pageSize,
      pageNumber,
    })
    setSpecies(data)
    setTotal(totalSpecies)
  }

  useEffect(() => {
    fetchSpecies()
  }, [router.query, pageNumber, pageSize])

  return (
    <PageLayout>
      <SpeciesFilter />
      <Divider />
      <div>{species.length === 0 && <Empty />}</div>
      <div className="mt-8 grid md:grid-cols-4 md:gap-8 grid-cols-1 gap-y-4">
        {species.length !== 0 &&
          species.map(({ speciesId, familyName, genusName, speciesName }) => (
            <SpeciesCard
              speciesName={speciesName}
              genusName={genusName}
              familyName={familyName}
              onClick={() => router.push(`/species/${speciesId}`)}
              key={speciesId}
            />
          ))}
      </div>
      <div className="mt-16 flex justify-center">
        <Pagination
          defaultCurrent={1}
          defaultPageSize={10}
          pageSize={pageSize}
          total={total}
          onChange={onPaginationChanged}
          hideOnSinglePage={true}
        />
      </div>
    </PageLayout>
  )
}

export default species

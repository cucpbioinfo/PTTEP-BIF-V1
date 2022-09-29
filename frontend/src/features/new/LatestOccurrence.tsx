import { Divider, Pagination } from 'antd'
import { getLatestOccurrences } from 'api/occurrences/getLatestOccurrences'
import { LatestOccurenceCard } from 'components/new/LatestOccurenceCard'
import React, { useEffect, useState } from 'react'
import { PaginationQuery } from 'types/pagination'

export const LatestOccurrence = () => {
  const [latestOccurrences, setLatestOccurrences] = useState([])
  const [pageNumber, setPageNumber] = useState(1)
  const [pageSize, setPageSize] = useState(3)
  const [total, setTotal] = useState(0)

  const onPageNumberChange = (page: number) => {
    setPageNumber(page)
  }

  const fetchLatestOccurrences = async ({
    pageNumber,
    pageSize,
  }: PaginationQuery) => {
    const { occurrences, total } = await getLatestOccurrences({
      pageNumber,
      pageSize,
    })
    setLatestOccurrences(occurrences)
    setTotal(total)
  }

  useEffect(() => {
    console.log({ pageNumber, pageSize })
    fetchLatestOccurrences({ pageNumber, pageSize })
  }, [pageNumber, pageSize])

  return (
    <>
      <div className="md:py-12 py-4 md:px-32 px-4 flex justify-center items-center space-x-8 md:text-3xl text-xl">
        <h2>Occurrence records</h2>
        <h2 className="md:text-4xl text-2xl text-primary">{total}</h2>
        <h2>Results</h2>
      </div>
      <Divider className="m-0" />
      <h2 className="md:text-3xl text-2xl md:mt-0 mt-8 md:mb-8 mb-4">
        Latest Occurrence
      </h2>
      <div className="grid grid-rows gap-8">
        {latestOccurrences.map((occurrence) => (
          <LatestOccurenceCard
            imageSrc={occurrence.imageSrc}
            occurrenceDetails={occurrence.occurrenceDetails}
            discoveredAt={occurrence.discoveredAt}
          />
        ))}
      </div>
      <div className="mt-16 flex justify-center">
        <Pagination
          current={pageNumber}
          onChange={onPageNumberChange}
          total={total}
        />
      </div>
    </>
  )
}

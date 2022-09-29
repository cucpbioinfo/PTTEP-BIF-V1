import { eDNAAxios } from 'api/const'
import { PaginationQuery } from 'types/pagination'
import { MOCK_LATEST_OCCURRENCES } from 'utils/mockData'

export const getLatestOccurrences = async ({
  pageNumber,
  pageSize,
}: PaginationQuery) => {
  const { data, total } = await (await eDNAAxios.get('/occurrences')).data

  return { occurrences: data, total }
}

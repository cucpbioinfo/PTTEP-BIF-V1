import { Select } from 'antd'
import { listMajorGroup } from 'api/major-group/listMajorGroup'
import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'

const { Option } = Select
export const SummaryFilterGroup = () => {
  const router = useRouter()
  const { locale } = router
  const [majorGroups, setMajorGroups] = useState([])

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

  useEffect(() => {
    fetchMajorGroups()
  }, [])

  return (
    <>
      <div className="text-sm">
        Major Group
      </div>
      <div className="md:w-auto w-auto">
        <Select
          placeholder="Major Group"
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
    </>
  )
}

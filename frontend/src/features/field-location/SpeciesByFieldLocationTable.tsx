import { Table } from 'antd'
import React from 'react'

export const SpeciesByFieldLocationTable = () => {
  const fieldLocations = [
    'APP-1A',
    'APP-2A',
    'APP-3A',
    'APP-4A',
    'APP-1B',
    'APP-2B',
  ]
  const fieldLocationColumns = fieldLocations.map((fieldLocation) => ({
    title: fieldLocation,
    children: [
      {
        title: '1',
        dataIndex: `${fieldLocation}-1`,
        key: `${fieldLocation}-1`,
      },
      {
        title: '2',
        dataIndex: `${fieldLocation}-2`,
        key: `${fieldLocation}-2`,
      },
      {
        title: '3',
        dataIndex: `${fieldLocation}-3`,
        key: `${fieldLocation}-3`,
      },
    ],
  }))
  const columns = [
    {
      title: 'ไฟลัม',
      dataIndex: 'phylumName',
      key: 'phylumName',
    },
    ...fieldLocationColumns,
  ]
  const data = [
    {
      phylumName: 'Annelida',
      'APP-1A-1': 1,
      'APP-1A-2': 1,
      'APP-1A-3': 1,
      'APP-2A-1': 1,
      'APP-2A-2': 1,
      'APP-2A-3': 1,
      'APP-3A-1': 1,
      'APP-3A-2': 1,
      'APP-3A-3': 1,
      'APP-4A-1': 1,
      'APP-4A-2': 1,
      'APP-4A-3': 1,
      'APP-1B-1': 1,
      'APP-1B-2': 1,
      'APP-1B-3': 1,
      'APP-2B-1': 1,
      'APP-2B-2': 1,
      'APP-2B-3': 1,
    },
    {
      phylumName: 'Annelida2',
      'APP-1A-1': 1,
      'APP-1A-2': 1,
      'APP-1A-3': 1,
      'APP-2A-1': 1,
      'APP-2A-2': 1,
      'APP-2A-3': 1,
      'APP-3A-1': 1,
      'APP-3A-2': 1,
      'APP-3A-3': 1,
      'APP-4A-1': 1,
      'APP-4A-2': 1,
      'APP-4A-3': 1,
      'APP-1B-1': 1,
      'APP-1B-2': 1,
      'APP-1B-3': 1,
      'APP-2B-1': 1,
      'APP-2B-2': 1,
      'APP-2B-3': 1,
    },
  ]
  return (
    <div>
      <Table
        columns={columns}
        dataSource={data}
        scroll={{ x: true }}
        rowClassName="cursor-pointer"
        rowKey={(record) => record.phylumName}
      ></Table>
    </div>
  )
}

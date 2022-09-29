// @ts-nocheck
import { Table } from 'antd'
import { listSite } from 'api/site/listSite'
import React, { useEffect, useState } from 'react'
import { Site } from 'types/site'
import { MOCK_SPECIES_FIELD_LOCATIONS } from 'utils/mockData'

const columns = [
  {
    title: 'ปริมาณความหนาแน่น (ตัว / ตร.ม.)',
    align: 'center',
    children: [
      {
        title: 'สถานี',
        dataIndex: 'station',
        key: 'station',
        align: 'center',
      },
      {
        title: 'APP-1B',
        dataIndex: 'app-1b',
        key: 'app-1b',
        children: [
          {
            title: '(1)',
            dataIndex: 'app-1b-1',
            key: 'app-1b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'app-1b-2',
            key: 'app-1b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'APP-3B',
        dataIndex: 'app-3b',
        key: 'app-3b',
        children: [
          {
            title: '(1)',
            dataIndex: 'app-3b-1',
            key: 'app-3b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'app-3b-2',
            key: 'app-3b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'AWP8-1B',
        dataIndex: 'awp8-1b',
        key: 'awp8-1b',
        children: [
          {
            title: '(1)',
            dataIndex: 'awp8-1b-1',
            key: 'awp8-1b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'awp8-1b-2',
            key: 'awp8-1b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'AWP8-3B',
        dataIndex: 'awp8-3b',
        key: 'awp8-3b',
        children: [
          {
            title: '(1)',
            dataIndex: 'awp8-3b-1',
            key: 'awp8-3b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'awp8-3b-2',
            key: 'awp8-3b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'AWP-1N-1B',
        dataIndex: 'awp-1n-1b',
        key: 'awp-1n-1b',
        children: [
          {
            title: '(1)',
            dataIndex: 'awp-1n-1b-1',
            key: 'awp-1n-1b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'awp-1n-1b-2',
            key: 'awp-1n-1b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'AWP-1N-3B',
        dataIndex: 'awp-1n-3b',
        key: 'awp-1n-3b',
        children: [
          {
            title: '(1)',
            dataIndex: 'awp-1n-3b-1',
            key: 'awp-1n-3b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'awp-1n-3b-2',
            key: 'awp-1n-3b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'AWP29-1B',
        dataIndex: 'awp29-1b',
        key: 'awp29-1b',
        children: [
          {
            title: '(1)',
            dataIndex: 'awp29-1b-1',
            key: 'awp29-1b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'awp29-1b-2',
            key: 'awp29-1b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
      {
        title: 'AWP29-3B',
        dataIndex: 'awp29-3b',
        key: 'awp29-3b',
        children: [
          {
            title: '(1)',
            dataIndex: 'awp29-3b-1',
            key: 'awp29-3b-1',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
          {
            title: '(2)',
            dataIndex: 'awp29-3b-2',
            key: 'awp29-3b-2',
            align: 'center',
            render: (text) => <p>{text?.toLocaleString()}</p>,
          },
        ],
      },
    ],
  },
]

export const LocationTable = () => {
  const [sites, setSites] = useState<Site[]>([])

  const fetchSites = async () => {
    const data = await listSite()
    console.log(data)
    setSites(data)
  }

  useEffect(() => {
    fetchSites()
  }, [])

  return (
    <Table
      columns={columns}
      bordered
      dataSource={MOCK_SPECIES_FIELD_LOCATIONS}
      scroll={{ x: true }}
    ></Table>
  )
}

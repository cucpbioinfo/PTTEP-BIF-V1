import React from 'react'
import Image from 'next/image'
import dayjs from 'dayjs'
import { Divider } from 'antd'

interface LatestOccurenceCardProps {
  imageSrc?: string
  occurrenceDetails: string
  discoveredAt: Date
}

export const LatestOccurenceCard = ({
  imageSrc,
  occurrenceDetails,
  discoveredAt,
}: LatestOccurenceCardProps) => {
  const parseDiscoveredAt = (discoveredAt: Date) => {
    return dayjs(discoveredAt).format('MMMM YYYY')
  }
  return (
    <div className="grid md:grid-cols-2 gap-4 grid-cols-1">
      <div className="m-auto">
        <div className="grid justify-items-end text-lg">
          {parseDiscoveredAt(discoveredAt)}
        </div>
        <Image
          src={imageSrc ? imageSrc : 'https://picsum.photos/300/150'}
          width={400}
          height={300}
        />
      </div>
      <div className="flex flex-col justify-between md:pt-12 pt-6">
        <p className="md:text-xl text-lg">{occurrenceDetails}</p>
        <div className="grid justify-items-end">
          <a className="text-lg text-primary">Read More</a>
        </div>
      </div>
      <Divider />
    </div>
  )
}

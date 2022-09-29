import dayjs from 'dayjs'
import { useRouter } from 'next/router'
import React from 'react'

interface InternalDataInteProps {
  title: string
  href: string
  updatedAt: Date
}

export const InternalDataItem = ({
  title,
  href,
  updatedAt,
}: InternalDataInteProps) => {
  const router = useRouter()
  return (
    <div onClick={() => router.push(href)} className="cursor-pointer">
      <h2 className="font-bold text-primary text-xl md:text-2xl">{title}</h2>
      <h5 className="font-normal text-gray-900 text:lg md:text-xl">
        Modified At {dayjs(updatedAt).format('MMMM YYYY')}
      </h5>
    </div>
  )
}

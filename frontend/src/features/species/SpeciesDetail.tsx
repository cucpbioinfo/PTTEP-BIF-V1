import { Descriptions } from 'antd'
import { useRouter } from 'next/router'
import React from 'react'
import { SpeciesDetailsResponse } from 'types/species'
import { SPECIES_DETAILS_KEYS } from './locale'

export const SpeciesDetail = ({
  speciesDetail,
}: {
  speciesDetail: SpeciesDetailsResponse
}) => {
  const router = useRouter()
  const {
    speciesName,
    majorGroupId,
    majorGroupName,
    kingdomName,
    phylumName,
    className,
    orderName,
    familyName,
    genusName,
    kingdomId, //26June Quicklink
    phylumId,
    classId,
    orderId,
    familyId,
    genusId,
  } = speciesDetail
  return (
    <Descriptions bordered column={1} size="default">
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].majorGroupName}>
      <a onClick={() => router.push(`/species?majorGroupId=${majorGroupId}`)}>{majorGroupName || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].kingdomName}>
        <a onClick={() => router.push(`/species?kingdomId=${kingdomId}`)}>{kingdomName || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].phylumName}>
        <a onClick={() => router.push(`/species?phylumId=${phylumId}`)}>{phylumName || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].className}>
        <a onClick={() => router.push(`/species?classId=${classId}`)}>{className || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].orderName}>
        <a onClick={() => router.push(`/species?orderId=${orderId}`)}>{orderName || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].familyName}>
        <a onClick={() => router.push(`/species?familyId=${familyId}`)}>{familyName || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].genusName}>
        <i><a onClick={() => router.push(`/species?genusId=${genusId}`)}>{genusName || "-"}</a></i>
      </Descriptions.Item>
      <Descriptions.Item
        label={SPECIES_DETAILS_KEYS[router.locale].speciesName}
      >
        <i>{speciesName}</i>
      </Descriptions.Item>
    </Descriptions>
  )
}

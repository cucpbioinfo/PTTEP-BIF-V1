import { Descriptions } from 'antd'
import { useRouter } from 'next/router'
import React from 'react'
import { SpeciesDetailsResponse } from 'types/species'
import { SPECIES_DETAILS_KEYS } from './locale'

function capitalizeFirstLetter(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

function SpeciesSplice(string:string) {
  if(string.search(/sp./i)==-1){
    let touppercase = string.charAt(0).toUpperCase() + string.slice(1);
    return <><i>{touppercase}</i></>;
  }
  else
  {
    let position = string.search(/sp./i);
    let speciesName = string.substring(0, position);
    let upperspeciesName = speciesName.charAt(0).toUpperCase() + speciesName.slice(1);
    let SpecialText = string.slice(position);
    return <><i>{upperspeciesName}</i> {SpecialText}</>;
  }
  }

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
      <a onClick={() => router.push(`/species?majorGroupId=${majorGroupId}`)}>{capitalizeFirstLetter(majorGroupName) || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].kingdomName}>
        <a onClick={() => router.push(`/species?kingdomId=${kingdomId}`)}>{capitalizeFirstLetter(kingdomName) || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].phylumName}>
        <a onClick={() => router.push(`/species?phylumId=${phylumId}`)}>{capitalizeFirstLetter(phylumName) || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].className}>
        <a onClick={() => router.push(`/species?classId=${classId}`)}>{capitalizeFirstLetter(className) || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].orderName}>
        <a onClick={() => router.push(`/species?orderId=${orderId}`)}>{capitalizeFirstLetter(orderName) || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].familyName}>
        <a onClick={() => router.push(`/species?familyId=${familyId}`)}>{capitalizeFirstLetter(familyName) || "-"}</a>
      </Descriptions.Item>
      <Descriptions.Item label={SPECIES_DETAILS_KEYS[router.locale].genusName}>
        <i><a onClick={() => router.push(`/species?genusId=${genusId}`)}>{capitalizeFirstLetter(genusName) || "-"}</a></i>
      </Descriptions.Item>
      <Descriptions.Item
        label={SPECIES_DETAILS_KEYS[router.locale].speciesName}
      >
        {SpeciesSplice(speciesName)}
      </Descriptions.Item>
    </Descriptions>
  )
}

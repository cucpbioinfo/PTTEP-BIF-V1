import React from 'react'
import { SpeciesDetailsResponse } from 'types/species'
import { SpeciesImages } from './speciesimages' 

export const SpeciesOverviewTab = ({
  speciesDetail,
}: {
  speciesDetail: SpeciesDetailsResponse
}) => {
  const {
    identificationName,
    methodName,
  } = speciesDetail
  return (
    <div className="py-6">
      <SpeciesImages />
      <br/>
      <div className="mb-6">
        <div className="text-xl font-semibold">Identification Technique</div>
        <div>{identificationName}</div>
      </div>
      <div className="mb-6">
        <div className="text-xl font-semibold">Method of Collection</div>
        <div>{methodName}</div>
      </div>
    </div>
  )
}

import React from 'react'
import { Tooltip } from 'antd'
import { InfoCircleOutlined } from '@ant-design/icons';
import { SpeciesDetailsResponse } from 'types/species'
import { SpeciesImages } from './speciesimages' 

const IdentificationInfo = <span>Definition of Identifacation Technique.</span>;
const MethodInfo = <span>Definition of about Method of Collection.</span>;
const AccessionInfo = <span>Definition of about Genbank.</span>;

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
        <div className="text-xl font-semibold">
          Identification Technique <Tooltip placement="top" title={IdentificationInfo}><InfoCircleOutlined/></Tooltip>
        </div>
        <div>{identificationName}</div>
      </div>
      <div className="mb-6">
        <div className="text-xl font-semibold">
          Method of Collection <Tooltip placement="top" title={MethodInfo}><InfoCircleOutlined/></Tooltip>
        </div>
        <div>{methodName}</div>
      </div>
      <div className="mb-6">
        <div className="text-xl font-semibold">
          Accession <Tooltip placement="top" title={AccessionInfo}><InfoCircleOutlined/></Tooltip>
        </div>
        <div>number@genbank : </div>
      </div>
      
    </div>
  )
}

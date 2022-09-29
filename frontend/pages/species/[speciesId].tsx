import { SpeciesDetailPage } from 'features/species/SpeciesDetailPage'
import { useRouter } from 'next/router'
import React from 'react'

const SpeciesDetails = () => {
  const router = useRouter()
  const { speciesId } = router.query
  return (
    <div>
      <SpeciesDetailPage speciesId={speciesId} />
    </div>
  )
}

export default SpeciesDetails

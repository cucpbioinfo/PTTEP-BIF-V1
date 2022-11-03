import React from 'react'

interface SpeciesCardProps {
  speciesName: string
  genusName: string
  familyName: string
  onClick?: () => void
}

export const SpeciesCard = ({
  speciesName,
  genusName,
  familyName,
  onClick = () => {},
}: SpeciesCardProps) => {
  return (
    <div
      className="w-full border shadow px-4 py-4 cursor-pointer"
      onClick={onClick}
    >
      <img
        src="species-images/Bacteriastrum_001.jpg"
        className="w-full object-cover"
      />
      <div className="mt-4">
        <h3 className="text-xl">Species Name: {speciesName}</h3>
        <h2 className="text-lg">Genus: {genusName}</h2>
        <h2 className="text-lg">Family: {familyName}</h2>
      </div>
    </div>
  )
}

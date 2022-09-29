import React from 'react'

// interface Speciesprops {
//     SpeciesId: string
//     SpeciesName: string
  
// }

// export const SpeciesImages = ({SpeciesId,SpeciesName}:Speciesprops) => {
export const SpeciesImages = () => {
  
  return (
    <div className="flex">
        <div className="w-full border py-4 ">
            <img
                src="../species-images/Bacteriastrum_001.jpg"
                className="w-full object-cover"/>
        </div>
        <div className="w-full border py-4 ">
            <img
                src="../species-images/Bacteriastrum_002.jpg"
                className="w-full object-cover"/>
        </div>
        <div className="w-full border py-4 ">
            <img
                src="../species-images/Bacteriastrum_003.jpg"
                className="w-full object-cover"/>
        </div>
        <div className="w-full border py-4 ">
            <img
                src="../species-images/Bacteriastrum_004.jpg"
                className="w-full object-cover"/>
        </div>
    </div>
  )
}

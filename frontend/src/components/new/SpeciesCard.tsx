import React from 'react'

interface SpeciesCardProps {
  speciesName: string
  genusName: string
  familyName: string
  onClick?: () => void
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
function ToUpperSplice(string:string) {
    let touppercase = string.charAt(0).toUpperCase() + string.slice(1);
    return <>{touppercase}</>;}

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
        <h3 className="text-xl">Scientific Name : {SpeciesSplice(speciesName)}</h3>
        <h2 className="text-lg">Genus : {SpeciesSplice(genusName)}</h2>
        <h2 className="text-lg">Family : {ToUpperSplice(familyName)}</h2>
      </div>
    </div>
  )
}

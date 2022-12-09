import React from 'react'
import { InfoCircleOutlined } from '@ant-design/icons';
import { Tooltip} from "antd";
// const DiversityDefInfo = <>
// <p><b>The Shannon-Weiner Species Diversity Index</b> is an index that is commonly used to characterize species diversity in a community. 
// This index accounts for both abundance and evenness of the species present. A large number of species can increase diversity. 
// Similarly, increasing the uniformity of individual distribution among species will also increase diversity. If each individual 
// belongs to a different species, the diversity index is the largest. In contrast, if there is/are dominant species (1 or more species), 
// the diversity index will be decreased.</p>
// <br/>
// <p><b>The Shannon-Weiner Species Diversity Index</b> is calculated by taking the number of each species, the proportion each species 
// is of the total number of individuals, and sums the proportion times the natural log of the proportion for each species. 
// The formula is as follows
// <img className="" src="/def/DiversityIndex.png"></img>
// Where H’ is the species diversity index, 
// s is the number of species
// pi is the proportion of individuals of each species belonging to the ith species of the total number of individuals.</p>
// <br/>
// <p><b>References</b>
// <a href="https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index"> https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index </a>
// Nolan, K.A. and J.E. Callahan. 2006. Beachcomber biology: The Shannon-Weiner Species Diversity Index. 
// Pages 334-338, in Tested Studies for Laboratory Teaching, Volume 27 
// (M.A. O'Donnell, Editor). Proceedings of the 27th Workshop/Conference of the Association for Biology Laboratory Education (ABLE), 383 pages.</p>
// </>;

// const EvennessDefInfo = <>
// <p><b>Pielou’s evenness index</b> was proposed by Pielou (1966). This index refers to how close in numbers each species in an environment is. The Pielou index J′ is defined as

// <img className="" src="/def/EvennessIndex.png"></img>

// Where J' is constrained between 0 and 1. 
// H’ is the Shannon species diversity index
// S is the total number of species.
// If all species are represented in equal numbers in the sample, then J′=1. If one species strongly dominates J′ is close to zero.</p>

// <br/>
// <p><b>References</b>
// <a href="https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index"> https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index </a>
// <a href="https://en.wikipedia.org/wiki/Species_evenness"> https://en.wikipedia.org/wiki/Species_evenness </a>
// </p>
// </>;
export const Diversityinfo = () => {
  const DivInfo = <>
    <p><b>The Shannon-Weiner Species Diversity Index</b> is an index that is commonly used to characterize species diversity in a community. 
      This index accounts for both abundance and evenness of the species present. <a href="">Read more...</a></p> 
  </>;
  return(
    <>
      <Tooltip placement="top" title={DivInfo}><InfoCircleOutlined/></Tooltip>
    </>
  )
}

export const DefDiversity = () => {
  return (
    <>
      <p><b>The Shannon-Weiner Species Diversity Index</b> is an index that is commonly used to characterize species diversity in a community. 
      This index accounts for both abundance and evenness of the species present. A large number of species can increase diversity. 
      Similarly, increasing the uniformity of individual distribution among species will also increase diversity. If each individual 
      belongs to a different species, the diversity index is the largest. In contrast, if there is/are dominant species (1 or more species), 
      the diversity index will be decreased.</p>
      <br/>
      <p><b>The Shannon-Weiner Species Diversity Index</b> is calculated by taking the number of each species, the proportion each species 
      is of the total number of individuals, and sums the proportion times the natural log of the proportion for each species. 
      The formula is as follows
      <img className="" src="/def/DiversityIndex.png"></img>
      Where H’ is the species diversity index, 
      s is the number of species
      pi is the proportion of individuals of each species belonging to the ith species of the total number of individuals.</p>
      <br/>
      <p><b>References</b></p>
      <p><a href="https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index"> https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index </a></p>
      <p>Nolan, K.A. and J.E. Callahan. 2006. Beachcomber biology: The Shannon-Weiner Species Diversity Index. 
      Pages 334-338, in Tested Studies for Laboratory Teaching, Volume 27 
      (M.A. O'Donnell, Editor). Proceedings of the 27th Workshop/Conference of the Association for Biology Laboratory Education (ABLE), 383 pages.</p>
    </>
  )
}

export const DefEvenness = () => {
  return (
    <>
      <p><b>Pielou’s evenness index</b> was proposed by Pielou (1966). This index refers to how close in numbers each species in an environment is. The Pielou index J′ is defined as

      <img className="" src="/def/EvennessIndex.png"></img>

      Where J' is constrained between 0 and 1. 
      H’ is the Shannon species diversity index
      S is the total number of species.
      If all species are represented in equal numbers in the sample, then J′=1. If one species strongly dominates J′ is close to zero.</p>

      <br/>
      <p><b>References</b><p/>
      <p><a href="https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index"> https://www.marinespecies.org/introduced/wiki/Measurements_of_biodiversity#Shannon-Wiener_diversity_index </a></p>
      <p><a href="https://en.wikipedia.org/wiki/Species_evenness"> https://en.wikipedia.org/wiki/Species_evenness </a></p>
      </p>
    </>
  )
}

export const DefNumber = () => {
  return (
    <>
    <p><b>Numbers of species</b> is a counted number of total species found in each area/study station.</p>
    </>
  )
}

export const DefSam = () => {
  return (
    <>
    <p><b>Surface depth</b> 1 meter below sea surface</p>
    <p><b>Euphotic depth</b> Euphotic zone depth, reflects the depth where photosynthetic available radiation (PAR) is 1% of its surface value. The depth of the disappearance of the Secchi disk multiplied by a factor of ~2.8 yields the depth of the 1% light level or the euphotic zone depth.</p>
    <br/>
    <p><b>References</b></p>
    <p>Lee, Z., A. Weidemann, J. Kindle, R. Arnone, K. L. Carder, and C. Davis (2007), Euphotic zone depth: Its derivation andimplication to ocean-color remote sensing,J. Geophys. Res.,112, C03009, doi:10.1029/2006JC003802.</p>
    </>
  )
}

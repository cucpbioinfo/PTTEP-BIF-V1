import React from 'react';

const SummaryData = [
  {
      summaryId: "a93f47ab-ce0b-44b4-8059-e591cf2db8ed",
      year: 2022,
      majorGroupName: "phytoplankton",
      identificationName: "morphology",
      assetName: "gbs",
      platformName: "pps",
      stationName: "pps1",
      surfaceShannon: 3.09,
      surfaceNumber: 47,
      surfaceMax: 3.85,
      surfaceEvenness: 0.8,
      euphoticzoneShannon: 2.83,
      euphoticzoneNumber: 53,
      euphoticzoneMax: 3.97,
      euphoticzoneEvenness: 0.71,
      averageShannon: 3.04,
      averageNumber: 58,
      averageMax: 4.06,
      averageEvenness: 0.75
  },
  {
      summaryId: "925cfc7c-a2ac-483c-8e7e-0870fe160d83",
      year: 2022,
      majorGroupName: "phytoplankton",
      identificationName: "morphology",
      assetName: "gbs",
      platformName: "pps",
      stationName: "pps2",
      surfaceShannon: 2.88,
      surfaceNumber: 33,
      surfaceMax: 3.5,
      surfaceEvenness: 0.82,
      euphoticzoneShannon: 2.8,
      euphoticzoneNumber: 49,
      euphoticzoneMax: 3.89,
      euphoticzoneEvenness: 0.72,
      averageShannon: 3,
      averageNumber: 55,
      averageMax: 4.01,
      averageEvenness: 0.75
  },
  {
      summaryId: "381350b9-0244-4514-9398-48dec9697c64",
      year: 2022,
      majorGroupName: "phytoplankton",
      identificationName: "morphology",
      assetName: "gbs",
      platformName: "pps",
      stationName: "pps3",
      surfaceShannon: 2.79,
      surfaceNumber: 34,
      surfaceMax: 3.53,
      surfaceEvenness: 0.79,
      euphoticzoneShannon: 3.18,
      euphoticzoneNumber: 38,
      euphoticzoneMax: 3.64,
      euphoticzoneEvenness: 0.87,
      averageShannon: 3.14,
      averageNumber: 46,
      averageMax: 3.83,
      averageEvenness: 0.82
  }
];

export const FilterData = () => {
 
  return (
    <div>
      {SummaryData.filter(summary => summary.year == 2022 && summary.stationName == "pps3").
      map(filteredData => (
        <li>
          {filteredData.surfaceShannon}
          {filteredData.surfaceNumber}
          {filteredData.surfaceMax}
          {filteredData.surfaceEvenness}
        </li>
      ))}
    </div>
  );
}

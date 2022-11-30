import React, { useEffect, useState } from 'react'
import { listDensity } from 'api/species/listDensity'
import ReactEcharts from 'echarts-for-react'
import { useRouter } from 'next/router'

interface Speciesprops {
  SpeciesId: string
}

export const BarSum = ({SpeciesId}:Speciesprops) => {
interface IProps {}
interface IState {}
var ThisSpeciesName =""
var DataXStation = new Array();
var DataXYear = new Array();
var DataXSurface = new Array();
var DataXZone = new Array(); 
console.log("data : " + DataXYear )
console.log("data : " + DataXSurface )
console.log("data : " + DataXZone )
const router = useRouter()
const { locale } = router
const [density, setDensity] = useState([])

const fetchDensity = async (speciesId?: string) => {
  const { data } = await listDensity(speciesId)
  setDensity(data)
}

useEffect(() => {
  fetchDensity(router.query.speciesId as string)
}, [router.query.speciesId])
var Loop = 0
{density.map((density) => (
  <li>
    {Loop=Loop+1}
    {ThisSpeciesName = density?.speciesName}
    {console.log("Loop is : " + Loop )}
    {console.log("new numbers is : " + density?.speciesName )}
    {console.log("new numbers is : " + density?.year )}
    {DataXStation.push(density?.stationName)}
    {DataXYear.push(density?.year)}
    {DataXSurface.push(density?.surface)}
    {DataXZone.push(density?.euphotic_zone)}
    {console.log("new numbers is : " + density?.assetName )}
    {console.log("new numbers is : " + density?.platformName )}
    {console.log("new numbers is : " + density?.stationName )}
    {console.log("new numbers is : " + density?.surface )}
    {console.log("new numbers is : " + density?.euphotic_zone )}
  </li>
))}
console.log("data : " + DataXYear )
console.log("data : " + DataXSurface )
console.log("data : " + DataXZone )
class BarSummary extends React.Component<IProps, IState> {
  getOption = () => ({
    title: {
      text: ThisSpeciesName,
      subtext: DataXYear[0]
    },
    tooltip: {
      trigger: "axis"
    },
    legend: {
      data: ["At Surface", "At Euphotic Zone"]
    },
    grid: {
      left: "3%",
      right: "4%",
      bottom: "3%",
      containLabel: true
    },
    toolbox: {
      feature: {
        saveAsImage: {}
      }
    },
    xAxis: {
      type: "category",
      data: DataXStation
    },
    yAxis: {
      type: "value"
    },
    series: [
      {
        name: "Surface",
        type: "bar",
        data: DataXSurface,
        animationDuration: 5000
      },
      {
        name: "Euphotic Zone",
        type: "bar",
        data: DataXZone,
        animationDuration: 5000
      }
      
    ]
  });

  render() {
    return <ReactEcharts option={this.getOption()} />;
  }
}
  return (
    <div>
      <BarSummary/>
    </div>
  )
}

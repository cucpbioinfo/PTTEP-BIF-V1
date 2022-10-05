import React from 'react'
import ReactEcharts from 'echarts-for-react'

interface Speciesprops {
  SpeciesName: string
}

export const PieBU = ({SpeciesName}:Speciesprops) => {
  
class BasicPieChartBU extends React.Component {
    getOption = () => ({
        title: {
          text: SpeciesName,
          // subtext: "by year",
          x: "center"
        },
        tooltip: {
          trigger: "item",
          formatter: "{a} <br/>{b} : {c} ({d}%)"
        },
        legend: {
          orient: "vertical",
          left: "left",
          data: [ "2019", "2020", "2021", "2022"]
        },
        series: [
          {
            name: SpeciesName,
            type: "pie",
            radius: "55%",
            center: ["50%", "60%"],
            animationDuration: 5000,
            data: [
              { value: 310, name: "2019" },
              { value: 540, name: "2020" },
              { value: 135, name: "2021" },
              { value: 1548, name: "2022" }
            ],
            itemStyle: {
              emphasis: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: "rgba(0, 0, 0, 0.5)"
              }
            }
          }
        ]
      });

  render() {
    return <ReactEcharts option={this.getOption()} />;
  }
}
  return (
    <div>
      <BasicPieChartBU/>
    </div>
  )
}

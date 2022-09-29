import React from 'react'
import ReactEcharts from 'echarts-for-react'

interface Speciesprops {
  SpeciesName: string
}

export const Bar2 = ({SpeciesName}:Speciesprops) => {
interface IProps {}
interface IState {}
var datax = [310, 540, 135, 1548]
class BasicBar2 extends React.Component<IProps, IState> {
  getOption = () => ({
    title: {
      text: SpeciesName,
      subtext: "by EChart"
    },
    tooltip: {
      trigger: "axis"
    },
    legend: {
      orient: "vertical",
      left: "left",
      data: ["2019", "2020", "2021", "2022"]
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
      data: ["2019", "2020", "2021", "2022"]
    },
    yAxis: {
      type: "value"
    },
    series: [
      // {
      //   name: "Specie Name 2",
      //   type: "bar",
      //   data: [120, 132, 101, 134, 90, 230, 210],
      //   animationDuration: 5000
      // },
      {
        name: SpeciesName,
        type: "line",
        data: datax,
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
      <BasicBar2/>
    </div>
  )
}

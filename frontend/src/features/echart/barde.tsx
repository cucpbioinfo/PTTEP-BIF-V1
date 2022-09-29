import React from 'react'
import ReactEcharts from 'echarts-for-react'

interface DEBar {
  barname: string
  stationId: string
}

export const BarDE = ({barname,stationId}:DEBar) => {
interface IProps {}
interface IState {}
var datax = [310, 540, 135, 600]
var datax2 = [350, 450, 235, 550]
var datayear = ["2019", "2020", "2021", "2022"]
class BarDE extends React.Component<IProps, IState> {
  getOption = () => ({
    title: {
      text: barname,
      subtext: "demo station"
    },
    tooltip: {
      trigger: "axis"
    },
    legend: {
      orient: "vertical",
      left: "left",
      data: datayear
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
      data: datayear
    },
    yAxis: {
      type: "value"
    },
    series: [
      {
        name: "Specie Name 2",
        type: "bar",
        data: datax2,
        animationDuration: 5000
      },
      {
        name: stationId,
        type: "bar",
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
      <BarDE/>
    </div>
  )
}

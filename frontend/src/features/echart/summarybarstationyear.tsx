import React, { useEffect, useState } from 'react'
import { listDensity } from 'api/species/listDensity'
import { listAsset } from 'api/dashboard/listAsset'
import { listPlatform } from 'api/dashboard/listPlatform'
import { listStation } from 'api/dashboard/listStation'
import ReactEcharts from 'echarts-for-react'
import { useRouter } from 'next/router'
import { Button , Checkbox , Select ,Table } from 'antd';
import type { CheckboxValueType } from 'antd/es/checkbox/Group';

// export const SummaryBarStationYear = ({summaryId,name,year,surface,zone}:any) => {
  export const SummaryBarStationYear = ({dataimport}:any) => {
  interface IProps {}
  interface IState {}
 
  const StationNameArr = []
  const YearArr = []
  const SurfaceDiversityArr = []
  const ZoneDiversityArr = []
  const SurfaceEvennessArr = []
  const ZoneEvennessArr = []

  
  dataimport.map((data) => 
    (
      StationNameArr.push(data.stationName),
      YearArr.push(data.year),
      SurfaceDiversityArr.push(data.surfaceShannon),
      ZoneDiversityArr.push(data.euphoticzoneShannon),
      SurfaceEvennessArr.push(data.surfaceEvenness),
      ZoneEvennessArr.push(data.euphoticzoneEvenness)
    )
  )
  // const dataSource = [];
  // const columns = [
  //   {
  //     title: 'Name',
  //     dataIndex: 'name',
  //     key: 'name',
  //   },
  //   {
  //     title: 'Year',
  //     dataIndex: 'year',
  //     key: 'year',
  //   },
  //   {
  //     title: 'Surface',
  //     dataIndex: 'surface',
  //     key: 'surface',
  //   },
  //   {
  //     title: 'Euphotic Zone',
  //     dataIndex: 'euphotic_zone',
  //     key: 'euphotic_zone',
  //   },
  // ];
  // dataSource.push({
  //         key: summaryId,
  //         name: name,
  //         year: year,
  //         surface: surface,
  //         euphotic_zone: zone,
  //       })
  const colors = ['#5470C6', '#91CC75', '#EE6666'];
  class SummaryBar extends React.Component<IProps, IState> {
    
    getOption = () => ({
  //     color: colors,

  // tooltip: {
  //   trigger: 'axis',
  //   axisPointer: {
  //     type: 'cross'
  //   }
  // },
  // grid: {
  //   right: '20%'
  // },
  // toolbox: {
  //   feature: {
  //     dataView: { show: true, readOnly: false },
  //     restore: { show: true },
  //     saveAsImage: { show: true }
  //   }
  // },
  // legend: {
  //   data: ['Evaporation', 'Precipitation', 'Temperature']
  // },
  // xAxis: [
  //   {
  //     type: 'category',
  //     axisTick: {
  //       alignWithLabel: true
  //     },
  //     // prettier-ignore
  //     data: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
  //   }
  // ],
  // yAxis: [
  //   {
  //     type: 'value',
  //     name: 'Evaporation',
  //     position: 'right',
  //     alignTicks: true,
  //     axisLine: {
  //       show: true,
  //       lineStyle: {
  //         color: colors[0]
  //       }
  //     },
  //     axisLabel: {
  //       formatter: '{value} ml'
  //     }
  //   },
  //   {
  //     type: 'value',
  //     name: 'Precipitation',
  //     position: 'right',
  //     alignTicks: true,
  //     offset: 80,
  //     axisLine: {
  //       show: true,
  //       lineStyle: {
  //         color: colors[1]
  //       }
  //     },
  //     axisLabel: {
  //       formatter: '{value} ml'
  //     }
  //   },
  //   {
  //     type: 'value',
  //     name: '温度',
  //     position: 'left',
  //     alignTicks: true,
  //     axisLine: {
  //       show: true,
  //       lineStyle: {
  //         color: colors[2]
  //       }
  //     },
  //     axisLabel: {
  //       formatter: '{value} °C'
  //     }
  //   }
  // ],
  // series: [
  //   {
  //     name: 'Evaporation',
  //     type: 'bar',
  //     data: [
  //       2.0, 4.9, 7.0, 23.2, 25.6, 76.7, 135.6, 162.2, 32.6, 20.0, 6.4, 3.3
  //     ]
  //   },
  //   {
  //     name: 'Precipitation',
  //     type: 'bar',
  //     yAxisIndex: 1,
  //     data: [
  //       2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3
  //     ]
  //   },
  //   {
  //     name: 'Temperature',
  //     type: 'line',
  //     yAxisIndex: 2,
  //     data: [2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 6.2]
  //   }
  // ]
      
      title: {
        text: StationNameArr
        //subtext: DataXYear[0]
      },
      tooltip: {
        trigger: "item",
        name:StationNameArr,
        axisPointer: {
          animation: true
        }
      },
      legend: {
        data: ["Surface", "Euphotic Zone"]
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
        data: YearArr
      },
      yAxis: {
        type: "value",
        name: "Y"
      },
      dataZoom: [
        {
          type: "slider"
        },
        {
          type: "inside"
        }
      ],
      series: [
        {
          name: "Surface",
          type: "bar",
          data: SurfaceDiversityArr,
          animationDuration: 250,
          label: {
            normal: {
                show: true,
                position: 'top'
              }
          }
        },
        {
          name: "Euphotic Zone",
          type: "bar",
          data: ZoneDiversityArr,
          animationDuration: 250,
          label: {
            normal: {
                show: true,
                position: 'top'
              }
          }
        }
        
      ]
    });

    render() {
      return <><ReactEcharts option={this.getOption()} />{555}</>;
    }
  }
    return (
      
      <div>
  
        <SummaryBar/>
        {/* <Table dataSource={dataSource} columns={columns} /> */}
        
      </div>

    )
}
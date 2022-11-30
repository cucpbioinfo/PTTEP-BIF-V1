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
  export const SummaryBarStationYear = ({type,dataimport}:any) => {
  interface IProps {}
  interface IState {}
 
  const StationNameArr = []
  const YearArr = []
  const SurfaceDiversityArr = []
  const ZoneDiversityArr = []
  const SurfaceEvennessArr = []
  const ZoneEvennessArr = []
  const SurfaceNumberArr = []
  const ZoneNumberArr = []
  const XBarName = []
  var Dat1 = []
  var Dat2 = []
  var HeaderName = ""
  const dataSourcediversity = []
  const dataSourceevenness = []
  const dataSourcenumber = []
  var dataSource = []
  
  dataimport.map((data) => 
    (
      StationNameArr.push(data.stationName),
      YearArr.push(data.year),
      SurfaceDiversityArr.push(data.surfaceShannon),
      ZoneDiversityArr.push(data.euphoticzoneShannon),
      SurfaceEvennessArr.push(data.surfaceEvenness),
      ZoneEvennessArr.push(data.euphoticzoneEvenness),
      SurfaceNumberArr.push(data.surfaceNumber),
      ZoneNumberArr.push(data.euphoticzoneNumber),
      XBarName.push(data.year),

      
      dataSourcediversity.push({
        key: data.summaryId,
        name: data.stationName.toUpperCase(),
        year: data.year,
        surface: data.surfaceShannon,
        euphotic_zone: data.euphoticzoneShannon
      }),
      dataSourceevenness.push({
        key: data.summaryId,
        name: data.stationName.toUpperCase(),
        year: data.year,
        surface: data.surfaceEvenness,
        euphotic_zone: data.euphoticzoneEvenness
      }),
      dataSourcenumber.push({
        key: data.summaryId,
        name: data.stationName.toUpperCase(),
        year: data.year,
        surface: data.surfaceNumber,
        euphotic_zone: data.euphoticzoneNumber
      })

    )
  )
if(type === "diversity"){
  HeaderName = "Shannon-Weiner Species Diversity Index"
  Dat1 = SurfaceDiversityArr;
  Dat2 = ZoneDiversityArr;
  dataSource = dataSourcediversity;
}
else if(type === "evenness"){
  HeaderName = "Evenness Index"
  Dat1 = SurfaceEvennessArr;
  Dat2 = ZoneEvennessArr;
  dataSource = dataSourceevenness;
}
else if(type === "number"){
  HeaderName = "Number of Species"
  Dat1 = SurfaceNumberArr;
  Dat2 = ZoneNumberArr;
  dataSource = dataSourcenumber;
}
  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'Year',
      dataIndex: 'year',
      key: 'year',
    },
    {
      title: 'At Surface',
      dataIndex: 'surface',
      key: 'surface',
    },
    {
      title: 'At Euphotic Zone',
      dataIndex: 'euphotic_zone',
      key: 'euphotic_zone',
    },
  ];
  
  const colors = ['#5470C6', '#91CC75', '#EE6666'];
  class SummaryBar extends React.Component<IProps, IState> {
    
    getOption = () => ({
        
      title: {
        text: HeaderName
      },
      tooltip: {
        trigger: "item",
        name:StationNameArr,
        //formatter:"test "+{}+" some number",
        axisPointer: {
          animation: true
        }
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
        data: XBarName
      },
      yAxis: {
        type: "value",
        //name: "Y"
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
          name: "At Surface",
          type: "bar",
          data: Dat1,
          animationDuration: 250,
          label: {
            normal: {
                show: true,
                position: 'top'
              }
          }
        },
        {
          name: "At Euphotic Zone",
          type: "bar",
          data: Dat2,
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
      return <><ReactEcharts option={this.getOption()} /></>;
    }
  }
  
    return (
      
      <div>
  
        <SummaryBar/>
        <Table dataSource={dataSource} columns={columns} />
          
      </div>

    )
}
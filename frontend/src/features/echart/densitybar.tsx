import React, { useEffect, useState } from 'react'
import ReactEcharts from 'echarts-for-react'
import { Button , Checkbox , Select ,Table } from 'antd';

// export const SummaryBarStationYear = ({summaryId,name,year,surface,zone}:any) => {
  export const DensityBar = ({dataimport}:any) => {
  interface IProps {}
  interface IState {}
 
  const StationNameArr = []
  const YearArr = []
  const SurfaceDensityArr = []
  const ZoneDensityArr = []

  const XBarName = []
  var Dat1 = []
  var Dat2 = []
  var HeaderName = ""
  const dataSourcedensity = []

  var dataSource = []
  
  function capitalizeFirstLetter(string) {
    return string.charAt(0).toUpperCase() + string.slice(1);
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

  dataimport.map((data) => 
    (
      StationNameArr.push(data.stationName),
      YearArr.push(data.year),
      SurfaceDensityArr.push(data.surface),
      ZoneDensityArr.push(data.euphotic_zone),
      XBarName.push(data.speciesName+" "+data.year),

      
      dataSourcedensity.push({
        key: data.summaryId,
        name: capitalizeFirstLetter(data.speciesName),
        station: data.stationName.toUpperCase(),
        year: data.year,
        surface: data.surface,
        euphotic_zone: data.euphotic_zone
      })

    )
  )

  HeaderName = "Density"
  Dat1 = SurfaceDensityArr;
  Dat2 = ZoneDensityArr;
  dataSource = dataSourcedensity;

  const columns = [
    {
      title: 'Name',
      dataIndex: 'name',
      key: 'name',
    },
    {
      title: 'station',
      dataIndex: 'station',
      key: 'station',
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
          name: "Surface",
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
          name: "Euphotic Zone",
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
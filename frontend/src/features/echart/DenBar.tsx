import React from 'react'
import ReactEcharts from 'echarts-for-react'

function SpeciesSplice(string:string) {
  if(string.search(/sp./i)==-1){
    let touppercase = string.charAt(0).toUpperCase() + string.slice(1);
    return <>{touppercase}</>;
  }
  else
  {
    let position = string.search(/sp./i);
    let speciesName = string.substring(0, position);
    let upperspeciesName = speciesName.charAt(0).toUpperCase() + speciesName.slice(1);
    let SpecialText = string.slice(position);
    return {upperspeciesName}+" "+{SpecialText};
  }
  }

export const DenBar = ({densityId, speciesId, speciesName,name,year,surface,zone}:any) => {
  interface IProps {}
  interface IState {}
  
  const dataSource = [];
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
      title: 'Surface',
      dataIndex: 'surface',
      key: 'surface',
    },
    {
      title: 'Euphotic Zone',
      dataIndex: 'euphotic_zone',
      key: 'euphotic_zone',
    },
  ];
  dataSource.push({
          key: densityId,
          name: name,
          year: year,
          surface: surface,
          euphotic_zone: zone,
        })
  
  class BarDensity extends React.Component<IProps, IState> {
    getOption = () => ({
      title: {
        //text: speciesName
        //subtext: DataXYear[0]
        text: speciesName.charAt(0).toUpperCase() + speciesName.slice(1),
        // "<div className='flex'><div className='italic'>A</div> <div className='not-italic'>B</div></div>"
        //subtext: "Sub Title",
        // left: "center",
        // top: "center",
        textStyle: {
          // fontSize: 30,
          fontStyle: "italic"
        },
        // subtextStyle: {
        //   fontSize: 20
        // }
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
      // dataZoom: [
      //   {
      //     type: "slider"
      //   },
      //   {
      //     type: "inside"
      //   }
      // ],
      xAxis: {
        type: "category",
        data: [year]
      },
      yAxis: {
        type: "value"
      },
      series: [
        {
          name: "At Surface",
          type: "bar",
          data: [surface],
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
          data: [zone],
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
      return <ReactEcharts option={this.getOption()} />;
    }
  }
    return (
      <div>
        <BarDensity/>
      </div>
    )
}
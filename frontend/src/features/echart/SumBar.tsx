import React, { useEffect, useState } from 'react'
import { listDensity } from 'api/species/listDensity'
import { listAsset } from 'api/dashboard/listAsset'
import { listPlatform } from 'api/dashboard/listPlatform'
import { listStation } from 'api/dashboard/listStation'
import ReactEcharts from 'echarts-for-react'
//import { useRouter } from 'next/router'
import { Button , Checkbox , Select ,Table } from 'antd';
import type { CheckboxValueType } from 'antd/es/checkbox/Group';

// interface Speciesprops {
//   speciesName: string
//   year:string
//   surface:string
//   zone:string
// }

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

export const SumBar = ({densityId, speciesId, speciesName,name,year,surface,zone}:any) => {
  interface IProps {}
  interface IState {}
  // console.log(surface)
  // console.log(zone)
  //const router = useRouter()
  // var DataXStation = new Array();
  // var DataXYear = new Array();
  // var DataXSurface = new Array();
  // var DataXZone = new Array(); 

  //for checkbox,select
  // var AllYearData = new Array();
  // var AllAssetData = new Array();
  // var AllPlatformData = new Array();
  // var AllStationData = new Array();

  // var CheckedStation = new Array();

  // // console.log("data : " + DataXYear )
  // // console.log("data : " + DataXSurface )
  // // console.log("data : " + DataXZone )

  // const router = useRouter()
  // const { locale } = router
  // const { Option } = Select;

  // const [asset, setAsset] = useState([])
  // const [platform, setPlatform] = useState([])
  // const [station, setStation] = useState([])

  // const [density, setDensity] = useState([])
  // const [densityfilter, setDensityfilter] = useState([])

  // // const [year, setYear] = useState([])
  // // const [asset, setAsset] = useState([])
  // // const [platform, setPlatform] = useState([])
  // // const [statio, setStation] = useState([])
  // //fetch data by species by API
  // const fetchDensity = async (speciesId?: string) => {
  //   const { data } = await listDensity(speciesId)
  //   setDensity(data)
  // }
  
  // const fetchAsset = async () => {
  //   const { data } = await listAsset()
  //   setAsset(data)
  // }
  // const fetchPlatform = async (assetId?: string) => {
  //   const { data } = await listPlatform(assetId)
  //   setPlatform(data)
  // }
  // const fetchStation = async (platformId?: string) => {
  //   const { data } = await listStation(platformId)
  //   setStation(data)
  // }

  // useEffect(() => {
  //   fetchDensity(router.query.speciesId as string)
  // }, [router.query.speciesId])

  // useEffect(() => {
  //   fetchAsset()
  // }, [])
  // useEffect(() => {
  //   fetchPlatform(router.query.assetId as string)
  // }, [router.query.assetId])
  // useEffect(() => {
  //   fetchStation(router.query.platformId as string)
  // }, [router.query.platformId])
  


  // const onChange = (checkedValues: CheckboxValueType[]) => {
  //   CheckedStation = checkedValues;
  //   console.log('checked = ', checkedValues);
  // };
  // const onChangeStation = (checkedValues: CheckboxValueType[]) => {
  //   CheckedStation = checkedValues;
  //   console.log('onchangestation = ', checkedValues);
  //   density.filter(function(density) {
  //     return density.stationName == "pps3";
      
  //   });
    
  // };

  // const plainOptions = ['Apple', 'Pear', 'Orange'];
  // const options = [
  //   { label: 'Apple', value: 'Apple' },
  //   { label: 'Pear', value: 'Pear' },
  //   { label: 'Orange', value: 'Orange' },
  // ];
  // const optionsWithDisabled = [
  //   { label: 'Apple', value: 'Apple' },
  //   { label: 'Pear', value: 'Pear' },
  //   { label: 'Orange', value: 'Orange', disabled: false },
  // ];



  //to push data to array for filter option without api
  // const CheckDataDupe = (TargetArray,Data) => {
  //   if(TargetArray.includes(Data))
  //   {    console.log("already have data");  }
  //   else
  //   {    TargetArray.push(Data);
  //       TargetArray.sort();  }
  // };

  // const onFilterChange = (key: string, value?: string) => {
  //   const url = {
  //     pathname: router.pathname,
  //     query: {
  //       ...router.query,
  //       [key]: value,
  //     },
  //   }
  //   router.push(url, url, { locale })
  // }

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
  
  // density.map((density) => 
  //   (
      
  //     ThisSpeciesName = density?.speciesName,
  //     DataXStation.push(density?.stationName),
  //     DataXYear.push(density?.year),
  //     DataXSurface.push(density?.surface),
  //     DataXZone.push(density?.euphotic_zone),

  //     CheckDataDupe(AllYearData,density?.year),
  //     CheckDataDupe(AllAssetData,density?.assetName),
  //     CheckDataDupe(AllPlatformData,density?.platformName),
  //     CheckDataDupe(AllStationData,density?.stationName),
      
  //     dataSource.push({
  //       key: density?.densityId,
  //       name: density?.stationName,
  //       year: density?.year,
  //       surface: density?.surface,
  //       euphotic_zone: density?.euphotic_zone,
  //     })
  //   )
  // )

  //var Loop = 0
  // density.map((density) => 
  //   (
  //     //Loop=Loop+1,
  //     ThisSpeciesName = density?.speciesName,
  //     // console.log("Loop is : " + Loop) ,
  //     // console.log("new numbers is : " + density?.speciesName ),
  //     DataXStation.push(density?.stationName),
  //     DataXYear.push(density?.year),
  //     DataXSurface.push(density?.surface),
  //     DataXZone.push(density?.euphotic_zone),
  //     //console.log("new numbers is : " + density?.year ),
  //     CheckDataDupe(AllYearData,density?.year),
  //     //AllYearData.push(density?.year),
  //     //console.log("new numbers is : " + density?.assetName ),
  //     CheckDataDupe(AllAssetData,density?.assetName),
  //     //AllAssetData.push(density?.assetName),     
  //     //console.log("new numbers is : " + density?.platformName ),
  //     CheckDataDupe(AllPlatformData,density?.platformName),
  //     //AllPlatformData.push(density?.platformName),   
  //     //console.log("new numbers is : " + density?.stationName ),
  //     CheckDataDupe(AllStationData,density?.stationName)
  //     //AllStationData.push(density?.stationName),    
  //     //console.log("new numbers is : " + density?.surface ),
  //     //console.log("new numbers is : " + density?.euphotic_zone )
  // )
  // )
  // console.log(AllYearData)
  // console.log(AllAssetData)
  // console.log(AllPlatformData)
  // console.log(AllStationData)
  // console.log("data : " + DataXYear )
  // console.log("data : " + DataXSurface )
  // console.log("data : " + DataXZone )
  //{SpeciesSpliceFirst(speciesName)} {SpeciesSpliceSpecial(speciesName)
  class BarDensity extends React.Component<IProps, IState> {
    getOption = () => ({
      title: {
        //text: speciesName
        //subtext: DataXYear[0]
        //text: speciesName,
        //subtext: "Sub Title",
        // left: "center",
        // top: "center",
        // textStyle: {
        //   // fontSize: 30,
        //   fontStyle: "italic"
        // },
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
          //saveAsImage: {}
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
        


        {/* <Select
          placeholder="Year"
          onChange={(value: string | undefined) => 
            {onFilterChange('year', value)}}>
          {AllYearData.map((year) => (
            <Option
              value={year}
              key={year}
            >
              {year}
            </Option>
          ))
          }
        </Select> */}

        {/* <div>Asset</div>
        <Checkbox.Group 
          options={AllAssetData} 
          onChange={onChange} /> */}


        {/* <Select
          placeholder="Asset"
          onChange={(value: string | undefined) => 
            {onFilterChange('assetName', value)}}>
          {AllAssetData.map((Asset) => (
            <Option
              value={Asset}
              key={Asset}
            >
              {Asset}
            </Option>
          ))
          }
        </Select> */}

        {/* <Select
          placeholder="Asset"
          onChange={(value: string | undefined) => 
            {onFilterChange('assetName', value)}}>
          {AllAssetData.map((Asset) => (
            <Option
              value={Asset}
              key={Asset}
            >
              {Asset}
            </Option>
          ))
          }
        </Select> */}

        {/* <div>Asset</div>
        <Select
          placeholder="Asset"
          onChange={(value: string | undefined) => 
            {onFilterChange('assetId', value)}}>
          {asset.map((asset) => (
            <Option
              value={asset?.assetId}
              key={asset?.assetId}
            >
              {asset?.assetName}
            </Option>
          ))
          }
        </Select>
        <div>Platform</div>
        <Select
          placeholder="Platform"
          onChange={(value: string | undefined) => 
            {onFilterChange('platformId', value)}}>
          {platform.map((platform) => (
            <Option
              value={platform?.platformId}
              key={platform?.platformId}
            >
              {platform?.platformName}
            </Option>
          ))
          }
        </Select>
        <div>Station</div>
        <Select
          placeholder="Station"
          onChange={(value: string | undefined) => 
            {onFilterChange('stationId', value)}}>
          {station.map((station) => (
            <Option
              value={station?.stationId}
              key={station?.stationId}
            >
              {station?.stationName}
            </Option>
          ))
          }
        </Select>


        <div>Year</div>
        <Checkbox.Group 
          options={AllYearData} 
          onChange={onChange} />
        <Button>Compare</Button> */}

        {/* <div>Platform</div>
        <Checkbox.Group 
          options={AllPlatformData} 
          onChange={onChange} /> */}

        {/* <Select
          placeholder="Platform"
          onChange={(value: string | undefined) => 
            {onFilterChange('platformName', value)}}>
          {AllPlatformData.map((Platform) => (
            <Option
              value={Platform}
              key={Platform}
            >
              {Platform}
            </Option>
          ))
          }
        </Select> */}

        {/* <div>Station</div>
        <Checkbox.Group 
          options={AllStationData} 
          onChange={onChangeStation} /> */}

        {/* <Select
          placeholder="Station"
          onChange={(value: string | undefined) => 
            {onFilterChange('stationName', value)}}>
          {AllStationData.map((Station) => (
            <Option
              value={Station}
              key={Station}
            >
              {Station}
            </Option>
          ))
          }
        </Select> */}

        {/* <Select defaultValue="lucy" style={{ width: 120 }} >
          <Option value="jack">Jack</Option>
          <Option value="lucy">Lucy</Option>
          <Option value="disabled" disabled>
            Disabled
          </Option>
        <Option value="Yiminghe">yiminghe</Option>
        </Select>
      <>
        <Checkbox.Group 
          options={plainOptions} 
          defaultValue={['Apple']} 
          onChange={onChange} />
        <br />
        <br />
        <Checkbox.Group 
          options={options} 
          defaultValue={['Pear']} 
          onChange={onChange} />
        <br />
        <br />
        <Checkbox.Group
          options={optionsWithDisabled}
          disabled
          defaultValue={['Apple']}
          onChange={onChange} /> 
      </> */}
      {/* <a onClick={() => router.push(`/filterdata?speciesId=${speciesId}`)}>See all</a> */}
        <BarDensity/>
        
        {/* <Table dataSource={dataSource} columns={columns} /> */}
        {/* <Table dataSource={dataSource} columns={columns} />; */}
      </div>
    )
}


// let array = [
//   { id: 3 },
//   { id: -1 },
//   { id: 0 },
//   { id: 15 },
//   { id: 12.2 },
//   { },
//   { id: null },
//   { id: NaN },
//   { id: 'undefined' }
// ]
  
// let countInvalidEntries = 0
  
// function filterById(obj) {
//   if (Number.isFinite(obj.id) && obj.id !== 0) 
//   {
//     return true
//   } 
//   countInvalidEntries++
//   return false;
// }
  
// let arrayById = array.filter(filterById);
  
// console.log(arrayById);
  
// console.log(countInvalidEntries);
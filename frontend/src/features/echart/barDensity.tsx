import React, { useEffect, useState } from 'react'
import { listDensity } from 'api/species/listDensity'
import ReactEcharts from 'echarts-for-react'
import { useRouter } from 'next/router'
import { Checkbox , Select ,Table } from 'antd';
import type { CheckboxValueType } from 'antd/es/checkbox/Group';

interface Speciesprops {
  SpeciesId: string
}

export const DensityBar = ({SpeciesId}:Speciesprops) => {
interface IProps {}
interface IState {}
var ThisSpeciesName ="";
var DataXStation = new Array();
var DataXYear = new Array();
var DataXSurface = new Array();
var DataXZone = new Array(); 

var AllYearData = new Array();
var AllAssetData = new Array();
var AllPlatformData = new Array();
var AllStationData = new Array();

var CheckedStation = new Array();

// console.log("data : " + DataXYear )
// console.log("data : " + DataXSurface )
// console.log("data : " + DataXZone )

const router = useRouter()
const { locale } = router
const { Option } = Select;
const [density, setDensity] = useState([])

// const [year, setYear] = useState([])
// const [asset, setAsset] = useState([])
// const [platform, setPlatform] = useState([])
// const [statio, setStation] = useState([])


const onChange = (checkedValues: CheckboxValueType[]) => {
  CheckedStation = checkedValues;
  console.log('checked = ', checkedValues);
};

const plainOptions = ['Apple', 'Pear', 'Orange'];
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

//fetch data by species by API
const fetchDensity = async (speciesId?: string) => {
  const { data } = await listDensity(speciesId)
  setDensity(data)
}

//to push data to array for filter option without api
const CheckDataDupe = (TargetArray,Data) => {
  if(TargetArray.includes(Data))
  {    console.log("already have data");  }
  else
  {    TargetArray.push(Data);
       TargetArray.sort();  }
};

const onFilterChange = (key: string, value?: string) => {
  const url = {
    pathname: router.pathname,
    query: {
      ...router.query,
      [key]: value,
    },
  }
  router.push(url, url, { locale })
}

useEffect(() => {
  fetchDensity(router.query.speciesId as string)
}, [router.query.speciesId])
//var Loop = 0
density.map((density) => 
  (
    //Loop=Loop+1,
    ThisSpeciesName = density?.speciesName,
    // console.log("Loop is : " + Loop) ,
    // console.log("new numbers is : " + density?.speciesName ),
    DataXStation.push(density?.stationName),
    DataXYear.push(density?.year),
    DataXSurface.push(density?.surface),
    DataXZone.push(density?.euphotic_zone),

    //console.log("new numbers is : " + density?.year ),
    CheckDataDupe(AllYearData,density?.year),
    //AllYearData.push(density?.year),
    
    //console.log("new numbers is : " + density?.assetName ),
    CheckDataDupe(AllAssetData,density?.assetName),
    //AllAssetData.push(density?.assetName),
    
    //console.log("new numbers is : " + density?.platformName ),
    CheckDataDupe(AllPlatformData,density?.platformName),
    //AllPlatformData.push(density?.platformName),
    
    //console.log("new numbers is : " + density?.stationName ),
    CheckDataDupe(AllStationData,density?.stationName)
    //AllStationData.push(density?.stationName),
    
    //console.log("new numbers is : " + density?.surface ),
    //console.log("new numbers is : " + density?.euphotic_zone )
)
)


console.log(AllYearData)
console.log(AllAssetData)
console.log(AllPlatformData)
console.log(AllStationData)


const dataSource = [
  {
    key: '1',
    name: 'Mike',
    age: 32,
    address: '10 Downing Street',
  },
  {
    key: '2',
    name: 'John',
    age: 42,
    address: '10 Downing Street',
  },
];

const columns = [
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age',
  },
  {
    title: 'Address',
    dataIndex: 'address',
    key: 'address',
  },
];


// console.log("data : " + DataXYear )
// console.log("data : " + DataXSurface )
// console.log("data : " + DataXZone )
class BarDensity extends React.Component<IProps, IState> {
  getOption = () => ({
    title: {
      text: ThisSpeciesName,
      subtext: DataXYear[0]
    },
    tooltip: {
      trigger: "axis"
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
      
      <Select
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
      </Select>

      <Select
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
      </Select>

      <Select
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
      </Select>

      <Select
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
      </Select>

      <Checkbox.Group 
        options={AllStationData} 
        onChange={onChange} />





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

      <BarDensity/>

      <Table dataSource={dataSource} columns={columns} />;
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
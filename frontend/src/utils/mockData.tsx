export const SPECIES_DETAIL = {
  speciesId: '1',
  scientificName: 'full scientificName',
  commonName: 'commonName',
  thaiName: 'thaiName',
  kingdom: 'kingdom',
  phylum: 'phylum',
  order: 'order',
  speciesClass: 'class',
  family: 'family',
  genus: 'genus',
}

export const MOCK_ZOOPLANKTONS = [
  {
    key: '1',
    speciesId: '1',
    scientificName: 'name1',
    commonName: 'commonName1',
    thaiName: 'Thainame1',
    higherTaxon: 'higherTaxon',
    location: 'loc',
    count: 'count',
  },
  {
    key: '2',
    speciesId: '2',
    scientificName: 'name2',
    commonName: 'commonName2',
    thaiName: 'Thainame2',
    higherTaxon: 'higherTaxon',
    location: 'loc',
    count: 'count',
  },
  {
    key: '3',
    speciesId: '3',
    scientificName: 'name3',
    commonName: 'commonName3',
    thaiName: 'Thainame3',
    higherTaxon: 'higherTaxon',
    location: 'loc',
    count: 'count',
  },
  {
    key: '4',
    speciesId: '4',
    scientificName: 'name4',
    commonName: 'commonName4',
    thaiName: 'Thainame4',
    higherTaxon: 'higherTaxon',
    location: 'loc',
    count: 'count',
  },
  {
    key: '5',
    speciesId: '5',
    scientificName: 'name5',
    commonName: 'commonName5',
    thaiName: 'Thainame5',
    higherTaxon: 'higherTaxon',
    location: 'loc',
    count: 'count',
  },
]

export const MOCK_LATEST_OCCURRENCES = [
  {
    occurrenceId: '1',
    imageSrc: '/blue-whale.jpeg',
    occurrenceDetails:
      'The blue whale (Balaenoptera musculus) is a marine mammal. Reaching a maximum confirmed length of 29.9 metres and weighing up to 199 tonnes, it is the largest animal known to have ever existed.',
    discoveredAt: new Date(),
  },
  {
    occurrenceId: '2',
    imageSrc: '/chaetoceros.jpeg',
    occurrenceDetails:
      'Chaetoceros is probably the largest genus of marine planktonic diatoms with approximately 400 species described, although many of these descriptions are no longer valid. It is often very difficult to distinguish between different Chaetoceros species.',
    discoveredAt: new Date(),
  },
]

export const MOCK_SPECIES = [
  {
    speciesId: '1',
    speciesName: {
      th: 'แอนเนลิด',
      en: 'annelid',
    },
    kingdomName: {
      th: 'แอนิมอลเลีย',
      en: 'Animalia',
    },
    phylumName: {
      th: 'แอนเนลิดา',
      en: 'annelida',
    },
  },
  {
    speciesId: '2',
    speciesName: {
      th: 'แอนเนลิด',
      en: 'annelid',
    },
    kingdomName: {
      th: 'แอนิมอลเลีย',
      en: 'Animalia',
    },
    phylumName: {
      th: 'แอนเนลิดา',
      en: 'annelida',
    },
  },
  {
    speciesId: '3',
    speciesName: {
      th: 'แอนเนลิด',
      en: 'annelid',
    },
    kingdomName: {
      th: 'แอนิมอลเลีย',
      en: 'Animalia',
    },
    phylumName: {
      th: 'แอนเนลิดา',
      en: 'annelida',
    },
  },
  {
    speciesId: '4',
    speciesName: {
      th: 'แอนเนลิด',
      en: 'annelid',
    },
    kingdomName: {
      th: 'แอนิมอลเลีย',
      en: 'Animalia',
    },
    phylumName: {
      th: 'แอนเนลิดา',
      en: 'annelida',
    },
  },
  {
    speciesId: '5',
    speciesName: {
      th: 'แอนเนลิด',
      en: 'annelid',
    },
    kingdomName: {
      th: 'แอนิมอลเลีย',
      en: 'Animalia',
    },
    phylumName: {
      th: 'แอนเนลิดา',
      en: 'annelida',
    },
  },
]

export const MOCK_SPECIES_FIELD_LOCATIONS = [
  {
    station: 'Division Cyanophyta',
    'app-1b-1': 5200,
    'app-1b-2': 21600,
    'app-3b-1': 8400,
    'app-3b-2': 14880,
    'awp8-1b-1': 2920,
    'awp8-1b-2': 9310,
    'awp8-3b-1': 11360,
    'awp8-3b-2': 33000,
    'awp-1n-1b-1': 34800,
    'awp-1n-1b-2': 40000,
    'awp-1n-3b-1': 0,
    'awp-1n-3b-2': 0,
    'awp29-1b-1': 0,
    'awp29-1b-2': 3810,
    'awp29-3b-1': 11840,
    'awp29-3b-2': 57600,
  },
  {
    station: 'Division Bacillariophyta',
    'app-1b-1': 29900,
    'app-1b-2': 60000,
    'app-3b-1': 93600,
    'app-3b-2': 91760,
    'awp8-1b-1': 125560,
    'awp8-1b-2': 171570,
    'awp8-3b-1': 147680,
    'awp8-3b-2': 157080,
    'awp-1n-1b-1': 233450,
    'awp-1n-1b-2': 321250,
    'awp-1n-3b-1': 28220,
    'awp-1n-3b-2': 41400,
    'awp29-1b-1': 75600,
    'awp29-1b-2': 97790,
    'awp29-3b-1': 124320,
    'awp29-3b-2': 127200,
  },
  {
    station: 'Division Chrysophyta',
    'app-1b-1': '-',
    'app-1b-2': 2400,
    'app-3b-1': '-',
    'app-3b-2': '-',
    'awp8-1b-1': '-',
    'awp8-1b-2': '-',
    'awp8-3b-1': '-',
    'awp8-3b-2': '-',
    'awp-1n-1b-1': '-',
    'awp-1n-1b-2': '-',
    'awp-1n-3b-1': '-',
    'awp-1n-3b-2': '-',
    'awp29-1b-1': '-',
    'awp29-1b-2': '-',
    'awp29-3b-1': '-',
    'awp29-3b-2': '-',
  },
]

export const MOCK_INTERNAL_DATA = [
  {
    title:
      'Environmental Monitoring Database (2022). Arthit platform. APP station. Phytoplankton',
    href: '/location',
    updatedAt: new Date(),
  },
  {
    title:
      'Environmental Monitoring Database (2022). Arthit platform. APP station. Bacteria',
    href: '/location',
    updatedAt: new Date(),
  },
  {
    title:
      'Environmental Monitoring Database (2022). Arthit platform. APP station. Zooplankton',
    href: '/location',
    updatedAt: new Date(),
  },
]

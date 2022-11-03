import { eDNAAxios } from 'api/const'

// export const listAsset = async () => {
//   const response = await (await eDNAAxios.get('/asset')).data
//   return response
// }

export const listYearStation = async (stationId?: string) => {
  const response = await (
    await eDNAAxios.get('/yeardensity', {
      params: {stationId},
    })
  ).data
  return response
}

// export const listEvenness = async () => {
//   const response = await (await eDNAAxios.get('/evenness')).data
//   return response
// }
// import { eDNAAxios } from 'api/const'
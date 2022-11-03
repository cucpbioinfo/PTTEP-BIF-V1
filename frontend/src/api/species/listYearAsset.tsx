import { eDNAAxios } from 'api/const'

// export const listAsset = async () => {
//   const response = await (await eDNAAxios.get('/asset')).data
//   return response
// }

export const listYearAsset = async (assetId?: string) => {
  const response = await (
    await eDNAAxios.get('/yearassetdensity', {
      params: {assetId},
    })
  ).data
  return response
}

// export const listEvenness = async () => {
//   const response = await (await eDNAAxios.get('/evenness')).data
//   return response
// }
// import { eDNAAxios } from 'api/const'
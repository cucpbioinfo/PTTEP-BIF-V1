import { eDNAAxios } from 'api/const'
import { ListAssetDensityQuery } from 'types/density'

export const listAssetDensity = async () => {
  const response = await (await eDNAAxios.get('/assetdensity')).data
  return response
}

// export const listAssetDensity = async (params:ListAssetDensityQuery) => {
//   const response = await (
//     await eDNAAxios.get('/assetdensity', {
//       params: { ...params },
//     })
//   ).data
//   return response
// }

// export const listAssetDensity = async (speciesId?: string) => {
//   const response = await (
//     await eDNAAxios.get('/assetdensity', {
//       params: { speciesId },
//     })
//   ).data
//   return response
// }

// export const listEvenness = async () => {
//   const response = await (await eDNAAxios.get('/evenness')).data
//   return response
// }
// import { eDNAAxios } from 'api/const'

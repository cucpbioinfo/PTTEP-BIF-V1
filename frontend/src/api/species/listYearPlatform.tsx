import { eDNAAxios } from 'api/const'

// export const listAsset = async () => {
//   const response = await (await eDNAAxios.get('/asset')).data
//   return response
// }

export const listYearPlatform = async (platformId?: string) => {
  const response = await (
    await eDNAAxios.get('/yearplatformdensity', {
      params: {platformId},
    })
  ).data
  return response
}

// export const listEvenness = async () => {
//   const response = await (await eDNAAxios.get('/evenness')).data
//   return response
// }
// import { eDNAAxios } from 'api/const'
import { eDNAAxios } from 'api/const'

export const listEvenness = async (stationId?: string) => {
  const response = await (
    await eDNAAxios.get('/evenness', {
      params: { stationId },
    })
  ).data
  return response
}

// export const listEvenness = async () => {
//   const response = await (await eDNAAxios.get('/evenness')).data
//   return response
// }
// import { eDNAAxios } from 'api/const'
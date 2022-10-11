import { eDNAAxios } from 'api/const'

export const listDensity = async (speciesId?: string) => {
  const response = await (
    await eDNAAxios.get('/density', {
      params: { speciesId },
    })
  ).data
  return response
}

// export const listEvenness = async () => {
//   const response = await (await eDNAAxios.get('/evenness')).data
//   return response
// }
// import { eDNAAxios } from 'api/const'
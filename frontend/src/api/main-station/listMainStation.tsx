import { eDNAAxios } from 'api/const'
import { MainStation } from 'types/main-station'
import { Site } from 'types/site'

export const listMainStation = async (
  siteId?: string,
): Promise<MainStation[]> => {
  const res = await (
    await eDNAAxios.get('/main-stations', {
      params: { siteId },
    })
  ).data.data
  return res
}

/** @jsxRuntime classic */
/** @jsx jsx */
import React, { useEffect, useState } from 'react'
import ReactECharts from 'echarts-for-react'
import { getPieChartOption } from 'utils/pieChart'
import { css, jsx } from '@emotion/react'
import { Carousel, Collapse, Divider, Input, Image } from 'antd'
import { SearchOutlined } from '@ant-design/icons'
import { PrimaryButton } from 'components/PrimaryButton'
import { SpeciesFilter } from 'components/SpeciesFilter'
import { PageLayout } from 'components/PageLayout'
import { AdvanceSearchDrawer } from './AdvanceSearchDrawer'
import { useRouter } from 'next/router'
import { SpeciesCount } from 'types/species'
import { DefaultSpeciesCounts } from './const'
import { getSpeciesCounts } from 'api/species/getSpeciesCounts'
import { searchSpecies } from 'api/species/searchSpecies'
import { getSpeciesImagePath } from 'utils/image'

const { Panel } = Collapse

export const MainPage = () => {
  const router = useRouter()

  const [searchText, setSearchText] = useState<string>()
  const [currentSearchValue, setCurrentSearchValue] = useState<string>()
  const [showAdvanceSearchDrawer, setShowAdvanceSearchDrawer] =
    useState<boolean>(false)
  const [showPreviewImage, setShowPreviewImage] = useState<boolean>(false)
  const [species, setSpecies] = useState([])

  const fetchSpecies = async (keyword: string) => {
    const data = await searchSpecies({ keyword, limit: 5 })
    setSpecies(data)
    console.log(data)
  }

  useEffect(() => {
    fetchSpecies(currentSearchValue)
  }, [currentSearchValue])

  const onSearch = () => {
    setCurrentSearchValue(searchText)
  }
  const onChartClick = (params: any) => {
    const { name } = params.data
    router.push(`/species-type/${name}`)
  }

  const parseSpeciesCounts = (data: SpeciesCount[]) => {
    return data.map((speciesCount: SpeciesCount) => ({
      name: speciesCount.speciesType,
      value: speciesCount.count,
    }))
  }
  const [speciesCounts, setSpeciesCounts] = useState(
    parseSpeciesCounts(DefaultSpeciesCounts),
  )
  const fetchSpeciesCounts = async () => {
    const data = await getSpeciesCounts()
    setSpeciesCounts(parseSpeciesCounts(data))
  }

  const onImageClicked = (speciesId: string) => {
    router.push({
      pathname: `/species/${speciesId}`,
      query: router.query,
    })
  }

  useEffect(() => {
    fetchSpeciesCounts()
  }, [])

  return (
    <PageLayout>
      <div>
        <Input
          placeholder="search"
          className="mr-4 rounded-xl w-full"
          suffix={<SearchOutlined onClick={onSearch} />}
          onPressEnter={onSearch}
          css={css`
            .ant-input {
              padding: 0.5rem 0.5rem !important;
              font-size: 1rem;
            }
            max-width: 400px;
          `}
          onChange={(e) => {
            setSearchText(e.target.value)
          }}
        />
        <div>
          <div
            className="mt-2 underline text-secondary cursor-pointer inline-block"
            onClick={() => setShowAdvanceSearchDrawer(true)}
          >
            advance search
          </div>
        </div>
        {currentSearchValue !== undefined && (
          <>
            <Divider />
            <div className="flex mt-8 text-lg">
              Search Result
              <div className="ml-1 font-semibold truncate w-48">
                {currentSearchValue}
              </div>
            </div>
            <div className="mt-2 grid md:grid-cols-4 md:gap-4 grid-cols-1 justify-items-center w-full">
              {species
                .filter((s) => s.speciesImages.length > 0)
                .map((s) => (
                  <div
                    style={{ width: '200px', height: '200px' }}
                    key={s.speciesImages[0]}
                    onClick={() => onImageClicked(s.speciesId)}
                  >
                    <img
                      src={getSpeciesImagePath(s.speciesImages[0])}
                      className="w-full h-full object-contain"
                    />
                  </div>
                ))}
            </div>
          </>
        )}
      </div>
      <AdvanceSearchDrawer
        visible={showAdvanceSearchDrawer}
        onClose={() => setShowAdvanceSearchDrawer(false)}
      />
      <Divider />
      <div className="mt-16">
        <ReactECharts
          option={getPieChartOption(speciesCounts)}
          style={{ height: '400px' }}
          onEvents={{ click: onChartClick }}
        />
      </div>
    </PageLayout>
  )
}

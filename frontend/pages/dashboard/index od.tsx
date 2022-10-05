/** @jsxRuntime classic */
/** @jsx jsx */
import { css, jsx } from '@emotion/react'
import { Divider } from 'antd'
import { LatestOccurenceCard } from 'components/new/LatestOccurenceCard'
import { Navbar } from 'components/new/Navbar'
import { LatestOccurrence } from 'features/new/LatestOccurrence'
import React from 'react'

const LandingPage = () => {
  return (
    <div>
      <Navbar />
      <div
        css={css`
          width: 100vw;
          background: url(landing-bg.png);
          background-repeat: no-repeat;
          height: 500px;
          background-size: cover;
          opacity: 0.5;
        `}
      >
        <div className="md:pt-8 pt-32 md:px-32 px-16 h-full">
          <h1 className="md:text-4xl text-2xl text-white font-bold">
            Biodiversity at Offshore Oil & Gas
            <br />
            Platform in Gulf of Thailand
          </h1>
        </div>
      </div>
      <div className="py-4 md:px-32 px-8 bg-primary md:flex block justify-center items-center">
        <h2 className="text-white md:text-2xl text-xl text-center">
          Looking for data
        </h2>
        <div className="w-1 bg-white h-10 ml-4 md:block hidden"></div>
        <div className="flex justify-center md:mt-0 mt-4">
          <div className="ml-8 text-white md:text-2xl text-xl cursor-pointer hover:underline">
            Access Data
          </div>
          <div className="ml-4 text-white md:text-2xl text-xl cursor-pointer hover:underline">
            Archieve Data
          </div>
        </div>
      </div>
      <div className="md:py-16 py-8 md:px-32 px-4">
        {/* <LatestOccurrence /> */}
      </div>
    </div>
  )
}

export default LandingPage

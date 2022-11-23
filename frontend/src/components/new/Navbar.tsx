/** @jsxRuntime classic */
/** @jsx jsx */
import { MenuOutlined } from '@ant-design/icons'
import { css, jsx } from '@emotion/react'
import { Dropdown, Input, Menu } from 'antd'
import { Content } from 'antd/lib/layout/layout'
import classNames from 'classnames'
import { OverlayMenu } from 'components/OverlayMenu'
import { useAuth } from 'contexts/AuthContext'
import Link from 'next/link'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { LOGIN_TEXT, NAVBAR_MENU_LIST } from './locale'

const { Search } = Input

export const Navbar = () => {
  const router = useRouter()
  const { isLogin, user, logout } = useAuth()

  const [showOverlayMenu, setShowOverlayMenu] = useState<boolean>(false)

  const setLocale = (locale: string) => {
    router.push(
      { pathname: router.pathname, query: router.query },
      { pathname: router.pathname, query: router.query },
      { locale },
    )
  }

  const onSearch = (keyword: string) => {
    let url = {
      pathname: '/species',
      query: { ...router.query, keyword },
    }
    router.push(url, url, { locale: router.locale })
  }

  const logoutDropdown = () => {
    return (
      <Menu onClick={() => logout()}>
        <Menu.Item className="text-base">Logout</Menu.Item>
      </Menu>
    )
  }

  return (
    <nav className="md:px-32 py-4 w-screen bg-primary px-8">
      <div className="flex justify-between items-center align-center">
        <Link href="/">
          <div className="flex m-1 text-white divide-x">
            <img className="h-32 w-32" src="/logo/PTTEP-logo_Sqr.png"/>
            <div className="m-8">
              <div className="ml-2 md:text-3xl">Biodiversity Information Facility</div>
              <div className="ml-2 md:text-2xl">PTTEP-BIF</div>
            </div>
          </div>
        </Link>
        {/* <div className="flex text-white content-end">
          <ul></ul>
          <ul></ul>
          <ul className="md:flex hidden flex-col mt-4 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium">
            {NAVBAR_MENU_LIST.map((item) => (
              <li
                className="text-lg text-white cursor-pointer hover:underline"
                onClick={() => router.push(item.href)}
                key={item[router.locale]}
              >
                {item[router.locale]}
              </li>
            ))}
          </ul>
        </div> */}
        <div className="md:flex justify-around hidden">
          {isLogin ? (
            <Dropdown overlay={logoutDropdown} trigger={['click']}>
              <div className="text-2xl text-white cursor-pointer">
                {user.email}
              </div>
            </Dropdown>
          ) : (
            <div
              className="text-2xl text-white cursor-pointer"
              onClick={() => router.push('/login')}
            >
              {LOGIN_TEXT[router.locale]}
            </div>
          )}
          <div className="ml-8 text-2xl text-white">
            <span
              className={classNames(
                {
                  'underline underline-offset-1': router.locale === 'en',
                },
                'cursor-pointer',
              )}
              onClick={() => setLocale('en')}
            >
              EN
            </span>{' '}
            <span>|</span>
            <span
              className={classNames(
                {
                  'underline underline-offset-1': router.locale === 'th',
                },
                'cursor-pointer',
                'ml-1',
              )}
              onClick={() => setLocale('th')}
            >
              TH
            </span>
          </div>
        </div>
        <div className="md:hidden">
          <MenuOutlined
            className="text-xl font-bold text-white"
            onClick={() => setShowOverlayMenu(true)}
          />
        </div>
      </div>
      <div className="flex md:justify-between justify-center items-center align-center">
        <ul className="md:flex hidden flex-col mt-4 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium">
          {NAVBAR_MENU_LIST.map((item) => (
            <li
              className="text-lg text-white cursor-pointer hover:underline"
              onClick={() => router.push(item.href)}
              key={item[router.locale]}
            >
              {item[router.locale]}
            </li>
          ))}
        </ul>
        <div>
          <Search
            placeholder="input search text"
            allowClear
            css={css`
              width: 250px;
            `}
            onSearch={onSearch}
            defaultValue={router.query.keyword}
          />
        </div>
      </div>
      <OverlayMenu
        visible={showOverlayMenu}
        onClose={() => setShowOverlayMenu(false)}
      />
    </nav>
  )
}

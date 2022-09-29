/** @jsxRuntime classic */
/** @jsx jsx */
import { Divider, Dropdown, Layout, Menu } from 'antd'
import React, { useState } from 'react'
import Link from 'next/link'
import { useRouter } from 'next/router'
import classnames from 'classnames'
import { css, jsx } from '@emotion/react'
import styled from '@emotion/styled'
import { PrimaryButton } from './PrimaryButton'
import { DownOutlined, MenuOutlined } from '@ant-design/icons'
import { OverlayMenu } from './OverlayMenu'
import { APPLICATION_MENU } from './const'
import { LANGUAGE_MAPPER } from 'utils/locale'
import { PAGE_LAYOUT_TEXT } from './locale'
import { useAuth } from 'contexts/AuthContext'

const { Header, Content, Footer } = Layout
const Logo = styled.div`
  font-size: 36px;
  font-weight: bold;
  color: #126e82;
  cursor: pointer;
`
export const LinkItemWrapper = ({ children, href = '#' }) => {
  const router = useRouter()
  return (
    <Link href={href}>
      <a
        className={classnames(
          'h-full px-8 whitespace-no-wrap flex items-center text-lg hover:text-primary border-b-2 border-transparent hover:border-primary text-current',
          {
            'text-primary border-b-2 border-primary': router.pathname === href,
          },
        )}
        css={css`
          transition: color 0.3s cubic-bezier(0.645, 0.045, 0.355, 1),
            border-color 0.3s cubic-bezier(0.645, 0.045, 0.355, 1),
            background 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
        `}
      >
        {children}
      </a>
    </Link>
  )
}

export const PageLayout = ({ children, logoHref = '/' }) => {
  const [showOverlayMenu, setShowOverlayMenu] = useState<boolean>(false)
  const router = useRouter()
  const { isLogin, user, logout } = useAuth()

  const changeLanguage = ({ key }) => {
    router.push(
      { pathname: router.pathname, query: router.query },
      router.pathname,
      { locale: key },
    )
  }
  const languageOverlayMenu = () => {
    return (
      <Menu onClick={changeLanguage}>
        {Object.keys(LANGUAGE_MAPPER).map((lang: string) => (
          <Menu.Item key={lang} className="text-base">
            {LANGUAGE_MAPPER[lang]}
          </Menu.Item>
        ))}
      </Menu>
    )
  }
  const logoutOverlayMenu = () => {
    return (
      <Menu onClick={() => logout()}>
        <Menu.Item key="logout" className="text-base">
          Logout
        </Menu.Item>
      </Menu>
    )
  }
  return (
    <Layout className="min-h-screen bg-gray-100">
      <Header className="bg-white flex justify-between items-center h-auto pt-2 pb-2 px-4 sm:h-24 sm:py-0 sm:px-12 shadow-md z-40">
        <div className="flex items-center">
          <Link href={logoHref}>
            <Logo>eDNA</Logo>
          </Link>
        </div>
        <div className="hidden h-full ml-32 w-full items-center md:flex">
          {APPLICATION_MENU.map((item) => (
            <LinkItemWrapper href={item.href} key={item.href}>
              {item.label[router.locale]}
            </LinkItemWrapper>
          ))}
          <div
            css={css`
              position: absolute;
              right: 0;
              padding-right: 3rem;
            `}
            className="hidden md:flex md:items-center"
          >
            <Dropdown overlay={languageOverlayMenu} trigger={['click']}>
              <div className="flex items-center">
                <div className="text-base mr-2">
                  {LANGUAGE_MAPPER[router.locale]}
                </div>
                <DownOutlined className="text-base" />
              </div>
            </Dropdown>
            <Divider className="h-8 bg-gray-400 ml-4 mr-4" type="vertical" />
            {isLogin ? (
              <Dropdown trigger={['click']} overlay={logoutOverlayMenu}>
                <div className="text-lg font-semibold text-secondary-dark hover:text-secondary cursor-pointer">
                  {user.email}
                </div>
              </Dropdown>
            ) : (
              <div
                className="text-lg font-semibold text-secondary-dark hover:text-secondary cursor-pointer"
                onClick={() => router.push('/login')}
              >
                {PAGE_LAYOUT_TEXT[router.locale].login}
              </div>
            )}
          </div>
        </div>
        <div className="md:hidden">
          <MenuOutlined
            className="text-xl font-semi-bold text-secondary-dark"
            onClick={() => setShowOverlayMenu(true)}
          />
        </div>
      </Header>
      <OverlayMenu
        visible={showOverlayMenu}
        onClose={() => setShowOverlayMenu(false)}
      />
      <Content className="container mx-auto px-4 py-8">{children}</Content>
    </Layout>
  )
}

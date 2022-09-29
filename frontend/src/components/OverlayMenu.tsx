/** @jsxRuntime classic */
/** @jsx jsx */
import { css, jsx } from '@emotion/react'
import styled from '@emotion/styled'
import { Divider, Drawer } from 'antd'
import classnames from 'classnames'
import { useAuth } from 'contexts/AuthContext'
import Link from 'next/link'
import { useRouter } from 'next/router'
import React from 'react'
import { APPLICATION_MENU } from './const'
import { PAGE_LAYOUT_TEXT } from './locale'
import { NAVBAR_MENU_LIST } from './new/locale'
import { LinkItemWrapper } from './PageLayout'
import { PrimaryButton } from './PrimaryButton'

interface OverlayMenuProps {
  visible: boolean
  onClose: () => void
}

const Logo = styled.div`
  font-size: 36px;
  font-weight: bold;
  color: #126e82;
  cursor: pointer;
`

export const OverlayMenu = ({ visible, onClose }: OverlayMenuProps) => {
  const router = useRouter()
  const { isLogin, user, logout } = useAuth()

  const onMenuClick = (href: string) => {
    router.push(href)
    onClose()
  }

  return (
    <Drawer
      className="bg-tertiary"
      visible={visible}
      onClose={onClose}
      css={css`
        .ant-drawer-close {
          font-size: 20px;
          color: #132c33;
        }
        .ant-drawer-body {
          padding-top: 4rem;
        }
        .ant-drawer-content {
          background-color: #d8e3e7;
        }
      `}
    >
      <Logo className="flex justify-center">eDNA</Logo>
      <div className="text-lg font-semibold text-secondary-dark hover:text-secondary cursor-pointer text-center">
        {isLogin && user.email}
      </div>
      <Divider />
      <div className="flex justify-center">
        <div>
          {NAVBAR_MENU_LIST.map((item, idx) => {
            return (
              <div
                className={classnames(
                  { 'text-primary': router.pathname === item.href },
                  'text-xl mt-8 text-center',
                )}
                key={`${idx}${item.href}`}
              >
                <div onClick={() => onMenuClick(item.href)}>
                  {item[router.locale]}
                </div>
              </div>
            )
          })}
        </div>
      </div>
      <div className="flex justify-center mt-8">
        {isLogin ? (
          <PrimaryButton size="large" onClick={() => logout()}>
            {PAGE_LAYOUT_TEXT[router.locale].logout}
          </PrimaryButton>
        ) : (
          <PrimaryButton size="large" onClick={() => router.push('/login')}>
            {PAGE_LAYOUT_TEXT[router.locale].login}
          </PrimaryButton>
        )}
      </div>
    </Drawer>
  )
}

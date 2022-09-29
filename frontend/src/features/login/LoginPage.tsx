/** @jsxRuntime classic */
/** @jsx jsx */
import React from 'react'
import { css, jsx } from '@emotion/react'
import styled from '@emotion/styled'
import { Form, Input, message } from 'antd'
import { LOGIN_PAGE_TEXT } from './locale'
import { useRouter } from 'next/router'
import { PrimaryButton } from 'components/PrimaryButton'
import { login } from 'api/auth/login'
import { eDNAAxios } from 'api/const'
import { useAuth } from 'contexts/AuthContext'

const Logo = styled.div`
  font-size: 48px;
  font-weight: bold;
  color: #126e82;
  cursor: pointer;
`

export const LoginPage = () => {
  const router = useRouter()
  const { login } = useAuth()
  const onFinish = async ({
    email,
    password,
  }: {
    email: string
    password: string
  }) => {
    login({ email, password })
  }
  return (
    <div className="pt-48 bg-tertiary h-screen">
      <div className="text-center">
        <Logo>eDNA</Logo>
        <div className="text-secondary-dark font-bold text-xl">
          {LOGIN_PAGE_TEXT[router.locale].login}
        </div>
      </div>
      <div
        className="mt-10 px-8 mx-auto"
        css={css`
          max-width: 24rem;
        `}
      >
        <Form
          name="loginForm"
          initialValues={{ remember: true }}
          onFinish={onFinish}
        >
          <Form.Item
            name="email"
            rules={[
              {
                required: true,
                message: LOGIN_PAGE_TEXT[router.locale].emailRequiredMessage,
              },
            ]}
          >
            <Input placeholder="email" />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[
              {
                required: true,
                message: LOGIN_PAGE_TEXT[router.locale].passwordRequiredMessage,
              },
            ]}
          >
            <Input.Password placeholder="password" />
          </Form.Item>
          <Form.Item className="mt-8">
            <PrimaryButton type="primary" htmlType="submit" block>
              Submit
            </PrimaryButton>
          </Form.Item>
        </Form>
      </div>
    </div>
  )
}

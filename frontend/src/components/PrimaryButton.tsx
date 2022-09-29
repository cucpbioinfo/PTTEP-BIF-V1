/** @jsxRuntime classic */
/** @jsx jsx */
import React from 'react'
import { Button as AntdButton, ButtonProps } from 'antd'
import { css, jsx } from '@emotion/react'
import classnames from 'classnames'

interface PrimaryButtonProps {
  children?: React.ReactNode
  css?: any
  className?: string
}

export const PrimaryButton = ({
  children,
  css,
  className,
  ...rest
}: PrimaryButtonProps & ButtonProps) => {
  return (
    <AntdButton
      type="primary"
      className={classnames(
        { 'px-12': rest.size === 'large' },
        `bg-primary border-primary hover:bg-white hover:text-primary hover:border-primary rounded-lg ${className}`,
      )}
      css={css}
      {...rest}
    >
      {children}
    </AntdButton>
  )
}

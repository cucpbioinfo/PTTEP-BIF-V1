import { Drawer } from 'antd'
import { PrimaryButton } from 'components/PrimaryButton'
import { SpeciesFilter } from 'components/SpeciesFilter'
import React from 'react'

interface AdvanceSearchDrawerProps {
  visible: boolean
  onClose: () => void
}

export const AdvanceSearchDrawer = ({
  visible,
  onClose,
}: AdvanceSearchDrawerProps) => {
  const onSearch = () => {
    console.log('advance search')
    onClose()
  }

  return (
    <Drawer visible={visible} title="Advance Search" onClose={onClose}>
      <SpeciesFilter />
      <div className="mt-20">
        <PrimaryButton size="large" block onClick={onSearch}>
          Search
        </PrimaryButton>
      </div>
    </Drawer>
  )
}

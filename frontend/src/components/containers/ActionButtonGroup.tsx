import React from "react"
import GetVersionButtonGroup from "../GetVersionButtonGroup"
import { ButtonGroup } from "../ui/button-group"
import { DropdownButtonGroup } from "./DropdownButtonGroup"
import { buttonVariants } from "../ui/button"


type ChangeModeButtonGroupProps = {
  href: string
  label: string
}

const ActionButtonGroup = ({ href, label }: ChangeModeButtonGroupProps) => {
  return (
    <div className="absolute top-6 right-6">
      <ButtonGroup>
        <GetVersionButtonGroup />
        <ButtonGroup className="hidden sm:flex">
          <a href={href} className={buttonVariants()}>
            {label}
          </a>
        </ButtonGroup>

        <DropdownButtonGroup />
      </ButtonGroup>
    </div>
  )
}

export default ActionButtonGroup

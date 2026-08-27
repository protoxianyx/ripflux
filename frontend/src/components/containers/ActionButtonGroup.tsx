import React from "react"
import GetVersionButtonGroup from "../GetVersionButtonGroup"
import { ButtonGroup } from "../ui/button-group"
// import { ChangeModeButtonGroup } from "../ChangeModeButtonGroup"
// import { DropdownButtonGroup } from "./DropdownButtonGroup"

const ActionButtonGroup = () => {
  return (
    <div className="absolute top-6 right-6">
      <ButtonGroup>
        <GetVersionButtonGroup />
        {/* <ChangeModeButtonGroup /> */}
        {/* <DropdownButtonGroup /> */}
      </ButtonGroup>
    </div>
    
  )
}

export default ActionButtonGroup

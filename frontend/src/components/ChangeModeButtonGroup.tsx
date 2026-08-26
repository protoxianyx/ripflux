import React from "react"
import { ButtonGroup } from "./ui/button-group"
import { Button } from "./ui/button"

export const ChangeModeButtonGroup = () => {
  return <ButtonGroup className="hidden sm:flex">
    <Button >
      Advanced Mode
    </Button>
  </ButtonGroup>
}

import React from "react"
import { ButtonGroup } from "../ui/button-group"
import { DropdownMenu, DropdownMenuTrigger } from "../ui/dropdown-menu"
import { Button } from "../ui/button"
import { ListIcon } from "@phosphor-icons/react"

export const DropdownButtonGroup = () => {
  return (
    <ButtonGroup>
      <DropdownMenu>
        <DropdownMenuTrigger
          render={
            <Button variant="outline" size="icon" aria-label="More Options">
              <ListIcon size={32} />
            </Button>
          }
        />
      </DropdownMenu>
    </ButtonGroup>
  )
}

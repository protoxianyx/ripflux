import React from "react"
import { ButtonGroup } from "../ui/button-group"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu"
import { Button } from "../ui/button"
import { ListIcon } from "@phosphor-icons/react"
import { clearTmp } from "@/api/utilsapi"

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
        <DropdownMenuContent>
          <DropdownMenuGroup>
            <DropdownMenuLabel>Utils</DropdownMenuLabel>

            <DropdownMenuItem
              onClick={() => {
                void clearTmp().catch((error) => console.error(error))
              }}
            >
              Clear Tmp
            </DropdownMenuItem>
          </DropdownMenuGroup>
        </DropdownMenuContent>
      </DropdownMenu>
    </ButtonGroup>
  )
}

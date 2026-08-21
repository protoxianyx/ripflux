import React from "react"
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "./ui/dialog"
import { Button } from "./ui/button"

type InvalidPopUpProps = {
  open: boolean
  onOpenChange: (open: boolean) => void
  onUseDefaults: () => void
}



export default function InvalidPopUp({
  open,
  onOpenChange,
  onUseDefaults,
}: InvalidPopUpProps) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent showCloseButton={true}>
        <DialogHeader>
          <DialogTitle>Incomplete download options</DialogTitle>
          <DialogDescription>
            Enter a URL and select a format and resolution before downloading.
          </DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <DialogClose render={<Button variant="outline">Cancel</Button>} />
          <Button type="button" variant="outline" onClick={onUseDefaults}>Set Defaults</Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

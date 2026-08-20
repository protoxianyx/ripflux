import React from 'react'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "./ui/dialog"

type InvalidPopUpProps = {
  open: boolean
  onOpenChange: (open: boolean) => void
}

export default function InvalidPopUp({
  open,
  onOpenChange,
}: InvalidPopUpProps) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Incomplete download options</DialogTitle>
          <DialogDescription>
            Enter a URL and select a format and resolution before downloading.
          </DialogDescription>
        </DialogHeader>

        <DialogFooter showCloseButton />
      </DialogContent>
    </Dialog>
  )
}
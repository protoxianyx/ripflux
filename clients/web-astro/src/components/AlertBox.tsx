import React from "react"
import { Alert, AlertDescription, AlertTitle } from "./ui/alert"
import { InfoIcon } from "@phosphor-icons/react"

interface AlertBoxProps {
  title: string
  description: React.ReactNode
}

export const AlertBox = ({ title, description }: AlertBoxProps) => {
  return (
    <Alert>
      <InfoIcon />
      <AlertTitle>{title}</AlertTitle>
      <AlertDescription>{description}</AlertDescription>
    </Alert>
  )
}

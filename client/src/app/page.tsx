"use client"

import { Button } from '@/components/ui/button'
import React from 'react'

const App = () => {
  return (
    <div> <Button onClick={() => console.log("Nice Click")}> Text</Button></div>
)
}

export default App
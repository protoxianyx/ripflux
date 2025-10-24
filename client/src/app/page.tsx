"use client"

import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import React from 'react'

const App = () => {
  return (
    <div> <Input type='url' placeholder='Image'/> <Button onClick={() => console.log("Nice Click")}> Text</Button></div>
)
}

export default App
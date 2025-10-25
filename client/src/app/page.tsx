"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import React, { useRef, useState } from "react";

const App = () => {
  const [inputValue, setInputValue] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);
  const [jsonData, setJsonData] = useState<object | null>(null);
  const [normaldata, setNormalData] = useState<string>("");

  const handleInput = () => {
    const value = inputRef.current?.value || "";
    const data = { inputValue: value };
    setJsonData(data);
    setNormalData(value);
  };

  const parseJSON = (dataBlock: object | null) => {
    if (dataBlock == null) {
      console.log("NULL");
      return null;
    }
    const data = JSON.stringify(dataBlock, null, 2);
    const parsedData = {
      parsed: data
    }

    console.log(`DataBlock: ${dataBlock} \n Data: ${data} \n`);
    return parsedData;
  };

  const parsedData = () => {
    const data = parseJSON(jsonData)
    console.log("Parsed DAta: ", data)
    
  }

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key == "Enter") {
      handleInput();
    }
  };

  return (
    <div>
      <div>
        <Input
          type="url"
          ref={inputRef}
          placeholder="input"
          onKeyDown={handleKeyDown}
          onInput={(e) => {
            const inputData = e.currentTarget.value;
            setInputValue(inputData);
          }}
        />
        <Button onClick={handleInput}>Enter</Button>
      </div>
      <Separator />
      <div>
        <p>Continious Input Value: {inputValue}</p>
        <p>JSON Data: {parseJSON(jsonData)?.parsed}</p>
        <p>Data: {normaldata} </p>
        <p>Parsed Data:  { }</p>
      </div>
    </div>
  );
};

export default App;

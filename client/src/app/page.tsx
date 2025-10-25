"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import React, { useRef, useState } from "react";

const App = () => {
  interface JSON_Data {
    inputValue: string;
  }

  const [inputValue, setInputValue] = useState("");
  const inputRef = useRef<HTMLInputElement>(null);
  const [jsonData, setJsonData] = useState<JSON_Data | null>(null);
  const [normaldata, setNormalData] = useState<string>("");

  const handleInput = () => {
    const value = inputRef.current?.value || "";
    const data: JSON_Data = { inputValue: value };
    setJsonData(data);
    setNormalData(value);
  };

  const parseJSON = (dataBlock: JSON_Data | null) => {
    if (dataBlock == null) {
      console.log("NULL");
      return null;
    }

    const jsonData: object = {
      jsonInputValue: dataBlock.inputValue,
    };
    const data = JSON.stringify(jsonData);
    const parsedData = {
      parsed: dataBlock.inputValue,
      json: data,
    };

    console.log(`DataBlock: ${dataBlock} \nData: ${data} \n`);
    return parsedData;
  };

  const parsedData = () => {
    const data = parseJSON(jsonData);
    console.log("Parsed DAta: ", data);
  };

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
        <p>JSON Data: {parseJSON(jsonData)?.json}</p>
        <p>Data: {normaldata} </p>
        <p>Parsed Data: {parseJSON(jsonData)?.parsed}</p>
      </div>
    </div>
  );
};

export default App;

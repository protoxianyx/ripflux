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

  const parseJSON = (dataBlock: JSON_Data | null) => {
    if (dataBlock == null) {
      console.log("NULL");
      return null;
    }

    const jsonData: object = {
      jsonInputValue: dataBlock.inputValue,
    };
    const data = JSON.stringify(jsonData, null, 2);

    const parsedData = {
      parsed: dataBlock.inputValue,
      json: data,
    };

    console.log(`DataBlock: ${dataBlock} \nData: ${data} \n`);
    return parsedData;
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key == "Enter") {
      handleInput();
    }
  };

  const sendData = async () => {
    await fetch("http://localhost:8080/api/send", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ text: parseJSON(jsonData)?.parsed }),
    });
  };

  const handleInput = () => {
    const value = inputRef.current?.value || "";
    const data: JSON_Data = { inputValue: value };
    setJsonData(data);
    sendData();
  };

  return (
    <div>
      <div className="flex border-0  border-black m-2">
        <Input
          type="url"
          ref={inputRef}
          placeholder="input"
          onKeyDown={handleKeyDown}
          onInput={(e) => {
            const inputData = e.currentTarget.value;
            setInputValue(inputData);
          }}
          className="m-1"
        />
        <Button
          onClick={() => {
            handleInput;
            // sendData;
          }}
          className="m-1"
        >
          Enter
        </Button>
      </div>
      <Separator />
      <div>
        <p>Continious Input Value: {inputValue}</p>
        <p>JSON Data: {parseJSON(jsonData)?.json}</p>
        <p>Parsed Data: {parseJSON(jsonData)?.parsed}</p>
      </div>
    </div>
  );
};

export default App;

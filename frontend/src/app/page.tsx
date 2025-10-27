"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { useRef, useState } from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

const App = () => {
  interface JSONDataObject {
    inputValue: string;
  }

  const [inputValue, setInputValue] = useState<string>("");
  const [jsonData, setJsonData] = useState<JSONDataObject>();
  const inputRef = useRef<HTMLInputElement | null>(null);

  const sendData = async () => {
    await fetch("http://localhost:8080/api/send", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(jsonData),
    });
    console.log("data sent");
  };

  const handleSubmit = () => {
    const inputData = inputRef.current?.value || "";
    const data: JSONDataObject = { inputValue: inputData };
    setJsonData(data);
  };

  const handleKeyUp = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      handleSubmit();
      sendData();
    }
  };

  return (
    <div>
      <div className="flex border-0  border-black m-2">
        <Input
          className="m-1"
          onInput={(e) => {
            setInputValue(e.currentTarget.value);
          }}
          onKeyUp={handleKeyUp}
          ref={inputRef}
        />
        <Select>
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Quality" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="light">Select 1</SelectItem>
            <SelectItem value="dark">Seleect 2</SelectItem>
            <SelectItem value="system">Select 3</SelectItem>
          </SelectContent>
        </Select>
        <Button
          className="m-1"
          onClick={() => {
            handleSubmit();
            sendData();
          }}
        >
          Enter
        </Button>
      </div>
      <Separator />
      <div>
        <p>Input Value as we type:{inputValue} </p>
        <p>Input Value:{JSON.stringify(jsonData)} </p>
        <p>JSON: </p>
        <p>JSON Data: {jsonData?.inputValue} </p>
      </div>
    </div>
  );
};

export default App;

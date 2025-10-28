"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { useRef, useState } from "react";
import Selector from "@/components/base/selector";

const App = () => {
  interface JSONDataObject {
    inputValue: string;
  }

  const [inputValue, setInputValue] = useState<string>("");
  const [jsonData, setJsonData] = useState<JSONDataObject>();
  const inputRef = useRef<HTMLInputElement | null>(null);
  const [response, setResponse] = useState("");

  const sendToBackend = async () => {
    const inputData = inputRef.current?.value || "";

    const res = await fetch("http://localhost:8080/api/send", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ text: inputData }),
    });

    const data = await res.json();
    setResponse(data.recieved);
  };

  const handleSubmit = () => {
    const inputData = inputRef.current?.value || "";
    const data: JSONDataObject = { inputValue: inputData };
    setJsonData(data);
  };

  const handleKeyUp = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      sendToBackend();
      handleSubmit();
      // sendData();
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
        <Selector />
        <Button
          className="m-1"
          onClick={() => {
            handleSubmit();
            // sendData();
            sendToBackend();
          }}
        >
          Enter
        </Button>
      </div>
      <Separator />
      <div>
        <p>Input Value as we type:{inputValue} </p>
        <p>Input Value:{JSON.stringify(jsonData)} </p>
        <p>Server: {response} </p>
        <p>JSON Data: {jsonData?.inputValue} </p>
      </div>
    </div>
  );
};

export default App;

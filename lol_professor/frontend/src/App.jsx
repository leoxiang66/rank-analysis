import wailsLogo from "./assets/wails.png";
import "./App.css";
import { Info } from "../wailsjs/go/main/App";
import { useState, useEffect } from "react";

function App() {
  return (
    <div className="w-screen h-screen bg-white text-black font-mono overflow-hidden">
      <div className="inline-block bg-blue-100 h-full w-[100px]">
        <div className="flex flex-col gap-10 items-center  h-full">
          <div className="logo w-10 h-10 border border-zinc-400 mt-8">Logo</div>

          <div>Profile</div>
          <div>Game</div>
          <div className="mt-auto mb-8">Others</div>
        </div>
      </div>
    </div>
  );
}

export default App;

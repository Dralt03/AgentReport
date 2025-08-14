"use client";
import { useState, useEffect } from "react";

export default function HomePage() {
  const [isCallActive, setIsCallActive] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [items, setItems] = useState();

  useEffect(() => {
    fetch("http://localhost:8000/items")
      .then((res) => res.json())
      .then((data) => setItems(data))
      .catch((err) => console.error(err));
  }, []);

  console.log(items);

  return (
    <div className="min-h-screen bg-white flex flex-col items-center justify-around">
      <h1 className="text-5xl py-10">WELCOME TO AGENT REPORT</h1>
      <div className="relative flex items-center justify-center mb-12">
        <div
          className={`blob w-32 h-24 bg-gradient-to-br from-purple-400 to-blue-400 shadow-xl opacity-90 ${
            isCallActive ? "animate-pulse" : ""
          }`}
        ></div>

        {isCallActive && (
          <div className="absolute -top-2 -right-2 w-4 h-4 bg-green-500 rounded-full animate-pulse"></div>
        )}
      </div>

      <button
        onClick={() => {
          setIsCallActive(!isCallActive);
        }}
        disabled={isLoading}
        className={`px-8 py-3 text-grey-800 text-2xl rounded-full transform hover:scale-105 transition-all duration-200 shadow-lg hover:shadow-xl ${
          isCallActive
            ? "bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700"
            : "bg-gradient-to-r from-purple-500 to-pink-500 hover:from-purple-600 hover:to-pink-600"
        } ${isLoading ? "opacity-50 cursor-not-allowed" : ""}`}
      >
        {isLoading ? "Connecting..." : isCallActive ? "Disconnect" : "Connect"}
      </button>

      <p className="mt-4 text-gray-600 text-sm">
        {isCallActive
          ? "Voice AI is active - speak now!"
          : "Click Connect to start voice conversation"}
      </p>
    </div>
  );
}

"use client";

import React from "react";
import { useVapi } from "../hooks/useVapi";
import { Mic, MicOff, Loader2 } from "lucide-react";

export default function VoiceAssistant() {
  const { isSessionActive, isConnected, toggleCall, volumeLevel } = useVapi();

  return (
    <div className="p-6 rounded-2xl shadow-lg bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 max-w-sm w-full transition-all duration-300 hover:shadow-xl">
      <div className="flex flex-col items-center gap-4">
        <div className="relative">
          <div
            className={`w-24 h-24 rounded-full flex items-center justify-center transition-all duration-300 ${
              isSessionActive
                ? "bg-red-500 shadow-[0_0_20px_rgba(239,68,68,0.5)]"
                : "bg-blue-600 shadow-[0_0_20px_rgba(37,99,235,0.3)]"
            }`}
          >
            {isSessionActive && isConnected ? (
               <div
                className="absolute inset-0 rounded-full border-4 border-white opacity-30 animate-ping"
                style={{
                  transform: `scale(${1 + volumeLevel * 2})`,
                }}
              />
            ) : null}
            
            <button
              onClick={toggleCall}
              className="z-10 w-full h-full flex items-center justify-center rounded-full focus:outline-none focus:ring-4 focus:ring-offset-2 focus:ring-blue-500"
              aria-label={isSessionActive ? "Stop voice assistant" : "Start voice assistant"}
            >
              {isSessionActive && !isConnected ? (
                 <Loader2 className="w-10 h-10 text-white animate-spin" />
              ) : isSessionActive ? (
                <Mic className="w-10 h-10 text-white" />
              ) : (
                <MicOff className="w-10 h-10 text-white" />
              )}
            </button>
          </div>
        </div>

        <div className="text-center space-y-2">
          <h2 className="text-2xl font-bold text-gray-900 dark:text-white">
            {isSessionActive ? (isConnected ? "Listening..." : "Connecting...") : "Voice Assistant"}
          </h2>
          <p className="text-gray-500 dark:text-gray-400 text-sm">
            {isSessionActive
              ? "Ask me about the latest news!"
              : "Click the microphone to start"}
          </p>
        </div>
      </div>
    </div>
  );
}

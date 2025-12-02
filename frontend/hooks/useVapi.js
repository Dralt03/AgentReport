import { useEffect, useState, useCallback, useRef } from "react";
import Vapi from "@vapi-ai/web";

export const useVapi = () => {
  const [isSessionActive, setIsSessionActive] = useState(false);
  const [isSpeechActive, setIsSpeechActive] = useState(false);
  const [isConnected, setIsConnected] = useState(false);
  const [volumeLevel, setVolumeLevel] = useState(0);
  const vapiRef = useRef(null);

  useEffect(() => {
    const vapi = new Vapi(process.env.NEXT_PUBLIC_VAPI_PUBLIC_KEY);
    vapiRef.current = vapi;

    const onCallStart = () => {
      setIsSessionActive(true);
      setIsConnected(true);
    };

    const onCallEnd = () => {
      setIsSessionActive(false);
      setIsConnected(false);
      setIsSpeechActive(false);
      setVolumeLevel(0);
    };

    const onSpeechStart = () => setIsSpeechActive(true);
    const onSpeechEnd = () => setIsSpeechActive(false);

    const onVolumeLevel = (level) => setVolumeLevel(level);

    const onMessage = (message) => {
      console.log("Vapi Message:", message);
    };

    const onError = (error) => {
      console.error("Vapi Error:", error);
      setIsSessionActive(false);
      setIsConnected(false);
    };

    vapi.on("call-start", onCallStart);
    vapi.on("call-end", onCallEnd);
    vapi.on("speech-start", onSpeechStart);
    vapi.on("speech-end", onSpeechEnd);
    vapi.on("volume-level", onVolumeLevel);
    vapi.on("message", onMessage);
    vapi.on("error", onError);

    return () => {
      vapi.stop();
      vapi.off("call-start", onCallStart);
      vapi.off("call-end", onCallEnd);
      vapi.off("speech-start", onSpeechStart);
      vapi.off("speech-end", onSpeechEnd);
      vapi.off("volume-level", onVolumeLevel);
      vapi.off("message", onMessage);
      vapi.off("error", onError);
    };
  }, []);

  const start = useCallback(async () => {
  const vapi = vapiRef.current;
  if (!vapi) return;

  try {
    // Register the tool BEFORE vapi.start()
    if (vapi.tools) {
      vapi.tools.add({
        name: "get_news",
        description:
          "Fetch the latest scraped news articles. Returns a JSON array.",
        async execute() {
          try {
            const res = await fetch("https://agentreport.onrender.com/items");

            if (!res.ok) {
              return { success: false, error: `HTTP ${res.status}` };
            }

            const data = await res.json();

            return {
              success: true,
              articles: data,
            };
          } catch (err) {
            console.error("Error fetching news:", err);
            return {
              success: false,
              error: "Failed to fetch news.",
            };
          }
        },
      });
    } else {
      console.warn("vapi.tools is undefined. Skipping tool registration.");
    }

      const assistantId = process.env.NEXT_PUBLIC_VAPI_ASSISTANT_ID;
      if (!assistantId) {
        console.error("Missing Vapi Assistant ID. Please set NEXT_PUBLIC_VAPI_ASSISTANT_ID in your .env file.");
        return;
      }

      await vapi.start(assistantId);
    } catch (error) {
      console.error("Failed to start Vapi session:", error);
    }
  }, []);

  const stop = useCallback(() => {
    vapiRef.current?.stop();
  }, []);

  const toggleCall = useCallback(() => {
    if (isSessionActive) {
      stop();
    } else {
      start();
    }
  }, [isSessionActive, start, stop]);

  return {
    isSessionActive,
    isSpeechActive,
    isConnected,
    volumeLevel,
    toggleCall,
    start,
    stop,
  };
};

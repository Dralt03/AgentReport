"use client";
import { useState } from "react";

export default function HomePage() {
  const [isCallActive, setIsCallActive] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const vapiRef = useRef(null);

  useEffect(() => {
    const initVapi = async () => {
      if (typeof window !== "undefined") {
        const Vapi = (await import("@vapi-ai/web")).default;
        vapiRef.current = new Vapi(
          process.env.NEXT_PUBLIC_VAPI_PUBLIC_KEY || "your-vapi-public-key"
        );

        // Set up event listeners
        vapiRef.current.on("call-start", () => {
          setIsCallActive(true);
          setIsLoading(false);
        });

        vapiRef.current.on("call-end", () => {
          setIsCallActive(false);
          setIsLoading(false);
        });

        vapiRef.current.on("error", (error) => {
          console.error("Vapi error:", error);
          setIsLoading(false);
        });
      }
    };

    initVapi();

    return () => {
      if (vapiRef.current) {
        vapiRef.current.stop();
      }
    };
  }, []);

  const handleConnect = async () => {
    if (!vapiRef.current) return;

    if (isCallActive) {
      vapiRef.current.stop();
    } else {
      setIsLoading(true);
      try {
        await vapiRef.current.start({
          model: {
            provider: "openai",
            model: "gpt-3.5-turbo",
            messages: [
              {
                role: "system",
                content:
                  "You are a helpful AI assistant. Keep responses concise and friendly.",
              },
            ],
          },
          voice: {
            provider: "11labs",
            voiceId: "21m00Tcm4TlvDq8ikWAM",
          },
        });
      } catch (error) {
        console.error("Failed to start call:", error);
        setIsLoading(false);
      }
    }
  };

  return (
    <div className="min-h-screen bg-white flex flex-col items-center justify-center">
      <div className="relative flex items-center justify-center mb-12">
        {/* Centered Blob with call indicator */}
        <div
          className={`blob w-32 h-24 bg-gradient-to-br from-purple-400 via-pink-400 via-cyan-400 to-blue-400 shadow-xl opacity-90 ${
            isCallActive ? "animate-pulse" : ""
          }`}
        ></div>

        {isCallActive && (
          <div className="absolute -top-2 -right-2 w-4 h-4 bg-green-500 rounded-full animate-pulse"></div>
        )}
      </div>

      <button
        onClick={handleConnect}
        disabled={isLoading}
        className={`px-8 py-3 text-black font-semibold rounded-full transform hover:scale-105 transition-all duration-200 shadow-lg hover:shadow-xl ${
          isCallActive
            ? "bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700"
            : "bg-gradient-to-r from-purple-500 to-pink-500 hover:from-purple-600 hover:to-pink-600"
        } ${isLoading ? "opacity-50 cursor-not-allowed" : ""}`}
      >
        {isLoading ? "Connecting..." : isCallActive ? "Disconnect" : "Connect"}
      </button>

      {/* Status text */}
      <p className="mt-4 text-gray-600 text-sm">
        {isCallActive
          ? "Voice AI is active - speak now!"
          : "Click Connect to start voice conversation"}
      </p>
    </div>
  );
}

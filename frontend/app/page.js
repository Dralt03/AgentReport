"use client";
import VoiceAssistant from "@/components/voiceAssistant";

export default function HomePage() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-50 dark:bg-gray-900 p-4">
      <div className="text-center mb-12">
        <h1 className="text-4xl font-extrabold text-gray-900 dark:text-white mb-4 tracking-tight">
          AgentReport
        </h1>
        <p className="text-lg text-gray-600 dark:text-gray-300 max-w-md mx-auto">
          Your AI-powered news aggregator. Listen to summaries and ask questions about current events.
        </p>
      </div>
      
      <VoiceAssistant />
    </div>
  );
}

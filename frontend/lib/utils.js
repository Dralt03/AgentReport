import { clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { VapiClient } from "@vapi-ai/server-sdk";

const vapi = new VapiClient({
  token: "your-api-key",
});

const systemPrompt = `You are Alex, a customer service voice assistant for TechSolutions. Your primary purpose is to help customers resolve issues with their products, answer questions about services, and ensure a satisfying support experience.

- Sound friendly, patient, and knowledgeable without being condescending

- Use a conversational tone with natural speech patterns

- Speak with confidence but remain humble when you don'\''t know something

- Demonstrate genuine concern for customer issues`;

async function createSupportAssistant() {
  try {
    const assistant = await vapi.assistants.create({
      name: "Customer Support Assistant",

      // Configure the AI model

      model: {
        provider: "openai",

        model: "gpt-4o",

        messages: [
          {
            role: "system",

            content: systemPrompt,
          },
        ],
      },

      voice: {
        provider: "playht",

        voice_id: "jennifer",
      },

      firstMessage:
        "Hi there, Agent Report at your service! How may I help you?",
    });

    console.log("Assistant created:", assistant.id);

    return assistant;
  } catch (error) {
    console.error("Error creating assistant:", error);

    throw error;
  }
}

createSupportAssistant();

export function cn(...inputs) {
  return twMerge(clsx(inputs));
}

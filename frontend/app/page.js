"use client";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Zap, ArrowRight } from "lucide-react";
import { DotLottieReact } from "@lottiefiles/dotlottie-react";

async function handleLogin() {}

export default function Home() {
  return (
    <div className="w-screen h-screen">
      <div className="flex h-16 w-full items-center justify-between px-4 md:px-6">
        <div className="flex items-center w-full">
          <div className="flex h-8 w-8 mx-2 items-center justify-center rounded-lg bg-primary">
            <Zap className="h-5 w-5 text-primary-foreground" />
          </div>
          <span className="text-xl font-bold">Agent Report</span>
        </div>
        <div className="flex items-center space-x-4">
          <Button variant="ghost" size="sm" className="hidden md:inline-flex">
            Sign In
          </Button>
          <Button size="sm">Get Started</Button>
        </div>
      </div>
      <section className="py-12 px-20 md:py-24">
        <div className="px-4 md:px-6 flex justify-between items-center">
          <div className="flex flex-col justify-center space-y-4">
            <div className="space-y-2">
              <Badge variant="secondary" className="w-fit">
                New: AI-Powered Workflows
              </Badge>
              <h1 className="text-5xl font-bold xl:text-6xl/none">
                Agent Report Ready for Ease and Convinience
              </h1>
              <p className="max-w-[600px] text-muted-foreground py-8 md:text-xl">
                Your personal AI assistant to help you save time and effort so
                that you can focus on the more important stuff
              </p>
            </div>
            <div className="flex flex-col gap-2 min-[400px]:flex-row">
              <Button size="lg" className="h-12 px-8">
                Sign Up
                <ArrowRight className="ml-2 h-4 w-4" />
              </Button>
            </div>
          </div>
          <DotLottieReact
            className="max-md:hidden"
            src="https://lottie.host/2fc1d95f-fca1-48d6-867f-45adbd454eaf/mmRzidBQhn.lottie"
            loop
            autoplay
          />
        </div>
      </section>
    </div>
  );
}

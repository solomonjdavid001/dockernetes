"use client";

import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { SidebarTrigger } from "@/components/ui/sidebar";
import { Moon, Sun } from "lucide-react";
import { useTheme } from "next-themes";

export function Header() {
  const { resolvedTheme, setTheme } = useTheme();

  const toggleTheme = () => {
    setTheme(resolvedTheme === "dark" ? "light" : "dark");
  };

  return (
    <header className="flex h-16 shrink-0 items-center justify-between gap-2 px-4 w-full">
      <div className="flex items-center gap-2">
        <SidebarTrigger className="-ml-1" />
        <Separator orientation="vertical" className="mr-2 h-4" />
      </div>
      <Button
        variant="outline"
        size="icon"
        onClick={toggleTheme}
        className="relative ml-auto"
      >
        <Sun
          className={`h-[1.2rem] w-[1.2rem] transition-all ${
            resolvedTheme === "dark"
              ? "rotate-90 scale-0"
              : "rotate-0 scale-100"
          }`}
        />
        <Moon
          className={`absolute h-[1.2rem] w-[1.2rem] transition-all ${
            resolvedTheme === "dark"
              ? "rotate-0 scale-100"
              : "-rotate-90 scale-0"
          }`}
        />
        <span className="sr-only">Toggle theme</span>
      </Button>
    </header>
  );
}

"use client";

import { Separator } from "@/components/ui/separator";
import { SidebarTrigger } from "@/components/ui/sidebar";
import { ModeToggle } from "../mode-toggle";

export function Header() {
  return (
    <header className="flex h-16 shrink-0 items-center justify-between gap-2 px-4 w-full">
      <div className="flex items-center gap-2">
        <SidebarTrigger className="-ml-1" />
        <Separator orientation="vertical" className="mr-2 h-4" />
      </div>
      <ModeToggle />
    </header>
  );
}

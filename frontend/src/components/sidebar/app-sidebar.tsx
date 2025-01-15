"use client";

import * as React from "react";
import { NavUser } from "@/components/sidebar/nav-user";
import {
  Sidebar,
  SidebarFooter,
  SidebarHeader,
  SidebarRail,
} from "@/components/ui/sidebar";
import { LogoAndNameDisplay } from "./logo-and-name";
import Image from "next/image";
import { sidebarData } from "./sidbar-data";
import { NavList } from "./nav-list";

// Function to generate initials from the user's name
const getInitials = (name: string) => {
  const nameParts = name.split(" ");
  const initials = nameParts
    .map((part) => part.charAt(0).toUpperCase())
    .join("");
  return initials;
};

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  const { user } = sidebarData;

  const avatarContent = user.avatar ? (
    <Image
      src={user.avatar}
      alt="User Avatar"
      className="w-8 h-8 rounded-full"
    />
  ) : (
    <div className="w-8 h-8 flex items-center justify-center bg-gray-500 text-white rounded-full">
      {getInitials(user.name)}
    </div>
  );

  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <LogoAndNameDisplay />
      </SidebarHeader>
      <NavList />
      <SidebarFooter>
        <NavUser user={user} avatarContent={avatarContent} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  );
}

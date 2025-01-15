import {
  ChartSpline,
  Container,
  HardDrive,
  Network,
  ShieldHalf,
  Wrench,
  Vault,
} from "lucide-react";

export const sidebarData = {
  user: {
    name: "Solomon David",
    email: "solomonjdavid@gmail.com",
    avatar: "",
  },
  navList: [
    {
      category: "Local",
      items: [
        { title: "Images", url: "/images", icon: Vault, isActive: true },
        { title: "Containers", url: "/containers", icon: Container },
        { title: "Volumes", url: "/volumes", icon: HardDrive },
        { title: "Builds", url: "/builds", icon: Wrench },
      ],
    },
    {
      category: "Cloud",
      items: [
        { title: "Docker Hub", url: "#", icon: Network },
        { title: "Docker Scout", url: "#", icon: ShieldHalf },
      ],
    },
    {
      category: "System",
      items: [{ title: "Usage", url: "/dashboard", icon: ChartSpline }],
    },
  ],
};

"use client";

import { Button } from "@/components/ui/button";
import { useRouter } from "next/navigation";

export default function Home() {
  const reload = () => {
    router.push("/images");
  };
  const router = useRouter();
  return (
    <div className="flex justify-center items-center min-h-screen">
      <Button onClick={() => reload()}>Go to Dashboard</Button>
    </div>
  );
}

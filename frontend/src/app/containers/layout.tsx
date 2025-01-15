import { Sidebar } from "@/components/sidebar/sidebar";

export default function ImagesLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <Sidebar>{children}</Sidebar>;
}

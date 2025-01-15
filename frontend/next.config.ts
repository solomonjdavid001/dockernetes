import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  output: "export",
  distDir: "dist",
  compress: false,
  trailingSlash: true,
};

export default nextConfig;

// https://nextjs.org/docs/app/guides/static-exports#configuration
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  /* config options here */
  reactCompiler: true,
  reactStrictMode: true,
  trailingSlash: true,
  output: 'export',
};

export default nextConfig;

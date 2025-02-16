import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: true,

  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: "https://blog-api-production-2267.up.railway.app/:path*",
      },
      {
        source: "/login",
        destination: "https://blog-api-production-2267.up.railway.app/login",
      },
    ];
  },
};

export default nextConfig;
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: true,

  images: {
    domains: ["ik.imagekit.io"],
  },

  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: "http://localhost:8080/:path*",
      },
      {
        source: "/login",
        destination: "http://localhost:8080/login",
      },
    ];
  },
};

export default nextConfig;

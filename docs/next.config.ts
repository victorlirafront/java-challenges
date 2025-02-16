import type { NextConfig } from "next";

const API_BASE_URL =
  process.env.NEXT_PUBLIC_BLOG_API_DEVELOPMENT ||
  process.env.NEXT_PUBLIC_BLOG_API_PRODUCTION;

if (!API_BASE_URL) {
  throw new Error("API_BASE_URL está indefinido. Verifique suas variáveis de ambiente.");
}

const nextConfig: NextConfig = {
  reactStrictMode: true,

  images: {
    domains: ["ik.imagekit.io"],
  },

  async rewrites() {
    return [
      {
        source: "/api/:path*",
        destination: `${API_BASE_URL}/:path*`,
      },
    ];
  },
};

export default nextConfig;

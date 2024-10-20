import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  async redirects() {
    return [
      {
        source: "/",
        destination: "/folders",
        permanent: false,
      },
    ];
  },
  /* config options here */
};

export default nextConfig;

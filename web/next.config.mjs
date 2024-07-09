/** @type {import('next').NextConfig} */

const nextConfig = {
    output: 'standalone',
    images: {unoptimized: true},
    env: {
        API_URL: process.env.API_URL,
    },
}

// module.exports = nextConfig
export default nextConfig;


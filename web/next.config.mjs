const nextConfig = {
    output: 'standalone',
    images: {unoptimized: true},
    env: {
        API_URL: process.env.API_URL,
    },
}

export default nextConfig;


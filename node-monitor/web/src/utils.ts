export const getApiServerUrl = () => {
    // TODO: get the url from .env.local file
    // const api_server_url = "http://localhost:3001";
    const api_server_url = process.env.NEXT_PUBLIC_API_SERVER_URL;
    return api_server_url;
}
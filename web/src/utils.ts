export const getApiServerUrl = () => {
    // TODO: get the url from .env.local file
    // const api_server_url = "http://localhost:3001";
    const api_server_url = process.env.NEXT_PUBLIC_API_SERVER_URL;
    return api_server_url;
}

export function unixTimestampToDateFormat(timestamp: number) {
    // 2006-01-02 15:04:05
    // Create a new Date object, multiplying by 1000 to convert seconds to milliseconds
    const date = new Date(timestamp * 1000);

    // Extract date components
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-indexed
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');

    // Combine components into desired format
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}


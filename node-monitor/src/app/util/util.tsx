export function unixTimestampToDateFormat(timestamp: number) {
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


const rawApiUrl = process.env.REACT_APP_API_URL || "http://localhost:3080";

// Prevent accidental double slashes when env vars include a trailing slash.
export const API_URL = rawApiUrl.replace(/\/+$/, "");

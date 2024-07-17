import axios, { AxiosRequestConfig, Method } from "axios";

export interface ClientOptions {
  address?: string;
  apiKey?: string;
}

export async function call<T>(
  options: ClientOptions,
  endpoint: string,
  data?: any,
  method: Method = "POST"
): Promise<T> {
  if (!options.address) {
    options.address = "127.0.0.1:58231";
  }
  const url = `${options.address}${endpoint}`;
  const headers = {
    Authorization: `Bearer ${options.apiKey}`,
  };

  const config: AxiosRequestConfig = {
    url,
    method,
    headers,
    data,
  };

  try {
    const response = await axios(config);
    return response.data as T;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.error(
        "Error:",
        error.response ? error.response.data : error.message
      );
    } else {
      console.error("Unexpected Error:", error);
    }
    throw error;
  }
}

export function uuid() {
  return (
    generateSegment(8) +
    "-" +
    generateSegment(4) +
    "-" +
    generateSegment(4) +
    "-" +
    generateSegment(4) +
    "-" +
    generateSegment(12)
  );
}

function generateSegment(length: number) {
  return Array.from({ length: length }, () =>
    Math.floor(Math.random() * 16).toString(16)
  ).join("");
}

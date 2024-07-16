import axios, { AxiosRequestConfig, Method } from "axios";

export interface ClientOptions {
    address?: string;
    apiKey?: string;
  }

  
export async function call<T>(
  address: string,
  apiKey: string,
  endpoint: string,
  method: Method,
  data?: any
): Promise<T> {
  const url = `${address}${endpoint}`;
  const headers = {
    Authorization: `Bearer ${apiKey}`,
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

import axios, { AxiosRequestConfig, Method } from "axios";

export interface Environment {
  production: boolean;
  brandName: string;
  shortBrandName: string;
  backendAddress: string;
  localPromptAddress: string;
  localtronAddress: string;
}

export interface ClientOptions {
  address?: string;
  apiKey?: string;
  env?: Environment;
}

export async function call<T>(
  options: ClientOptions,
  endpoint: string,
  data?: any,
  method: Method = "POST"
): Promise<T> {
  const address = options.address || options.env?.localtronAddress || "http://127.0.0.1:58231";
  const url = `${address}${endpoint}`;
  const headers: Record<string, string> = {};

  if (options.apiKey) {
    headers.Authorization = `Bearer ${options.apiKey}`;
  }

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
      console.error("Error:", error.response ? error.response.data : error.message);
    } else {
      console.error("Unexpected Error:", error);
    }
    throw error;
  }
}

export function get<T>(options: ClientOptions, endpoint: string): Promise<T> {
  return call<T>(options, endpoint, undefined, "GET");
}

export function post<T>(options: ClientOptions, endpoint: string, data: any): Promise<T> {
  return call<T>(options, endpoint, data, "POST");
}

export function put<T>(options: ClientOptions, endpoint: string, data: any): Promise<T> {
  return call<T>(options, endpoint, data, "PUT");
}

export function deleteRequest<T>(options: ClientOptions, endpoint: string, data?: any): Promise<T> {
  return call<T>(options, endpoint, data, "DELETE");
}

export function uuid(): string {
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

function generateSegment(length: number): string {
  return Array.from({ length: length }, () =>
    Math.floor(Math.random() * 16).toString(16)
  ).join("");
}

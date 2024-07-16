import { Method } from "axios";
export interface ClientOptions {
    address?: string;
    apiKey?: string;
}
export declare function call<T>(address: string, apiKey: string, endpoint: string, method: Method, data?: any): Promise<T>;
export declare function uuid(): string;

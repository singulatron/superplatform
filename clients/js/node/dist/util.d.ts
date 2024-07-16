import { Method } from "axios";
export interface ClientOptions {
    address?: string;
    apiKey?: string;
}
export declare function call<T>(address: string, apiKey: string, endpoint: string, data?: any, method?: Method): Promise<T>;
export declare function uuid(): string;

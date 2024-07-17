import { Method } from "axios";
export interface ClientOptions {
    address?: string;
    apiKey?: string;
}
export declare function call<T>(options: ClientOptions, endpoint: string, data?: any, method?: Method): Promise<T>;
export declare function uuid(): string;

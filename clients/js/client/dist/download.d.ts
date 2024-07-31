import { ClientOptions } from "./util";
import * as download from "@singulatron/types";
export declare class DownloadService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    do(url: string): Promise<void>;
    pause(url: string): Promise<void>;
    list(): Promise<download.ListResponse>;
    get(url: string): Promise<download.GetResponse>;
}

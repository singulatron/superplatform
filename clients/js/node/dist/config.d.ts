import { ClientOptions } from "./util";
import * as config from "@singulatron/types";
export declare class ConfigService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    configGet(): Promise<config.ConfigGetResponse>;
}

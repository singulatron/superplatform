import { ClientOptions } from "./util";
import * as docker from "@singulatron/types";
export declare class DockerService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    dockerInfo(): Promise<docker.DockerInfoResponse>;
}

import { ClientOptions } from "./util";
import * as prompt from "@singulatron/types";
export declare class PromptService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    promptAdd(prompt: prompt.AddPromptRequest): Promise<void>;
    promptRemove(promptId: string): Promise<void>;
    promptList(request: prompt.ListPromptsRequest): Promise<prompt.ListPromptsResponse>;
}

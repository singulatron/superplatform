import { Query } from "./generic";
export type PromptStatus = "scheduled" | "running" | "completed" | "errored" | "abandoned" | "canceled";
export interface PromptCreateFields {
    id: string;
    /**
     ThreadId is the ID of the thread a prompt belongs to.
       Clients subscribe to Thread Streams to see the answer to a prompt,
       or set `prompt.sync` to true for a blocking answer.
    */
    threadId: string;
    prompt: string;
    template?: string;
    /** ModelId is just the Singulatron internal ID of the model. */
    modelId?: string;
    /** MaxRetries specified how many times the system should retry a prompt when it keeps erroring. */
    maxRetries?: number;
    /** Sync drives whether prompt add request should wait and hang until
       the prompt is done executing. By default the prompt just gets put on a queue
       and the client will just subscribe to a Thread Stream.
       For quick and dirty scripting however it's often times easier to do things syncronously.
       In those cases set Sync to true.
     */
    sync?: boolean;
}
export interface Prompt extends PromptCreateFields {
    characterId?: string;
    createdAt?: string;
    status?: PromptStatus;
    lastRun?: string;
    runCount?: number;
    error?: string;
    userId?: string;
}
export interface AddPromptRequest extends PromptCreateFields {
}
export interface AddPromptResponse {
    answer?: string;
    prompt?: Prompt;
}
export interface RemovePromptRequest {
    promptId: string;
}
export interface ListPromptsRequest {
    query?: Query;
}
export interface ListPromptsResponse {
    prompts: Prompt[];
    after: string;
    count?: number;
}
export interface PromptRequest {
    prompt: string;
    stream?: boolean;
    max_tokens?: number;
}
export interface CompletionChoice {
    text: string;
    index: number;
    logprobs: any;
    finish_reason: string;
}
export interface CompletionUsage {
    prompt_tokens: number;
    completion_tokens: number;
    total_tokens: number;
}
export interface CompletionResponse {
    id: string;
    object: string;
    created: number;
    model: string;
    choices: CompletionChoice[];
    usage: CompletionUsage;
}

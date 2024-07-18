import { Query } from "./generic";
export type PromptStatus = "scheduled" | "running" | "completed" | "errored" | "abandoned" | "canceled";
export interface Prompt {
    id: string;
    threadId: string;
    userId?: string;
    prompt: string;
    template?: string;
    characterId?: string;
    modelId?: string;
    createdAt?: string;
    status?: PromptStatus;
    lastRun?: string;
    runCount?: number;
    error?: string;
    maxRetries?: number;
    sync?: boolean;
}
export interface AddPromptRequest {
    prompt: Prompt;
}
export interface RemovePromptRequest {
    prompt: Prompt;
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

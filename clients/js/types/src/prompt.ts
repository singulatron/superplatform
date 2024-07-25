import { Query } from "./generic";

export type PromptStatus =
  | "scheduled"
  | "running"
  | "completed"
  | "errored"
  | "abandoned"
  | "canceled";

export interface PromptCreateFields {
  id: string;
  threadId: string;
  userId?: string;
  /*
		Prompt without template, eg. `What is a banana`?
	*/
  prompt: string;
  /*
		Prompt template (eg. `[INST]{prompt}[/INST]`)
		Optional. Might be derived from/through the modelId
	*/
  template?: string;

  modelId?: string;

  maxRetries?: number;
  sync?: boolean;
}

export interface Prompt extends PromptCreateFields {
  characterId?: string;

  createdAt?: string;
  status?: PromptStatus;
  lastRun?: string;
  runCount?: number; // How many times this was ran (retries are due to errors)
  error?: string;
}

export interface AddPromptRequest extends PromptCreateFields {}

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

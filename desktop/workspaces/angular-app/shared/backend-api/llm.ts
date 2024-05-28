/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
export const promptEndpoint = '/llm/prompt';

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

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
import { Injectable, Inject, InjectionToken } from '@angular/core';

import { HttpClient, HttpHeaders } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';

import { CookieService } from 'ngx-cookie-service';
import { catchError } from 'rxjs/operators';
import { throwError } from 'rxjs';
import { BehaviorSubject } from 'rxjs';

export interface Event {
	type: string;
}

export interface FilterSaveEvent extends Event {
	type: 'filter-save';
	filters: string[];
}

export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
	backendAddress: string;
	localPromptAddress: string;
}

export interface ApiServiceConfig {
	env: Environment;
}
export const API_SERVICE_CONFIG = new InjectionToken<ApiServiceConfig>(
	'ApiServiceConfig'
);

export interface ChatThread {
	id: string;
	name: string;
	messages?: Array<ChatMessage>;
}

export const defaultThreadName = 'New chat';

export interface ChatMessage {
	id: string;
	messageContent?: string;
	isUserMessage?: boolean;
}

@Injectable({
	providedIn: 'root',
})
export class ApiService {
	private locale = 'en';

	public firehose: BehaviorSubject<Event> = new BehaviorSubject<Event>({
		type: 'noop',
	});

	private headers: HttpHeaders;
	private config: ApiServiceConfig;

	constructor(
		private http: HttpClient,
		private cs: CookieService,
		@Inject(API_SERVICE_CONFIG) config: ApiServiceConfig
	) {
		this.headers = new HttpHeaders();
		this.config = config;
	}

	public setLocale(s: string) {
		this.locale = s;
	}

	public getLocale(): string {
		return this.locale;
	}

	call(path: string, request: any): Promise<any> {
		let uri = this.config.env.backendAddress + path;

		let body = JSON.stringify(request);

		let headers = this.headers.set(
			'Authorization',
			'Bearer ' + this.cs.get('the_token')
		);

		return firstValueFrom(
			this.http
				.post<any>(uri, body, {
					headers: headers,
					responseType: 'text' as 'json',
				})
				.pipe(
					map((response) => {
						return JSON.parse(response);
					}),
					catchError((error) => {
						if (error.status >= 400) {
							throw error.error;
						}
						return throwError(error);
					})
				)
		);
	}

	getModelsFromServer(): Promise<ModelResponse> {
		return this.call('/app/models', {});
	}

	getVersion(): Promise<VersionResponse> {
		return this.call('/app/version', {});
	}

	modelsFromServer = [];
	serverModelsChecked = false;

	async getModels(): Promise<Model[]> {
		if (!this.serverModelsChecked) {
			this.serverModelsChecked = true;

			try {
				let rsp = await this.getModelsFromServer();
				if (rsp.models && Array.isArray(rsp.models) && rsp.models.length > 0) {
					this.modelsFromServer = rsp.models as any;
				}
			} catch (e) {
				console.error('Error getting models from server', e);
			}
		}

		models.push(...this.modelsFromServer);
		return models;
	}
}

export interface ReadByWebsitesRequest {
	host: string;
}

export interface VersionResponse {
	windows?: Version;
	linux?: Version;
	mac?: Version;
}

export interface Version {
	version?: string;
	downloadPageURL?: string;
	downloadURL?: string;
	releaseDate?: Date | string;
	changeLog?: string;
}

export interface ModelResponse {
	models: Model[];
}

export interface Model {
	/** id is the download url of the model */
	id: string;
	name: string;
	parameters?: string;
	flavour?: string;
	version?: string;
	quality?: string;
	extension?: string;
	fullName?: string;
	tags?: string[];
	mirrors?: string[];
	size?: number;
	uncensored?: boolean;
	maxRam?: number;
	description?: string;
	promptTemplate?: string;
	quantComment?: string;
	maxBits?: number;
	bits?: number;
}

// HERE BE DRAGONS
// I TOLD YOU

const mistralDescription = `Mistral excels in understanding and generating human-like text, making it a versatile tool across a multitude of domains. Its proficiency extends from generating coherent and contextually relevant text passages to providing detailed answers to queries, showcasing an impressive grasp of knowledge across a wide array of subjects.
Mistral stands out for its ability to perform tasks with remarkable accuracy and fewer resources, a leap forward in making state-of-the-art AI more accessible and sustainable.
`;

const codellamaDescription = `CodeLlama is a powerful AI model that specializes in generating code snippets and providing detailed explanations for programming-related queries. It is designed to assist developers in writing code, debugging, and understanding complex programming concepts.`;

const llamaChatUncensoredPrompt = `### HUMAN:
{prompt}
      
### RESPONSE:
`;

export let models: Model[] = [
	//
	// MISTRAL 7B
	//
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q2_K',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q2_K',
		size: 3.08,
		maxRam: 5.58,
		quantComment:
			'smallest, significant quality loss - not recommended for most purposes',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q3_K_S',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q3_K_S',
		size: 3.16,
		maxRam: 5.66,
		quantComment: 'very small, high quality loss',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_M.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q3_K_M',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q3_K_M',
		size: 3.52,
		maxRam: 6.02,
		quantComment: 'very small, high quality loss',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_L.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q3_K_L',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q3_K_L',
		size: 3.82,
		maxRam: 6.32,
		quantComment: 'small, substantial quality loss',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q4_K_S.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q4_K_S',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q4_K_S',
		size: 4.14,
		maxRam: 6.64,
		quantComment: 'small, greater quality loss',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q4_K_M.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q4_K_M',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q4_K_M',
		size: 4.37,
		maxRam: 6.87,
		description: mistralDescription,
		quantComment: 'medium, balanced quality - recommended',
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q5_K_S.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q5_K_S',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q5_K_S',
		size: 5,
		maxRam: 7.5,
		description: mistralDescription,
		quantComment: 'large, very low quality loss - recommended',
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q5_K_M.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q5_K_M',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q5_K_M',
		size: 5.13,
		maxRam: 7.63,
		quantComment: 'large, very low quality loss - recommended',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q6_K.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q6_K',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q6_K',
		size: 5.94,
		maxRam: 8.44,
		quantComment: 'very large, extremely low quality loss',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	{
		id: 'https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q8_0.gguf',
		name: 'Mistral',
		parameters: '7B',
		flavour: 'Instruct',
		version: 'v0.2',
		quality: 'Q8_0',
		extension: 'GGUF',
		fullName: 'Mistral 7B Instruct v0.2 Q8_0',
		size: 7.7,
		maxRam: 10.2,
		quantComment: 'very large, extremely low quality loss - not recommended',
		description: mistralDescription,
		promptTemplate: '[INST] {prompt} [/INST]',
	},
	//
	// CodeLLAMA 7B
	//
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q2_K.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q2_K',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q2_K',
		size: 2.83,
		maxRam: 5.33,
		quantComment:
			'smallest, significant quality loss - not recommended for most purposes',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q3_K_S.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q3_K_S',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q3_K_S',
		size: 2.95,
		maxRam: 5.45,
		quantComment: 'very small, high quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q3_K_M.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q3_K_M',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q3_K_M',
		size: 3.3,
		maxRam: 5.8,
		quantComment: 'very small, high quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q3_K_L.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q3_K_L',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q3_K_L',
		size: 3.6,
		maxRam: 6.1,
		quantComment: 'small, substantial quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q4_K_S.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q4_K_S',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q4_K_S',
		size: 3.86,
		maxRam: 6.36,
		quantComment: 'small, greater quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q4_K_M.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q4_K_M',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q4_K_M',
		size: 4.08,
		maxRam: 6.58,
		quantComment: 'medium, balanced quality - recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q5_K_S.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q5_K_S',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q5_K_S',
		size: 4.65,
		maxRam: 7.15,
		quantComment: 'large, low quality loss - recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q5_K_M.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q5_K_M',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q5_K_M',
		size: 4.78,
		maxRam: 7.28,
		quantComment: 'large, very low quality loss - recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q6_K.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q6_K',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q6_K',
		size: 5.53,
		maxRam: 8.03,
		quantComment: 'very large, extremely low quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-7B-GGUF/resolve/main/codellama-7b.Q8_0.gguf',
		name: 'CodeLlama',
		parameters: '7B',
		flavour: 'Code',
		version: '1',
		quality: 'Q8_0',
		extension: 'GGUF',
		fullName: 'CodeLlama 7B Q8_0',
		size: 7.16,
		maxRam: 9.66,
		quantComment: 'very large, extremely low quality loss - not recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	// CodeLLama 13B
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q2_K.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q2_K',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q2_K',
		size: 5.43,
		maxRam: 7.93,
		quantComment:
			'smallest, significant quality loss - not recommended for most purposes',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q3_K_S.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q3_K_S',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q3_K_S',
		size: 5.66,
		maxRam: 8.16,
		quantComment: 'very small, high quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q3_K_M.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q3_K_M',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q3_K_M',
		size: 6.34,
		maxRam: 8.84,
		quantComment: 'very small, high quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q3_K_L.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q3_K_L',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q3_K_L',
		size: 6.93,
		maxRam: 9.43,
		quantComment: 'small, substantial quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q4_K_S.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q4_K_S',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q4_K_S',
		size: 7.41,
		maxRam: 9.91,
		quantComment: 'small, greater quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q4_K_M.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q4_K_M',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q4_K_M',
		size: 7.87,
		maxRam: 10.37,
		quantComment: 'medium, balanced quality - recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q5_K_S.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q5_K_S',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q5_K_S',
		size: 8.97,
		maxRam: 11.47,
		quantComment: 'large, low quality loss - recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q5_K_M.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q5_K_M',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q5_K_M',
		size: 9.23,
		maxRam: 11.73,
		quantComment: 'large, very low quality loss - recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q6_K.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q6_K',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q6_K',
		size: 10.68,
		maxRam: 13.18,
		quantComment: 'very large, extremely low quality loss',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	{
		id: 'https://huggingface.co/TheBloke/CodeLlama-13B-GGUF/resolve/main/codellama-13b.Q8_0.gguf',
		name: 'CodeLlama',
		parameters: '13B',
		flavour: 'Code',
		version: '1',
		quality: 'Q8_0',
		extension: 'GGUF',
		fullName: 'CodeLlama 13B Q8_0',
		size: 13.83,
		maxRam: 16.33,
		quantComment: 'very large, extremely low quality loss - not recommended',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
	//
	// Llama 2 7B Chat Uncensored
	//
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q2_K.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q2_K',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q2_K',
		size: 2.83,
		maxRam: 5.33,
		quantComment:
			'smallest, significant quality loss - not recommended for most purposes',
		description:
			'A version of LLaMA2 model tailored for uncensored chat applications, optimized for smaller size and RAM usage with significant quality loss, making it less suitable for most purposes.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q3_K_S.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q3_K_S',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q3_K_S',
		size: 2.95,
		maxRam: 5.45,
		quantComment: 'very small, high quality loss',
		description:
			'A specialized version of the LLaMA2 model for chat applications with a focus on reduced size and memory requirements, featuring a high quality loss.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q3_K_M.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q3_K_M',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q3_K_M',
		size: 3.3,
		maxRam: 5.8,
		quantComment: 'very small, high quality loss',
		description:
			'This iteration of the LLaMA2 model is optimized for chat purposes, balancing size and efficiency with a slight compromise on quality.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q3_K_L.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q3_K_L',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q3_K_L',
		size: 3.6,
		maxRam: 6.1,
		quantComment: 'small, substantial quality loss',
		description:
			'Designed for chat applications, this LLaMA2 model version offers a compact size with manageable memory requirements at the cost of some quality loss.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q4_K_S.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q4_K_S',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q4_K_S',
		size: 3.86,
		maxRam: 6.36,
		quantComment: 'small, greater quality loss',
		description:
			'A compact and efficient version of the LLaMA2 model for uncensored chat, optimized to maintain a balance between size, memory usage, and quality.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q4_K_M.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q4_K_M',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q4_K_M',
		size: 4.08,
		maxRam: 6.58,
		quantComment: 'medium, balanced quality - recommended',
		description:
			'The LLaMA2 7B Chat Uncensored Q4_K_M model offers a balanced compromise between file size, RAM requirements, and quality, making it a recommended choice for chat applications.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q5_K_S.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q5_K_S',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q5_K_S',
		size: 4.65,
		maxRam: 7.15,
		quantComment: 'large, low quality loss - recommended',
		description:
			'A large-scale model version of LLaMA2 for chat, the Q5_K_S variant minimizes quality loss while requiring more memory, recommended for its efficient performance.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q5_K_M.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q5_K_M',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q5_K_M',
		size: 4.78,
		maxRam: 7.28,
		quantComment: 'large, very low quality loss - recommended',
		description:
			'LLaMA2 7B Chat Uncensored Q5_K_M is tailored for high-demand chat applications, offering substantial capacity with minimal compromise on quality.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q6_K.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q6_K',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q6_K',
		size: 5.53,
		maxRam: 8.03,
		quantComment: 'very large, extremely low quality loss',
		description:
			'Optimized for expansive chat integrations, the Q6_K version of LLaMA2 ensures extensive capacity with remarkably low quality loss, suitable for advanced applications.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	{
		id: 'https://huggingface.co/TheBloke/llama2_7b_chat_uncensored-GGUF/resolve/main/llama2_7b_chat_uncensored.Q8_0.gguf',
		name: 'LLaMA2',
		parameters: '7B',
		flavour: 'Chat',
		uncensored: true,
		version: '2',
		quality: 'Q8_0',
		extension: 'GGUF',
		fullName: 'LLaMA2 7B Chat Uncensored Q8_0',
		size: 7.16,
		maxRam: 9.66,
		quantComment: 'very large, extremely low quality loss - not recommended',
		description:
			'The LLaMA2 7B Chat Uncensored Q8_0 variant represents the upper echelon in terms of size and memory requirements.',
		promptTemplate: llamaChatUncensoredPrompt,
	},
	// Llama 3 8B
	{
		id: 'https://huggingface.co/QuantFactory/Meta-Llama-3-8B-Instruct-GGUF/resolve/main/Meta-Llama-3-8B-Instruct.Q3_K_M.gguf',
		name: 'Llama 3',
		parameters: '8B',
		flavour: 'Code',
		version: '3',
		quality: 'Q3_K_M',
		extension: 'GGUF',
		fullName: 'Llama 3 8B Q3_K_M',
		size: 4.02,
		maxRam: 5.33,
		quantComment:
			'smallest, significant quality loss - not recommended for most purposes',
		description: codellamaDescription,
		promptTemplate: '{prompt}',
	},
];

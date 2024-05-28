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
import * as crypto from 'crypto-js';
import { firstValueFrom, map, throwError } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';
import { catchError, switchMap } from 'rxjs/operators';
import { LapiService } from './lapi.service';
import { Config } from 'shared-lib/models/types';
import { ReplaySubject, Observable } from 'rxjs';

import {
	CompletionResponse,
	promptEndpoint,
	PromptRequest,
} from '../../../shared/backend-api/llm';

export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
	encrypt: boolean;
	encryptKey: string;
	backendAddress: string;
	localPromptAddress: string;
	localtronAddress: string;
	stripeCustomerPortalURL: string;
	stripeSFBSoloSubscribeLink: string;
	stripeSFBBusinessSubscribeLink: string;
}

export interface LocaltronServiceConfig {
	env: Environment;
}
export const LOCALTRON_SERVICE_CONFIG =
	new InjectionToken<LocaltronServiceConfig>('LocaltronServiceConfig');

@Injectable({
	providedIn: 'root',
})
export class LocaltronService {
	public activeThreadId: string;
	public llmAddressOverride: string;

	private headers: HttpHeaders;
	private config: LocaltronServiceConfig;

	constructor(
		private http: HttpClient,
		private cs: CookieService,
		private lapi: LapiService,
		@Inject(LOCALTRON_SERVICE_CONFIG) config: LocaltronServiceConfig
	) {
		this.config = config;
		this.headers = new HttpHeaders();

		this.start();
	}

	onPromptListUpdateSubject = new ReplaySubject<Prompt[]>(1);
	onPromptListUpdate$ = this.onPromptListUpdateSubject.asObservable();

	start() {
		this.lapi.onLLMAddressChange$.subscribe((v) => {
			this.llmAddressOverride = v;
		});

		const pollModelStatus = async () => {
			try {
				let rsp = await this.modelStatus();
				if (rsp?.status?.modelAddress) {
					let llmAddr = rsp.status?.modelAddress;

					if (!llmAddr?.startsWith('http://')) {
						llmAddr = 'http://' + llmAddr;
					}

					this.llmAddressOverride = llmAddr;
					this.lapi.onLLMAddressChangeSubject.next(llmAddr);
				}
				if (rsp?.status?.selectedExists) {
					await this.call('/model/start', {}).catch((e) => {
						console.error('Error starting model', {
							error: e,
						});
					});
				}
				if (rsp?.status?.running) {
					this.lapi.onModelLaunchSubject.next({});
				}
				this.lapi.onModelCheckSubject.next({
					selectedExists: rsp?.status?.selectedExists,
				});
			} catch (error) {
				console.error('Error in pollModelStatus', error);
			}
			setTimeout(pollModelStatus, 1000); // Call again after 1 second
		};

		const pollDockerInfo = async () => {
			try {
				let rsp = await this.dockerInfo();

				this.lapi.onDockerInfoSubject.next({
					hasDocker: rsp?.info?.hasDocker,
				});
			} catch (error) {
				console.error('Error in pollDockerInfo', error);
			}
			setTimeout(pollDockerInfo, 1000); // Call again after 1 second
		};

		const pollPromptList = async () => {
			try {
				let rsp = await this.promptList();

				this.onPromptListUpdateSubject.next(rsp.prompts);
			} catch (error) {
				console.error('Error in pollPromptList', error);
			}
			setTimeout(pollPromptList, 1000); // Call again after 1 second
		};

		const pollConfig = async () => {
			try {
				let rsp = await this.call('/config/get', {});
				this.lapi.lastConfig = rsp?.config;
				this.lapi.onConfigUpdateSubject.next(rsp?.config as Config);
			} catch (error) {
				console.error('Error in pollConfig', error);
			}
			setTimeout(pollConfig, 1000); // Call again after 1 second
		};

		const pollFileDownloadStatus = async () => {
			try {
				let rsp = await this.call('/download/list', {});
				this.lapi.onFileDownloadStatusSubject.next({
					allDownloads: rsp?.downloads as DownloadDetails[],
				});
			} catch (error) {
				console.error('Error in pollFileDownloadStatus', error);
			}
			setTimeout(pollFileDownloadStatus, 1000); // Call again after 1 second
		};

		// Start the polling loops
		pollModelStatus();
		pollDockerInfo();
		pollPromptList();
		pollConfig();
		pollFileDownloadStatus();
	}

	private c(): string {
		return this.config.env.encryptKey;
	}

	call(path: string, request: any): Promise<any> {
		if (!this.config.env.localtronAddress) {
			console.log('Localtron address is not set', {
				config: this.config,
			});
			throw 'Localtron address seems to be empty';
		}

		let uri = this.config.env.localtronAddress + path;
		// console.log("calling", uri);

		// Encrypt the request if env.encrypt is true
		let body = // this.config.env.encrypt
			//? this.encrypt(request)
			JSON.stringify(request);

		let headers = this.headers.set(
			'Authorization',
			'Bearer ' + this.cs.get('the_token')
		);
		if (this.config.env.encrypt) {
			// headers = headers.set('X-Content-Encrypted', 'true');
		}

		return firstValueFrom(
			this.http
				.post<any>(uri, body, {
					headers: headers,
					responseType: 'text' as 'json',
				})
				.pipe(
					map((response) => {
						// this.config.env.encrypt
						// ? this.decrypt(response)
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

	async appLogDisable(): Promise<void> {
		return this.call('/app/log/disable', {});
	}

	async appLogEnable(): Promise<void> {
		return this.call('/app/log/enable', {});
	}

	async appLogStatus(): Promise<LoggingStatus> {
		return this.call('/app/log/status', {});
	}

	async dockerInfo(): Promise<DockerInfoResponse> {
		return this.call('/docker/info', {});
	}

	async modelStatus(): Promise<ModelStatusResponse> {
		return this.call('/model/status', {});
	}

	async downloadDo(url: string) {
		this.call('/download/do', { url: url });
	}

	async modelStart(url?: string) {
		this.call('/model/start', { url: url });
	}

	async downloadPause(url: string) {
		this.call('/download/pause', { url: url });
	}

	async downloadList(): Promise<DownloadsResponse> {
		return this.call('/download/list', {});
	}

	async chatMessageAdd(message: ChatMessage): Promise<void> {
		let req: AddChatMessageRequest = { message: message };
		return this.call('/chat/message/add', req);
	}

	async chatMessageDelete(messageId: string): Promise<GetChatThreadResponse> {
		let req: DeleteChatMessageRequest = { messageId: messageId };
		return this.call('/chat/message/delete', req);
	}

	async chatMessages(threadId: string): Promise<GetChatMessagesResponse> {
		let req: GetChatMessagesRequest = { threadId: threadId };
		return this.call('/chat/messages', req);
	}

	async chatThread(threadId: string): Promise<GetChatThreadResponse> {
		let req: GetChatThreadRequest = { threadId: threadId };
		return this.call('/chat/thread', req);
	}

	async chatThreadAdd(thread: ChatThread): Promise<AddChatThreadResponse> {
		let req: AddChatThreadRequest = { thread: thread };
		return this.call('/chat/thread/add', req);
	}

	async chatThreadUpdate(
		thread: ChatThread
	): Promise<UpdateChatThreadResponse> {
		let req: UpdateChatThreadRequest = { thread: thread };
		return this.call('/chat/thread/update', req);
	}

	async chatThreadDelete(threadId: string): Promise<void> {
		let req: DeleteChatThreadRequest = { threadId: threadId };
		return this.call('/chat/thread/delete', req);
	}

	async chatThreads(): Promise<GetChatThreadsResponse> {
		let req: GetChatThreadsRequest = {};
		return this.call('/chat/threads', req);
	}

	setActiveThreadId(id: string) {
		localStorage.setItem(this.activeThreadId, id);
	}

	getActiveThreadId(): string {
		const activeThreadId = localStorage.getItem(this.activeThreadId);
		if (!activeThreadId) {
			return null;
		}
		return activeThreadId;
	}

	uuid() {
		function generateSegment(length) {
			return Array.from({ length: length }, () =>
				Math.floor(Math.random() * 16).toString(16)
			).join('');
		}

		return (
			generateSegment(8) +
			'-' +
			generateSegment(4) +
			'-' +
			generateSegment(4) +
			'-' +
			generateSegment(4) +
			'-' +
			generateSegment(12)
		);
	}

	async promptAdd(prompt: Prompt): Promise<void> {
		if (!prompt.id) {
			prompt.id = this.uuid();
		}
		let req: AddPromptRequest = { prompt: prompt };
		return this.call('/prompt/add', req);
	}

	async promptList(): Promise<ListPromptsResponse> {
		return this.call('/prompt/list', {});
	}

	promptSubscribe(threadId: string): Observable<CompletionResponse> {
		let uri =
			this.config.env.localtronAddress +
			'/prompt/subscribe?threadId=' +
			threadId;

		const token = this.cs.get('the_token');
		const headers = {
			Authorization: 'Bearer ' + token,
			'Content-Type': 'application/json',
		};

		return new Observable((observer) => {
			const controller = new AbortController();
			const { signal } = controller;

			fetch(uri, {
				method: 'GET',
				headers: headers,
				signal: signal,
			})
				.then((response) => {
					if (!response.ok) {
						observer.error(`HTTP error! status: ${response.status}`);
						return;
					}
					const reader = response.body.getReader();
					return new ReadableStream({
						start(controller) {
							function push() {
								reader
									.read()
									.then(({ done, value }) => {
										if (done) {
											console.debug('Prompt stream completed');
											controller.close();
											observer.complete();
											return;
										}
										// Convert the Uint8Array to string
										const text = new TextDecoder().decode(value);
										let lines = text.split('\n');
										lines.forEach((line) => {
											const trimmedLine = line.trim();

											if (
												trimmedLine === '' ||
												trimmedLine === 'data: ' ||
												trimmedLine === 'data: [DONE]'
											) {
												// Skip empty lines, lines containing only 'data: ', or "[DONE]" markers
												return;
											}

											const cleanedText = trimmedLine
												.replace(/^data: /gm, '')
												.trim();

											try {
												const json = JSON.parse(cleanedText);
												observer.next(json);
											} catch (error) {
												console.error(
													'Error parsing prompt response chunk JSON',
													{
														error: error,
														promptResponseChunk: cleanedText,
													}
												);
												// Decide how you want to handle parsing errors.
												// For continuous streaming, you might not want to call observer.error() here
												// unless it's a critical error that requires stopping the stream.
											}
										});

										// Call push again outside the loop to continue reading
										push();
									})
									.catch((err) => {
										if (
											err instanceof Error &&
											err.message.includes('BodyStreamBuffer was aborted')
										) {
											// we ignore this because this is normal
										} else {
											console.error('Error reading from stream', {
												error: JSON.stringify(err),
											});

											observer.error(err);
											controller.error(err);
										}
										observer.error(err);
										controller.error(err);
									});
							}
							push();
						},
					});
				})
				.catch((err) => {
					observer.error(err);
				});

			return () => {
				controller.abort(); // This ensures fetch is aborted when unsubscribing
			};
		});
	}

	prompt(request: PromptRequest): Observable<CompletionResponse> {
		let uri;
		if (this.llmAddressOverride) {
			uri = this.llmAddressOverride + '/v1/completions';
		} else if (this.config.env.localPromptAddress) {
			console.debug('Using local prompt', {
				address: this.config.env.localPromptAddress,
			});
			uri = this.config.env.localPromptAddress + '/v1/completions';
		} else {
			uri = this.config.env.backendAddress + promptEndpoint;
		}

		const token = this.cs.get('the_token');

		// Prepare headers
		const headers = {
			Authorization: 'Bearer ' + token,
			'Content-Type': 'application/json',
		};

		if (!request.max_tokens) {
			request.max_tokens = 2048;
		}

		console.debug('Prompt sync started');
		return new Observable((observer) => {
			fetch(uri, {
				method: 'POST',
				headers: headers,
				body: JSON.stringify(request),
			})
				.then((response) => {
					if (!response.ok) {
						throw new Error(`HTTP error! status: ${response.status}`);
					}
					return response.json();
				})
				.then((data) => {
					console.debug('Prompt sync completed');
					observer.next(data);
					observer.complete();
				})
				.catch((err) => {
					observer.error(err);
				});
		});
	}
}

interface ModelStatus {
	selectedExists: boolean;
	currentModelId: string;
	/** Running triggers onModelLaunch on the frontend.
	 * Running is true when the model is both running and answering
	 * - fully loaded.
	 */
	running: boolean;
	modelAddress: string;
}

// {
//   "status": {
//     "selectedExists": false,
//     "currentModelId": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf"
//   }
// }
interface ModelStatusResponse {
	status: ModelStatus | null;
}

interface DockerInfo {
	hasDocker: boolean;
	dockerDaemonAddress?: string;
	error?: string;
}

// {
//   "info": {
//     "hasDocker": true
//   }
// }
interface DockerInfoResponse {
	info: DockerInfo;
}

interface LoggingStatus {
	enabled: boolean;
}

export interface DownloadDetails {
	id: string;
	url: string;
	fileName: string;
	dir?: string;
	progress?: number;
	downloadedBytes: number;
	fullFileSize?: number;
	status: 'inProgress' | 'completed' | 'paused' | 'cancelled' | 'failed';
	filePath?: string;
	paused?: boolean;
	cancelled?: boolean;
	error?: string;
}

type DownloadsResponse = {
	downloads: DownloadDetails[];
};

export interface ChatThread {
	id?: string;
	topicId?: string;
	name?: string;
	time?: string;
}

export interface ChatMessage {
	id?: string;
	threadId: string;
	messageContent: string;
	isUserMessage: boolean;
	time?: string;
}

export interface ChatFile {
	threads: ChatThread[];
	messages: ChatMessage[];
}

export interface AddChatMessageRequest {
	message: ChatMessage;
}

export interface AddChatThreadRequest {
	thread: ChatThread;
}

export interface AddChatThreadResponse {
	thread: ChatThread;
}

export interface UpdateChatThreadRequest {
	thread: ChatThread;
}

export interface UpdateChatThreadResponse {
	thread: ChatThread;
}

export interface DeleteChatThreadRequest {
	threadId: string;
}

export interface DeleteChatMessageRequest {
	messageId: string; // Corrected field name from "threadId" to "messageId"
}

export interface GetChatThreadRequest {
	threadId: string;
}

type GetChatThreadResponse = {
	thread: ChatThread;
};

type GetChatThreadsRequest = {};

type GetChatThreadsResponse = {
	threads: ChatThread[];
};

type GetChatMessagesRequest = {};

type GetChatMessagesResponse = {
	messages: ChatMessage[];
};

// also duplicated in api service
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
}

export interface Prompt {
	id?: string;
	threadId: string;
	prompt: string;
	modelId: string;
}

export interface AddPromptRequest {
	prompt: Prompt;
}

export interface ListPromptsRequest {}

export interface ListPromptsResponse {
	prompts: Prompt[];
}

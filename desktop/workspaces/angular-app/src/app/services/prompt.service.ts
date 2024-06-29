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
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { ReplaySubject, Observable, catchError } from 'rxjs';
import { FirehoseService } from './firehose.service';
import { first } from 'rxjs';
import { UserService } from './user.service';

@Injectable({
	providedIn: 'root',
})
export class PromptService {
	onPromptListUpdateSubject = new ReplaySubject<Prompt[]>(1);
	onPromptListUpdate$ = this.onPromptListUpdateSubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'promptRemoved':
				case 'promptProcessingStarted':
				case 'promptProcessingFinished':
				case 'promptAdded': {
					const rsp = await this.promptList();
					this.onPromptListUpdateSubject.next(rsp.prompts);
					break;
				}
			}
			return;
		});

		try {
			const rsp = await this.promptList();

			this.onPromptListUpdateSubject.next(rsp.prompts);
		} catch (error) {
			console.error('Error in pollPromptList', {
				error: JSON.stringify(error),
			});
		}
	}

	async promptAdd(prompt: Prompt): Promise<void> {
		if (!prompt.id) {
			prompt.id = this.localtron.uuid();
		}
		const request: AddPromptRequest = { prompt: prompt };
		return this.localtron.call('/prompt/add', request);
	}

	async promptRemove(prompt: Prompt): Promise<void> {
		const request: RemovePromptRequest = { prompt: prompt };
		return this.localtron.call('/prompt/remove', request);
	}

	async promptList(): Promise<ListPromptsResponse> {
		return this.localtron.call('/prompt/list', {});
	}

	private resubCount = 0;
	promptSubscribe(threadId: string): Observable<CompletionResponse> {
		if (!threadId) {
			console.log('No thread id');
			throw 'no thread id';
		}

		return new Observable<CompletionResponse>((observer) => {
			const controller = new AbortController();
			const { signal } = controller;

			const subscribe = () => {
				console.log('Subscribing to thread', {
					threadId: threadId,
				});

				const uri =
					this.localtron.config.env.localtronAddress +
					'/prompt/subscribe?threadId=' +
					threadId;

				const headers = {
					Authorization: 'Bearer ' + this.userService.getToken(),
					'Content-Type': 'application/json',
				};

				fetch(uri, {
					method: 'GET',
					headers: headers,
					signal: signal,
				})
					.then((response) => {
						if (!response || !response.body) {
							observer.error(`Response is empty`);
							return;
						}
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
											const lines = text.split('\n');
											for (const line of lines) {
												const trimmedLine = line.trim();

												if (
													trimmedLine === '' ||
													trimmedLine === 'data: ' ||
													trimmedLine === 'data: [DONE]'
												) {
													// Skip empty lines, lines containing only 'data: ', or "[DONE]" markers
													continue
												}

												const cleanedText = trimmedLine
													.replaceAll(/^data: /gm, '')
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
													// Decide how we want to handle parsing errors.
													// For continuous streaming, we might not want to call observer.error() here
													// unless it's a critical error that requires stopping the stream.
												}
											}

											// Call push again outside the loop to continue reading
											push();
										})
										.catch((error) => {
											if (
												error instanceof Error &&
												error.message.includes('BodyStreamBuffer was aborted')
											) {
												// we ignore this because this is normal
											} else {
												console.error('Error reading from stream', {
													error: JSON.stringify(error),
												});

												observer.error(error);
												controller.error(error);
											}
											observer.error(error);
											controller.error(error);
										});
								}
								push();
							},
						});
					})
					.catch((error) => {
						observer.error(error);
					});
			};

			sleep(this.resubCount * 20).then(() => {
				subscribe();
			});

			return () => {
				controller.abort(); // This ensures fetch is aborted when unsubscribing
			};
		}).pipe(
			catchError((error) => {
				console.error('Prompt subscription error', {
					error: JSON.stringify(error),
				});
				this.resubCount++;
				return this.promptSubscribe(threadId);
			})
		);
	}
}

function sleep(ms: number) {
	return new Promise((resolve) => setTimeout(resolve, ms));
}

export type PromptStatus =
	| 'scheduled'
	| 'running'
	| 'completed'
	| 'errored'
	| 'abandoned'
	| 'canceled';

export interface Prompt {
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
	characterId?: string;
	modelId?: string;
	createdAt?: string;
	status?: PromptStatus;
	lastRun?: string;
	runCount?: number; // How many times this was ran (retries are due to errors)
	error?: string;
	maxRetries?: number;
}

export interface AddPromptRequest {
	prompt: Prompt;
}

export interface RemovePromptRequest {
	prompt: Prompt;
}

// eslint-disable-next-line
export interface ListPromptsRequest {}

export interface ListPromptsResponse {
	prompts: Prompt[];
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

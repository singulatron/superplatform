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
import { ReplaySubject, Observable } from 'rxjs';
import { FirehoseService } from './firehose.service';

@Injectable({
	providedIn: 'root',
})
export class PromptService {
	onPromptListUpdateSubject = new ReplaySubject<Prompt[]>(1);
	onPromptListUpdate$ = this.onPromptListUpdateSubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private firehoseService: FirehoseService
	) {
		this.init();
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'promptAdded':
					let rsp1 = await this.promptList();
					this.onPromptListUpdateSubject.next(rsp1.prompts);
					break;
				case 'promptProcessingStarted':
					let rsp2 = await this.promptList();
					this.onPromptListUpdateSubject.next(rsp2.prompts);
					break;
				case 'promptProcessingFinished':
					let rsp3 = await this.promptList();
					this.onPromptListUpdateSubject.next(rsp3.prompts);
					break;
			}
			return null;
		});

		try {
			let rsp = await this.promptList();

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
		let req: AddPromptRequest = { prompt: prompt };
		return this.localtron.call('/prompt/add', req);
	}

	async promptList(): Promise<ListPromptsResponse> {
		return this.localtron.call('/prompt/list', {});
	}

	promptSubscribe(threadId: string): Observable<CompletionResponse> {
		let uri =
			this.localtron.config.env.localtronAddress +
			'/prompt/subscribe?threadId=' +
			threadId;

		const token = ''; // this.cs.get('the_token');
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
	/*
		Prompt without template, eg. `What is a banana`?
	*/
	prompt: string;
	/*
		Prompt template (eg. `[INST]{prompt}[/INST]`)
		Optional. Might be derived from/through the modelId
	*/
	template: string;
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

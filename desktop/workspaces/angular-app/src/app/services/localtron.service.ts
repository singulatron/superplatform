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
import { firstValueFrom, map, throwError } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';
import { catchError } from 'rxjs/operators';
import { ElectronAppService } from './electron-app.service';
import { Config } from 'shared-lib/models/types';
import { Observable } from 'rxjs';

export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
	backendAddress: string;
	localPromptAddress: string;
	localtronAddress: string;
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
	public llmAddressOverride: string = '';

	private headers: HttpHeaders;
	public config: LocaltronServiceConfig;

	constructor(
		private http: HttpClient,
		private cs: CookieService,
		private eapp: ElectronAppService,
		@Inject(LOCALTRON_SERVICE_CONFIG) config: LocaltronServiceConfig
	) {
		this.config = config;
		this.headers = new HttpHeaders();

		this.start();
	}

	start() {
		this.firehoseSubscribe().subscribe((event) => {
			switch (event.name) {
				case 'chatMessageAdded':
			}
		});

		this.eapp.onLLMAddressChange$.subscribe((v) => {
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
					this.eapp.onLLMAddressChangeSubject.next(llmAddr);
				}
				if (rsp?.status?.selectedExists) {
					await this.call('/model/start', {}).catch((e) => {
						console.error('Error starting model', {
							error: e,
						});
					});
				}
				if (rsp?.status?.running) {
					this.eapp.onModelLaunchSubject.next({});
				}
				this.eapp.onModelCheckSubject.next({
					selectedExists: rsp?.status?.selectedExists,
				});
			} catch (error) {
				console.error('Error in pollModelStatus', {
					error: JSON.stringify(error),
				});
			}
			setTimeout(pollModelStatus, 1000);
		};

		const pollDockerInfo = async () => {
			try {
				let rsp = await this.dockerInfo();

				this.eapp.onDockerInfoSubject.next({
					hasDocker: rsp?.info?.hasDocker,
				});
			} catch (error) {
				console.error('Error in pollDockerInfo', {
					error: JSON.stringify(error),
				});
			}
			setTimeout(pollDockerInfo, 1000);
		};

		const pollPromptList = async () => {
			try {
				let rsp = await this.promptList();

				this.onPromptListUpdateSubject.next(rsp.prompts);
			} catch (error) {
				console.error('Error in pollPromptList', {
					error: JSON.stringify(error),
				});
			}
			setTimeout(pollPromptList, 1000);
		};

		const pollConfig = async () => {
			try {
				let rsp = await this.call('/config/get', {});
				this.eapp.lastConfig = rsp?.config;
				this.eapp.onConfigUpdateSubject.next(rsp?.config as Config);
			} catch (error) {
				console.error('Error in pollConfig', {
					error: JSON.stringify(error),
				});
			}
			setTimeout(pollConfig, 1000);
		};

		const pollFileDownloadStatus = async () => {
			try {
				let rsp = await this.call('/download/list', {});
				this.eapp.onFileDownloadStatusSubject.next({
					allDownloads: rsp?.downloads as DownloadDetails[],
				});
			} catch (error) {
				console.error('Error in pollFileDownloadStatus', {
					error: JSON.stringify(error),
				});
			}
			setTimeout(pollFileDownloadStatus, 1000);
		};

		pollModelStatus();
		pollDockerInfo();
		pollPromptList();
		pollConfig();
		pollFileDownloadStatus();
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

	uuid() {
		function generateSegment(length: number) {
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

	// @todo a lot of this is duplication from promptSubscribe
	firehoseSubscribe(): Observable<FirehoseEvent> {
		console.info('Subscribing to the firehose');

		let uri = this.config.env.localtronAddress + '/firehose/subscribe';

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

export interface FirehoseEvent {
	name: string;
	data: any;
}

export interface ChatMessageAddedEvent {
	threadId: string;
}

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
import { Observable, Subject } from 'rxjs';
import { UserService } from './user.service';
import { first } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class FirehoseService {
	private firehoseEventSubject = new Subject<FirehoseEvent>();
	public firehoseEvent$ = this.firehoseEventSubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private userService: UserService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		this.firehoseSubscribe().subscribe((event) => {
			this.firehoseEventSubject.next(event);
		});
	}

	// @todo a lot of this is duplication from promptSubscribe
	private firehoseSubscribe(): Observable<FirehoseEvent> {
		console.info('Subscribing to the firehose');

		let uri =
			this.localtron.config.env.localtronAddress + '/firehose/subscribe';

		const headers = {
			Authorization: 'Bearer ' + this.userService.getToken(),
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

export interface FirehoseEvent {
	name: string;
	data: any;
}

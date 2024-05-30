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
import {
	Component,
	ViewEncapsulation,
	Input,
	Output,
	EventEmitter,
	OnChanges,
	SimpleChanges,
} from '@angular/core';
import { Subscription } from 'rxjs';

import { ChangeDetectorRef } from '@angular/core';
import { finalize } from 'rxjs';
import { ApiService } from '../../stdlib/api.service';

import {
	ChatThread,
	ChatMessage,
	LocaltronService,
	Model,
} from '../../../src/app/services/localtron.service';

import { LapiService } from '../../../src/app/services/lapi.service';

const defaultThreadName = 'New chat';

@Component({
	selector: 'app-chat-box',
	templateUrl: './chat-box.component.html',
	styleUrl: './chat-box.component.css',
	encapsulation: ViewEncapsulation.None,
})
export class ChatBoxComponent implements OnChanges {
	@Input() promptTemplate: string = '[INST] {prompt} [/INST]';
	@Input() userMessageTemplate: string = 'User: {message}\n';
	@Input() modelMessageTemplate: string = 'Model: {message}\n';
	@Input() latestMessageTemplate: string = "User's latest message: {message}\n";

	@Input() historyEnabled = false;
	@Input() advancedHistoryEnabled = false;
	@Input() contextTemplate: string =
		`These are the previous messages from the User:\n{message} Answer only to the last message.`;
	@Input() advancedContextTemplate: string =
		`This is your previous conversation with the User (you are the "Model"):\n{message}`;

	@Input() threadNameSummaryTemplate =
		'Summarize, shorten this question in 3-5 words, keep it as a question: {message}';

	@Input() thread: ChatThread;

	@Output() onThreadUpdate = new EventEmitter<ChatThread>();
	@Output() onFirstMessageSend = new EventEmitter<ChatThread>();
	@Output() onCopyToClipboard = new EventEmitter<string>();

	private model: Model | undefined;
	private models: Model[] = [];

	public message: string = '';
	public messages: ChatMessage[] = [];
	public messageCurrentlyStreamed = '';

	constructor(
		private api: ApiService,
		private localtron: LocaltronService,
		public lapi: LapiService,
		private cd: ChangeDetectorRef
	) {}

	private subscriptions: Subscription[] = [];

	async ngOnInit() {
		this.models = await this.api.getModels();
		this.subscriptions.push(
			this.lapi.onConfigUpdate$.subscribe(async (config) => {
				this.model = this.models?.find(
					(m) => m.id == config?.model?.currentModelId
				);
			})
		);
	}

	streamSubscription: Subscription;

	ngOnDestroy() {
		this.streamSubscription.unsubscribe();
		this.subscriptions.forEach((s) => {
			s.unsubscribe();
		});
	}

	async ngOnChanges(changes: SimpleChanges): Promise<void> {
		if (changes.thread) {
			if (this.streamSubscription) {
				this.streamSubscription.unsubscribe();
			}

			let threadId;

			if (!this.thread) {
				this.thread = {
					id: this.localtron.uuid(),
				};
				threadId = this.thread.id;
			} else {
				threadId = changes.thread.currentValue.id;
				let rsp = await this.localtron.chatMessages(threadId);
				this.messages = rsp.messages;
			}

			this.messageCurrentlyStreamed = '';
			let first = true;

			this.streamSubscription = this.localtron
				.promptSubscribe(threadId)
				.subscribe(async (response) => {
					if (
						response?.choices &&
						response?.choices?.length > 0 &&
						response?.choices[0]?.text
					) {
						let insidePre =
							(this.messageCurrentlyStreamed.match(/```/g) || []).length % 2 ===
							1;
						let addVal = insidePre
							? response?.choices[0].text
							: escapeHtml(response?.choices[0].text);

						if (first) {
							addVal = addVal.trimLeft();
							first = false;
						}
						this.messageCurrentlyStreamed += addVal;
					}
					// finish_reason: "stop" might be model specific
					if (
						response?.choices?.length > 0 &&
						response?.choices[0]?.finish_reason === 'stop'
					) {
						this.onThreadUpdate.emit(changes.thread.currentValue);

						if (this.messages?.length == 1) {
							this.setThreadName(this.messages[0].messageContent);
						}
						let rsp = await this.localtron.chatMessages(threadId);
						this.messages = rsp.messages;
						this.messageCurrentlyStreamed = '';
					}
					this.cd.detectChanges();
				});
		}
	}

	async send() {
		if (this.messages?.length == 0) {
			this.thread.name = this.message.slice(0, 100);
			this.onThreadUpdate.emit(this.thread);
		}

		let msg = this.message;
		this.message = '';

		await this.localtron.chatMessageAdd({
			threadId: this.thread.id as string,
			messageContent: msg,
			isUserMessage: true,
		});

		this.prompt(msg);
	}

	// Handle keydown event to differentiate between Enter and Shift+Enter
	handleKeydown(event: KeyboardEvent): void {
		if (event.key === 'Enter' && !event.shiftKey) {
			event.preventDefault();
			if (this.hasNonWhiteSpace(this.message)) {
				this.send();
			}
		} else if (event.key === 'Enter' && event.shiftKey) {
			event.preventDefault();
			this.message += '\n';
		}
	}

	public hasNonWhiteSpace(value: string): boolean {
		if (!value) {
			return false;
		}
		return /\S/.test(value);
	}

	async prompt(msg: string): Promise<void> {
		let userMessages = this.messages?.filter((m) => m.isUserMessage) || [];
		let modelMessages = this.messages?.filter((m) => !m.isUserMessage) || [];
		let exchange = zigzagArrays(
			userMessages.map((m) =>
				this.userMessageTemplate.replace('{message}', m?.messageContent)
			),
			this.advancedHistoryEnabled
				? modelMessages.map((m) =>
						this.modelMessageTemplate.replace('{message}', m?.messageContent)
					)
				: []
		).join('');

		let fullContext =
			(this.historyEnabled || this.advancedHistoryEnabled) &&
			userMessages?.length > 0
				? this.advancedHistoryEnabled
					? this.advancedContextTemplate.replace('{message}', exchange)
					: this.contextTemplate.replace('{message}', exchange)
				: '';

		fullContext +=
			(this.historyEnabled || this.advancedHistoryEnabled) &&
			userMessages?.length > 0
				? this.latestMessageTemplate.replace('{message}', msg)
				: msg;
		let fullPrompt = this.promptTemplate
			? this.promptTemplate.replace('{prompt}', fullContext)
			: fullContext;

		this.messages.push({
			threadId: this.thread.id as string,
			id: this.localtron.uuid(),
			messageContent: msg,
			isUserMessage: true,
		});

		await this.localtron.promptAdd({
			prompt: fullPrompt,
			threadId: this.thread.id as string,
			modelId: this.model?.id as string,
		});
	}

	setThreadName(msg: string) {
		if (this.thread?.name !== defaultThreadName) {
			return;
		}

		let prompt = this.promptTemplate
			? this.promptTemplate.replace(
					'{prompt}',
					this.threadNameSummaryTemplate.replace('{message}', msg)
				)
			: msg;
		let newThreadName = '';
		this.localtron
			.prompt({
				prompt: prompt,
				stream: true,
			})
			.pipe(finalize(() => {}))
			.subscribe((response) => {
				if (response?.choices?.length > 0 && response?.choices[0]?.text) {
					newThreadName += response?.choices[0].text;
					this.thread.name = newThreadName;
					this.localtron.chatThreadUpdate(this.thread);

					this.cd.detectChanges();
				}
			});
	}

	propagateCopyToClipboard(text: string) {
		this.onCopyToClipboard.emit(text);
	}

	deleteMessage(messageId: string) {
		this.localtron.chatMessageDelete(messageId);
		this.messages = this.messages.filter((m) => m.id !== messageId);
	}

	getLastUserMessage(): ChatMessage | undefined {
		let userMessages = this.messages?.filter((message) => {
			return message.isUserMessage;
		});
		let length = userMessages?.length;
		return length ? userMessages[length - 1] : undefined;
	}

	regenerateAnswer(message: ChatMessage) {
		if (message.isUserMessage) {
			return;
		}
		this.deleteMessage(message.id as string);
		let lastUserMessage = this.getLastUserMessage();
		if (lastUserMessage) {
			this.prompt(lastUserMessage.messageContent);
		}
	}
}

function escapeHtml(unsafe) {
	return unsafe
		.replace(/&/g, '&amp;')
		.replace(/</g, '&lt;')
		.replace(/>/g, '&gt;')
		.replace(/"/g, '&quot;')
		.replace(/'/g, '&#039;');
}

function zigzagArrays<T>(array1: T[], array2: T[]): T[] {
	const result: T[] = [];
	const maxLength = Math.max(array1.length, array2.length);

	for (let i = 0; i < maxLength; i++) {
		if (i < array1.length) {
			result.push(array1[i]);
		}
		if (i < array2.length) {
			result.push(array2[i]);
		}
	}

	return result;
}

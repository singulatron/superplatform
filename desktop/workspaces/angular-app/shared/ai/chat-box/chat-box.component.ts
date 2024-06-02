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
import { ApiService } from '../../stdlib/api.service';

import { LocaltronService } from '../../../src/app/services/localtron.service';
import {
	ChatService,
	ChatThread,
	ChatMessage,
} from '../../../src/app/services/chat.service';
import {
	Prompt,
	PromptService,
} from '../../../src/app/services/prompt.service';
import { Model } from '../../../src/app/services/model.service';

import { ElectronAppService } from '../../../src/app/services/electron-app.service';
import { ConfigService } from '../../../src/app/services/config.service';

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

	@Input() thread!: ChatThread;

	@Output() onThreadUpdate = new EventEmitter<ChatThread>();
	@Output() onFirstMessageSend = new EventEmitter<ChatThread>();
	@Output() onCopyToClipboard = new EventEmitter<string>();

	private model: Model | undefined;
	private models: Model[] = [];
	public promptQueue: Prompt[] = [];

	public message: string = '';
	public messages: ChatMessage[] = [];
	public messageCurrentlyStreamed = '';

	constructor(
		private api: ApiService,
		private localtron: LocaltronService,
		public lapi: ElectronAppService,
		private cd: ChangeDetectorRef,
		private configService: ConfigService,
		private promptService: PromptService,
		private chatService: ChatService
	) {}

	private subscriptions: Subscription[] = [];

	async ngOnInit() {
		this.models = await this.api.getModels();
		this.subscriptions.push(
			this.configService.onConfigUpdate$.subscribe(async (config) => {
				this.model = this.models?.find(
					(m) => m.id == config?.model?.currentModelId
				);
			})
		);
		this.subscriptions.push(
			this.chatService.onChatMessageAdded$.subscribe(async (event) => {
				if (this.thread?.id && this.thread.id == event.threadId) {
					let rsp = await this.chatService.chatMessages(this.thread?.id);
					this.messages = rsp.messages;
				}
			})
		);
	}

	streamSubscription!: Subscription;
	promptSubscription!: Subscription;

	ngOnDestroy() {
		this.streamSubscription.unsubscribe();
		this.subscriptions.forEach((s) => {
			s.unsubscribe();
		});
	}

	async ngOnChanges(changes: SimpleChanges): Promise<void> {
		if (changes.thread) {
			// @todo investigate this if only the ID changed
			if (this.streamSubscription) {
				this.streamSubscription.unsubscribe();
			}
			if (this.promptSubscription) {
				this.promptSubscription.unsubscribe();
			}

			let threadId: string;

			if (!this.thread) {
				this.thread = {
					id: this.localtron.uuid(),
				};
				threadId = this.thread.id as string;
			} else {
				threadId = changes.thread.currentValue.id;
				let rsp = await this.chatService.chatMessages(threadId);
				this.messages = rsp.messages;
			}

			this.promptSubscription =
				this.promptService.onPromptListUpdate$.subscribe((promptList) => {
					let promptQueue = promptList?.filter((p) => {
						return p.threadId == threadId;
					});
					this.promptQueue = promptQueue;
				});

			this.messageCurrentlyStreamed = '';
			let first = true;

			// We are always subscribed to this, even if streaming is not happening
			// right now. There is always one streaming that is subscribed to
			// - the current thread on screen.
			this.streamSubscription = this.promptService
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
							addVal = addVal.trimStart();
							first = false;
						}
						this.messageCurrentlyStreamed += addVal;
					}

					if (
						response?.choices?.length > 0 &&
						response?.choices[0]?.finish_reason === 'stop'
					) {
						this.onThreadUpdate.emit(changes.thread.currentValue);

						if (this.messages?.length == 1) {
							this.setThreadName(this.messages[0].messageContent);
						}
						// @todo might not be needed now we have the `chatMessageAdded`
						// event coming from the firehose
						let rsp = await this.chatService.chatMessages(threadId);
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

		this.sendMessage(msg);
	}

	async sendMessage(msg: string) {
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

		await this.promptService.promptAdd({
			id: this.localtron.uuid(),
			prompt: fullPrompt,
			message: msg,
			threadId: this.thread.id as string,
			modelId: this.model?.id as string,
		});
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

	setThreadName(msg: string) {
		if (!msg) {
			return;
		}
		if (this.thread?.name !== defaultThreadName) {
			return;
		}

		// @todo this we dont talk to LLM locally anymore
		// add syncron prompt to prompt service that talks to localtron

		// let prompt = this.promptTemplate
		// 	? this.promptTemplate.replace(
		// 			'{prompt}',
		// 			this.threadNameSummaryTemplate.replace('{message}', msg)
		// 		)
		// 	: msg;
		// let newThreadName = '';

		//this.promptService
		//	.prompt({
		//		prompt: prompt,
		//		stream: true,
		//	})
		//	.pipe(finalize(() => {}))
		//	.subscribe((response) => {
		//		if (response?.choices?.length > 0 && response?.choices[0]?.text) {
		//			newThreadName += response?.choices[0].text;
		//			this.thread.name = newThreadName;
		//			this.localtron.chatThreadUpdate(this.thread);
		//
		//			this.cd.detectChanges();
		//		}
		//	});
	}

	propagateCopyToClipboard(text: string | undefined) {
		if (text === undefined) {
			return;
		}
		this.onCopyToClipboard.emit(text);
	}

	deleteMessage(messageId: string | undefined) {
		if (messageId === undefined) {
			return;
		}
		this.chatService.chatMessageDelete(messageId);
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
			this.sendMessage(lastUserMessage.messageContent);
		}
	}
}

function escapeHtml(unsafe: string) {
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

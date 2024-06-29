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
	OnChanges,
	SimpleChanges,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	ViewChild,
	ElementRef,
	AfterViewInit,
	OnDestroy,
	ViewContainerRef,
	ComponentRef,
} from '@angular/core';
import { Subscription, filter } from 'rxjs';

import { LocaltronService } from '../../services/localtron.service';
import {
	ChatService,
	ChatThread,
	ChatMessage,
	Asset,
} from '../../services/chat.service';
import { Prompt, PromptService } from '../../services/prompt.service';

import { ElectronAppService } from '../../services/electron-app.service';

import { TranslatePipe } from '../../translate.pipe';
import { FormsModule } from '@angular/forms';
import { MessageComponent } from './message/message.component';
import { NgFor, NgIf, AsyncPipe, NgStyle } from '@angular/common';
import { IonicModule } from '@ionic/angular';
import {
	ChatInputComponent,
	SendOutput,
} from './chat-input/chat-input.component';
import { MobileService } from '../../services/mobile.service';
import { FooterService } from '../../services/footer.service';
import { Router, NavigationEnd } from '@angular/router';

const defaultThreadName = 'New chat';

@Component({
	selector: 'app-chat-box',
	templateUrl: './chat-box.component.html',
	styleUrl: './chat-box.component.scss',
	encapsulation: ViewEncapsulation.None,
	standalone: true,
	imports: [
		IonicModule,
		NgFor,
		MessageComponent,
		NgIf,
		FormsModule,
		TranslatePipe,
		ChatInputComponent,
		AsyncPipe,
		NgStyle,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatBoxComponent implements OnChanges, AfterViewInit, OnDestroy {
	@Input() promptTemplate: string = '[INST] {prompt} [/INST]';

	// @todo push this to the backend too
	@Input() threadNameSummaryTemplate =
		'Summarize, shorten this question in 3-5 words, keep it as a question: {message}';

	@Input() thread!: ChatThread;

	public promptQueue: Prompt[] = [];

	public messages: ChatMessage[] = [];
	public assets: Asset[] = [];
	public messageCurrentlyStreamed: ChatMessage = {} as any;
	private subscriptions: Subscription[] = [];

	@ViewChild('footerContainer', { read: ViewContainerRef, static: true })
	container!: ViewContainerRef;

	// eslint-disable-next-line
	private footerComponentRef: ComponentRef<ChatInputComponent> | null = null;

	@ViewChild('scrollableElement') private scrollableElement!: ElementRef;
	private scrollListener!: () => void;
	private shouldScrollToBottom = true;
	private mutationObserver!: MutationObserver;

	constructor(
		private localtron: LocaltronService,
		public lapi: ElectronAppService,
		private cd: ChangeDetectorRef,
		private promptService: PromptService,
		private chatService: ChatService,
		public mobile: MobileService,
		public footer: FooterService,

		private router: Router
	) {}

	getFooterComponent(): ComponentRef<ChatInputComponent> {
		if (this.footerComponentRef) {
			return this.footerComponentRef;
		}

		this.footerComponentRef =
			this.container.createComponent(ChatInputComponent);

		const chatInputComponentInstance = this.footerComponentRef.instance;

		chatInputComponentInstance.sends.subscribe((event) => {
			this.handleSend(event);
		});

		return this.footerComponentRef;
	}

	async ngOnInit() {
		if (this.router.url === '/chat' && this.mobile.getMobileStatus()) {
			this.footer.updateFooterComponent(this.getFooterComponent());
		}

		this.subscriptions.push(
			this.router.events
				.pipe(filter((event): event is NavigationEnd => event instanceof NavigationEnd))
				.subscribe((navEnd) => {
					if (navEnd.url === '/chat' && this.mobile.getMobileStatus()) {
						this.footer.updateFooterComponent(this.getFooterComponent());
					}
				})
		);

		if (this.thread?.id) {
			const rsp = await this.chatService.chatMessages(this.thread.id);
			this.messages = rsp.messages;
			this.assets = rsp.assets;
			// The mutationObserver triggers before the app-messages components are rendered.
			// This ensures scrollToBottom is called when the app loads for the first time,
			// after the app-messages have been rendered.
			// TODO: Find a better solution
			setTimeout(async () => {
				this.scrollToBottom();
			}, 1000)
		}

		this.cd.markForCheck();

		this.subscriptions.push(
			this.chatService.onChatMessageAdded$.subscribe(async (event) => {
				if (this.thread?.id && this.thread.id == event.threadId) {
					const rsp = await this.chatService.chatMessages(this.thread?.id);
					this.messages = rsp.messages;
					this.assets = rsp.assets;
					this.cd.markForCheck();
				}
			})
		);
	}

	ngAfterViewInit() {
		this.mutationObserver = new MutationObserver(() => {
			this.scrollToBottom();
		});

		this.mutationObserver.observe(this.scrollableElement.nativeElement, {
			childList: true,
			subtree: true,
		});
		this.scrollListener = this.onScroll.bind(this);
		this.scrollableElement.nativeElement.addEventListener('scroll', this.scrollListener);
	}

	ngOnDestroy() {
		try {
			this.scrollableElement?.nativeElement?.removeEventListener('scroll', this.scrollListener);
		} catch (error) { }
		try {
			this.mutationObserver.disconnect();
		} catch (error) { }
	}

	private onScroll(): void {
		const element = this.scrollableElement.nativeElement;
		const atBottom = element.scrollHeight - element.scrollTop < (element.clientHeight + element.clientHeight * 0.05);
		console.log("aha", atBottom)
		this.shouldScrollToBottom = atBottom;
	}

	getAssets(message: ChatMessage): Asset[] {
		return this.assets?.filter((a) => message.assetIds?.includes(a.id));
	}

	async handleSend(emitted: SendOutput) {
		if (!this.thread?.title) {
			this.thread.title = emitted.message.slice(0, 100);
		}

		await this.promptService.promptAdd({
			id: this.localtron.uuid(),
			prompt: emitted.message,
			characterId: emitted.characterId,
			template: this.promptTemplate,
			threadId: this.thread.id as string,
			modelId: emitted.modelId as string,
		});

		this.cd.markForCheck();
	}

	streamSubscription!: Subscription;
	promptSubscription!: Subscription;

	ionViewWillLeave() {
		this.streamSubscription.unsubscribe();
		for (const s of this.subscriptions) {
			s.unsubscribe();
		}
	}

	async ngOnChanges(changes: SimpleChanges): Promise<void> {
		if (changes.thread) {
			this.shouldScrollToBottom = true;
			this.messages = [];
			this.assets = [];
			this.cd.markForCheck();

			if (this.streamSubscription) {
				this.streamSubscription.unsubscribe();
			}
			if (this.promptSubscription) {
				this.promptSubscription.unsubscribe();
			}

			let threadId: string;

			if (this.thread) {
				threadId = changes.thread.currentValue.id;
				const rsp = await this.chatService.chatMessages(threadId);
				this.messages = rsp.messages;
				this.assets = rsp.assets;
			} else {
				this.thread = {
					id: this.localtron.uuid(),
				};
				threadId = this.thread.id as string;
			}

			this.promptSubscription =
				this.promptService.onPromptListUpdate$.subscribe((promptList) => {
					const promptQueue = promptList?.filter((p) => {
						return p.threadId == threadId;
					});
					this.promptQueue = promptQueue;
					this.cd.markForCheck();
				});

			this.messageCurrentlyStreamed.content = '';
			let first = true;

			this.cd.markForCheck();

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
						const insidePre =
							(this.messageCurrentlyStreamed.content.match(/```/g) || [])
								.length %
							2 ===
							1;
						let addValue = insidePre
							? response?.choices[0].text
							: escapeHtml(response?.choices[0].text);

						if (first) {
							addValue = addValue.trimStart();
							first = false;
						}

						this.messageCurrentlyStreamed = {
							...this.messageCurrentlyStreamed,
							content: this.messageCurrentlyStreamed.content + addValue,
						} as any;
					}

					if (
						response?.choices?.length > 0 &&
						response?.choices[0]?.finish_reason === 'stop'
					) {
						if (this.messages?.length == 1) {
							this.setThreadName(this.messages[0].content);
						}
						// @todo might not be needed now we have the `chatMessageAdded`
						// event coming from the firehose
						const rsp = await this.chatService.chatMessages(threadId);
						this.messages = rsp.messages;
						this.assets = rsp.assets;

						this.messageCurrentlyStreamed = {
							...this.messageCurrentlyStreamed,
							content: '',
						} as any;
					}
					this.cd.detectChanges();
				});
		}
	}

	setThreadName(message: string) {
		if (!message) {
			return;
		}
		if (this.thread?.title !== defaultThreadName) {
			return;
		}
		// @todo summarize with llm at the end of the streaming
	}

	trackById(_: number, message: { id?: string }): string {
		return message.id || '';
	}

	removePromptFromQueue(prompt: Prompt): void {
		this.promptService.promptRemove(prompt)
	}

	private scrollToBottom(): void {
		if (!this.shouldScrollToBottom) {
			return;
		}
		try {
			this.scrollableElement.nativeElement.scrollTop = this.scrollableElement.nativeElement.scrollHeight;
		} catch (err) {
			console.error('Scroll to bottom failed:', err);
		}
	}
}

function escapeHtml(unsafe: string) {
	return unsafe
		.replaceAll('&', '&amp;')
		.replaceAll('<', '&lt;')
		.replaceAll('>', '&gt;')
		.replaceAll('"', '&quot;')
		.replaceAll("'", '&#039;');
}

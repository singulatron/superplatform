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
import { Component, OnInit } from '@angular/core';
import { LocaltronService } from '../services/localtron.service';

import { ElectronIpcService } from '../services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { Subscription } from 'rxjs';
import { ApiService } from '../../../shared/stdlib/api.service';
import { ChatService, ChatThread, ChatMessage } from '../services/chat.service';
import { Prompt, PromptService } from '../services/prompt.service';
import { Model } from '../services/model.service';
import { ConfigService } from '../services/config.service';

@Component({
	selector: 'app-chat',
	templateUrl: './chat.component.html',
	styleUrl: './chat.component.scss',
})
export class ChatComponent implements OnInit {
	public defaultPrompt = '[INST] {prompt} [/INST]';
	public chatThreads: Array<ChatThread> = [];
	public activeThread: ChatThread | null = null;
	public messages: ChatMessage[] = [];

	public model: Model | null = null;
	private models: Model[] = [];

	private subscriptions: Subscription[] = [];

	constructor(
		private localtron: LocaltronService,
		private chatService: ChatService,
		private configService: ConfigService,
		public promptService: PromptService,
		private api: ApiService,
		private ipcService: ElectronIpcService
	) {}

	async ngOnInit() {
		await this.refreshThreadList();

		let activeThreadId = this.chatService.getActiveThreadId();
		if (activeThreadId) {
			let activeThread = this.chatThreads?.find((v) => v.id === activeThreadId);
			if (activeThread) {
				this.activeThread = activeThread;
			}
		}
		if (!this.activeThread && this.chatThreads?.length) {
			this.activeThread = this.chatThreads[0];
		}
		if (!this.activeThread) {
			this.activeThread = {
				id: this.localtron.uuid(),
			};
		}

		this.models = await this.api.getModels();
		this.subscriptions.push(
			this.configService.onConfigUpdate$.subscribe((conf) => {
				let model = this.models?.find(
					(m) => m.id == conf?.model?.currentModelId
				);
				if (model) {
					this.model = model;
				}
			})
		);
	}

	ngOnDestroy() {
		this.subscriptions.forEach((sub) => sub.unsubscribe());
	}

	public async setThreadAsActive(thread: ChatThread) {
		this.activeThread = thread;
		console.debug('Loading thread', {
			threadId: thread.id,
		});
		if (!thread.id) {
			return;
		}
		let rsp = await this.chatService.chatMessages(thread.id);
		this.messages = rsp.messages;
		this.chatService.setActiveThreadId(thread.id);
	}

	public num(
		threadId: string | undefined,
		promptList: Prompt[] | null
	): number {
		if (!promptList) {
			return -1;
		}
		if (!threadId) {
			return -1;
		}
		let ind = -1;
		promptList?.forEach((p, index) => {
			if (p.threadId == threadId) {
				ind = index;
			}
		});
		return ind;
	}

	public async openNewThread() {
		this.activeThread = {
			id: this.localtron.uuid(),
		};
		console.debug('Opened empty thread', {
			threadId: this.activeThread.id,
		});
		this.messages = [];
	}

	public removeChatThread(thread: ChatThread) {
		if (!thread.id) {
			return;
		}
		this.chatService.chatThreadDelete(thread.id);
		this.refreshThreadList();
	}

	public async refreshThreadList() {
		let rsp = await this.chatService.chatThreads();
		this.chatThreads = rsp.threads;
		if (!this.chatThreads?.length) {
			this.chatThreads = [];
		}
	}

	public onThreadUpdate(updatedThread: ChatThread) {
		if (!updatedThread.id) {
			throw 'Id for thread is required';
		}
		let found = this.chatThreads?.find((thread, index) => {
			if (thread.id === updatedThread.id) {
				console.debug('Updating thread', {
					threadId: thread.id,
				});
				this.chatThreads[index] = updatedThread;
				return true;
			}
			return false;
		});
		if (!found) {
			if (this.chatThreads === undefined) {
				throw 'No threads';
			}
			console.debug('Adding new thread', {
				threadId: updatedThread?.id,
			});
			this.chatThreads.unshift(updatedThread);
		}
		this.chatService.chatThreadUpdate(updatedThread);
		if (!this.activeThread?.id) {
			this.setThreadAsActive(updatedThread);
		}
	}

	public onCopyToClipboard(text: string) {
		this.ipcService.send(WindowApiConst.COPY_TO_CLIPBOARD_REQUEST, text);
	}
}

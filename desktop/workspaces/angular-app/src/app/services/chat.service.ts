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
import { ReplaySubject } from 'rxjs';
import { FirehoseService } from './firehose.service';

@Injectable({
	providedIn: 'root',
})
export class ChatService {
	public activeThreadId: string = '';

	onChatMessageAddedSubject = new ReplaySubject<ChatMessageAddedEvent>(1);
	onChatMessageAdded$ = this.onChatMessageAddedSubject.asObservable();

	onChatThreadAddedSubject = new ReplaySubject<ChatThreadAddedEvent>(1);
	onChatThreadAdded$ = this.onChatMessageAddedSubject.asObservable();

	onChatThreadUpdateSubject = new ReplaySubject<ChatMessageAddedEvent>(1);
	onChatThreadUpdate$ = this.onChatMessageAddedSubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private firehoseService: FirehoseService
	) {
		this.init();
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'chatMessageAdded': {
					this.onChatMessageAddedSubject.next(event.data);
					break;
				}
				case 'chatThreadAdded': {
					this.onChatMessageAddedSubject.next(event.data);
					break;
				}
				case 'chatThreadUpdate': {
					this.onChatThreadUpdateSubject.next(event.data);
					break;
				}
			}
		});
	}

	async chatMessageDelete(messageId: string): Promise<GetChatThreadResponse> {
		const request: DeleteChatMessageRequest = { messageId: messageId };
		return this.localtron.call('/chat/message/delete', request);
	}

	async chatMessages(threadId: string): Promise<GetChatMessagesResponse> {
		const request: GetChatMessagesRequest = { threadId: threadId };
		return this.localtron.call('/chat/messages', request);
	}

	async chatThread(threadId: string): Promise<GetChatThreadResponse> {
		const request: GetChatThreadRequest = { threadId: threadId };
		return this.localtron.call('/chat/thread', request);
	}

	async chatThreadAdd(thread: ChatThread): Promise<AddChatThreadResponse> {
		const request: AddChatThreadRequest = { thread: thread };
		return this.localtron.call('/chat/thread/add', request);
	}

	async chatThreadUpdate(
		thread: ChatThread
	): Promise<UpdateChatThreadResponse> {
		const request: UpdateChatThreadRequest = { thread: thread };
		return this.localtron.call('/chat/thread/update', request);
	}

	async chatThreadDelete(threadId: string): Promise<void> {
		const request: DeleteChatThreadRequest = { threadId: threadId };
		return this.localtron.call('/chat/thread/delete', request);
	}

	async chatThreads(): Promise<GetChatThreadsResponse> {
		const request: GetChatThreadsRequest = {};
		return this.localtron.call('/chat/threads', request);
	}

	setActiveThreadId(id: string) {
		localStorage.setItem(this.activeThreadId, id);
	}

	getActiveThreadId(): string {
		const activeThreadId = localStorage.getItem(this.activeThreadId);
		if (!activeThreadId) {
			return '';
		}
		return activeThreadId;
	}
}

export interface ChatThread {
	id?: string;
	createdAt?: string;
	updatedAt?: string;

	topicIds?: string;
	userIds?: string[];

	title?: string;
}

export interface ChatMessage {
	id?: string;
	createdAt?: string;
	updatedAt?: string;

	threadId: string;
	userId?: string;
	content: string;
	assetIds: string[];
}

export interface Asset {
	id: string;
	url: string;
	/* Some assets might have the content directly in them as base64
	encoded strings */
	content: string;
	type: string;
	decription: string;
	createdAt: string;
	updatedAt: string;
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
	messageId: string;
}

export interface GetChatThreadRequest {
	threadId: string;
}

type GetChatThreadResponse = {
	thread: ChatThread;
};

// eslint-disable-next-line
type GetChatThreadsRequest = {};

type GetChatThreadsResponse = {
	threads: ChatThread[];
};

// eslint-disable-next-line
type GetChatMessagesRequest = {};

type GetChatMessagesResponse = {
	messages: ChatMessage[];
	assets: Asset[];
};

export interface ChatMessageAddedEvent {
	threadId: string;
}

export interface ChatThreadAddedEvent {
	threadId: string;
}

export interface ChatThreadUpdateEvent {
	threadId: string;
}

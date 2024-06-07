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

	constructor(
		private localtron: LocaltronService,
		private firehoseService: FirehoseService
	) {
		this.init();
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'chatMessageAdded':
					this.onChatMessageAddedSubject.next(event.data);
					break;
			}
		});
	}

	async chatMessageDelete(messageId: string): Promise<GetChatThreadResponse> {
		let req: DeleteChatMessageRequest = { messageId: messageId };
		return this.localtron.call('/chat/message/delete', req);
	}

	async chatMessages(threadId: string): Promise<GetChatMessagesResponse> {
		let req: GetChatMessagesRequest = { threadId: threadId };
		return this.localtron.call('/chat/messages', req);
	}

	async chatThread(threadId: string): Promise<GetChatThreadResponse> {
		let req: GetChatThreadRequest = { threadId: threadId };
		return this.localtron.call('/chat/thread', req);
	}

	async chatThreadAdd(thread: ChatThread): Promise<AddChatThreadResponse> {
		let req: AddChatThreadRequest = { thread: thread };
		return this.localtron.call('/chat/thread/add', req);
	}

	async chatThreadUpdate(
		thread: ChatThread
	): Promise<UpdateChatThreadResponse> {
		let req: UpdateChatThreadRequest = { thread: thread };
		return this.localtron.call('/chat/thread/update', req);
	}

	async chatThreadDelete(threadId: string): Promise<void> {
		let req: DeleteChatThreadRequest = { threadId: threadId };
		return this.localtron.call('/chat/thread/delete', req);
	}

	async chatThreads(): Promise<GetChatThreadsResponse> {
		let req: GetChatThreadsRequest = {};
		return this.localtron.call('/chat/threads', req);
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
	topicId?: string;
	title?: string;
	createdAt?: string;
}

export interface ChatMessage {
	id?: string;
	threadId: string;
	messageContent: string;
	isUserMessage: boolean;
	createdAt?: string;
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
	messageId: string;
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

export interface ChatMessageAddedEvent {
	threadId: string;
}

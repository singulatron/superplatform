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

	onMessageAddedSubject = new ReplaySubject<MessageAddedEvent>(1);
	onMessageAdded$ = this.onMessageAddedSubject.asObservable();

	onThreadAddedSubject = new ReplaySubject<ThreadAddedEvent>(1);
	onThreadAdded$ = this.onMessageAddedSubject.asObservable();

	onThreadUpdateSubject = new ReplaySubject<MessageAddedEvent>(1);
	onThreadUpdate$ = this.onMessageAddedSubject.asObservable();

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
					this.onMessageAddedSubject.next(event.data);
					break;
				}
				case 'chatThreadAdded': {
					this.onMessageAddedSubject.next(event.data);
					break;
				}
				case 'chatThreadUpdate': {
					this.onThreadUpdateSubject.next(event.data);
					break;
				}
			}
		});
	}

	async chatMessageDelete(messageId: string): Promise<GetThreadResponse> {
		const request: DeleteMessageRequest = { messageId: messageId };
		return this.localtron.call('/chat/message/delete', request);
	}

	async chatMessages(threadId: string): Promise<GetMessagesResponse> {
		const request: GetMessagesRequest = { threadId: threadId };
		return this.localtron.call('/chat/messages', request);
	}

	async chatThread(threadId: string): Promise<GetThreadResponse> {
		const request: GetThreadRequest = { threadId: threadId };
		return this.localtron.call('/chat/thread', request);
	}

	async chatThreadAdd(thread: Thread): Promise<AddThreadResponse> {
		const request: AddThreadRequest = { thread: thread };
		return this.localtron.call('/chat/thread/add', request);
	}

	async chatThreadUpdate(
		thread: Thread
	): Promise<UpdateThreadResponse> {
		const request: UpdateThreadRequest = { thread: thread };
		return this.localtron.call('/chat/thread/update', request);
	}

	async chatThreadDelete(threadId: string): Promise<void> {
		const request: DeleteThreadRequest = { threadId: threadId };
		return this.localtron.call('/chat/thread/delete', request);
	}

	async chatThreads(): Promise<GetThreadsResponse> {
		const request: GetThreadsRequest = {};
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

export interface Thread {
	id?: string;
	createdAt?: string;
	updatedAt?: string;

	topicIds?: string;
	userIds?: string[];

	title?: string;
}

export interface Message {
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

export interface AddMessageRequest {
	message: Message;
}

export interface AddThreadRequest {
	thread: Thread;
}

export interface AddThreadResponse {
	thread: Thread;
}

export interface UpdateThreadRequest {
	thread: Thread;
}

export interface UpdateThreadResponse {
	thread: Thread;
}

export interface DeleteThreadRequest {
	threadId: string;
}

export interface DeleteMessageRequest {
	messageId: string;
}

export interface GetThreadRequest {
	threadId: string;
}

type GetThreadResponse = {
	thread: Thread;
};

// eslint-disable-next-line
type GetThreadsRequest = {};

type GetThreadsResponse = {
	threads: Thread[];
};

// eslint-disable-next-line
type GetMessagesRequest = {};

type GetMessagesResponse = {
	messages: Message[];
	assets: Asset[];
};

export interface MessageAddedEvent {
	threadId: string;
}

export interface ThreadAddedEvent {
	threadId: string;
}

export interface ThreadUpdateEvent {
	threadId: string;
}

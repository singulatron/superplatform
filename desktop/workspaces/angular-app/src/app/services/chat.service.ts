/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { ReplaySubject, Subject } from 'rxjs';
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

	onStartNewThreadSubject = new Subject<void>();
	// emitted when a new thread should be started
	onStartNewThread$ = this.onStartNewThreadSubject.asObservable();

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

	async chatThreadUpdate(thread: Thread): Promise<UpdateThreadResponse> {
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

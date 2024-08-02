/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { ReplaySubject, Subject } from 'rxjs';
import { FirehoseService } from './firehose.service';
import * as chat from '@singulatron/types';

@Injectable({
	providedIn: 'root',
})
export class ChatService {
	public activeThreadId: string = '';

	onMessageAddedSubject = new ReplaySubject<chat.MessageAddedEvent>(1);
	onMessageAdded$ = this.onMessageAddedSubject.asObservable();

	onThreadAddedSubject = new ReplaySubject<chat.ThreadAddedEvent>(1);
	onThreadAdded$ = this.onMessageAddedSubject.asObservable();

	onThreadUpdateSubject = new ReplaySubject<chat.MessageAddedEvent>(1);
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

	async chatMessageDelete(messageId: string): Promise<chat.GetThreadResponse> {
		return this.localtron.delete('/chat-service/message/' + messageId);
	}

	async chatMessages(threadId: string): Promise<chat.GetMessagesResponse> {
		return this.localtron.get(`/chat-service/thread/${threadId}/messages`);
	}

	async chatThread(threadId: string): Promise<chat.GetThreadResponse> {
		return this.localtron.get(`/chat-service/thread/${threadId}`);
	}

	async chatThreadAdd(thread: chat.Thread): Promise<chat.AddThreadResponse> {
		const request: chat.AddThreadRequest = { thread: thread };
		return this.localtron.post('/chat-service/thread', request);
	}

	async chatThreadUpdate(
		thread: chat.Thread
	): Promise<chat.UpdateThreadResponse> {
		const request: chat.UpdateThreadRequest = { thread: thread };
		return this.localtron.put(`/chat-service/thread/${thread.id}`, request);
	}

	async chatThreadDelete(threadId: string): Promise<void> {
		return this.localtron.delete(`/chat-service/thread/${threadId}`);
	}

	async chatThreads(): Promise<chat.GetThreadsResponse> {
		return this.localtron.post('/chat-service/threads', {});
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

import { Injectable } from '@angular/core';
import { ChatMessageAddedEvent, LocaltronService } from './localtron.service';
import { ReplaySubject } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class ChatService {
	public activeThreadId: string = '';

	onChatMessageAddedSubject = new ReplaySubject<ChatMessageAddedEvent>(1);
	onChatMessageAdded$ = this.onChatMessageAddedSubject.asObservable();

	constructor(private localtron: LocaltronService) {}

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
	name?: string;
	time?: string;
}

export interface ChatMessage {
	id?: string;
	threadId: string;
	messageContent: string;
	isUserMessage: boolean;
	time?: string;
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

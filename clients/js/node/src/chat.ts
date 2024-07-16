import { ClientOptions } from "./util";

export class ChatService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  async chatMessageDelete(messageId: string): Promise<chat.GetThreadResponse> {
    const request: chat.DeleteMessageRequest = { messageId: messageId };
    return this.localtron.call("/chat/message/delete", request);
  }

  async chatMessages(threadId: string): Promise<chat.GetMessagesResponse> {
    const request: chat.GetMessagesRequest = { threadId: threadId };
    return this.localtron.call("/chat/messages", request);
  }

  async chatThread(threadId: string): Promise<chat.GetThreadResponse> {
    const request: chat.GetThreadRequest = { threadId: threadId };
    return this.localtron.call("/chat/thread", request);
  }

  async chatThreadAdd(thread: chat.Thread): Promise<chat.AddThreadResponse> {
    const request: chat.AddThreadRequest = { thread: thread };
    return this.localtron.call("/chat/thread/add", request);
  }

  async chatThreadUpdate(
    thread: chat.Thread
  ): Promise<chat.UpdateThreadResponse> {
    const request: chat.UpdateThreadRequest = { thread: thread };
    return this.localtron.call("/chat/thread/update", request);
  }

  async chatThreadDelete(threadId: string): Promise<void> {
    const request: chat.DeleteThreadRequest = { threadId: threadId };
    return this.localtron.call("/chat/thread/delete", request);
  }

  async chatThreads(): Promise<chat.GetThreadsResponse> {
    const request: chat.GetThreadsRequest = {};
    return this.localtron.call("/chat/threads", request);
  }

  setActiveThreadId(id: string) {
    localStorage.setItem(this.activeThreadId, id);
  }

  getActiveThreadId(): string {
    const activeThreadId = localStorage.getItem(this.activeThreadId);
    if (!activeThreadId) {
      return "";
    }
    return activeThreadId;
  }
}

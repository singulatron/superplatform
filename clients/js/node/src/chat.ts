import { ClientOptions, call } from "./util";
import * as chat from "@singulatron/types";

export class ChatService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any): Promise<any> {
    return call(this.options.address!, this.options.apiKey!, endpoint, request);
  }

  async chatMessageDelete(messageId: string): Promise<chat.GetThreadResponse> {
    const request: chat.DeleteMessageRequest = { messageId: messageId };
    return this.call("/chat/message/delete", request);
  }

  async chatMessages(threadId: string): Promise<chat.GetMessagesResponse> {
    const request: chat.GetMessagesRequest = { threadId: threadId };
    return this.call("/chat/messages", request);
  }

  async chatThread(threadId: string): Promise<chat.GetThreadResponse> {
    const request: chat.GetThreadRequest = { threadId: threadId };
    return this.call("/chat/thread", request);
  }

  async chatThreadAdd(thread: chat.Thread): Promise<chat.AddThreadResponse> {
    const request: chat.AddThreadRequest = { thread: thread };
    return this.call("/chat/thread/add", request);
  }

  async chatThreadUpdate(
    thread: chat.Thread
  ): Promise<chat.UpdateThreadResponse> {
    const request: chat.UpdateThreadRequest = { thread: thread };
    return this.call("/chat/thread/update", request);
  }

  async chatThreadDelete(threadId: string): Promise<void> {
    const request: chat.DeleteThreadRequest = { threadId: threadId };
    return this.call("/chat/thread/delete", request);
  }

  async chatThreads(): Promise<chat.GetThreadsResponse> {
    const request: chat.GetThreadsRequest = {};
    return this.call("/chat/threads", request);
  }
}

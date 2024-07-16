import { ClientOptions } from "./util";
import * as chat from "@singulatron/types";
export declare class ChatService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    chatMessageDelete(messageId: string): Promise<chat.GetThreadResponse>;
    chatMessages(threadId: string): Promise<chat.GetMessagesResponse>;
    chatThread(threadId: string): Promise<chat.GetThreadResponse>;
    chatThreadAdd(thread: chat.Thread): Promise<chat.AddThreadResponse>;
    chatThreadUpdate(thread: chat.Thread): Promise<chat.UpdateThreadResponse>;
    chatThreadDelete(threadId: string): Promise<void>;
    chatThreads(): Promise<chat.GetThreadsResponse>;
}

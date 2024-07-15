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
export type GetThreadResponse = {
    thread: Thread;
};
export type GetThreadsRequest = {};
export type GetThreadsResponse = {
    threads: Thread[];
};
export type GetMessagesRequest = {};
export type GetMessagesResponse = {
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

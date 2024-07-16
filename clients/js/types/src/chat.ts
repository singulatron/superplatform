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

export type GetThreadResponse = {
  thread: Thread;
};

// eslint-disable-next-line
export type GetThreadsRequest = {};

export type GetThreadsResponse = {
  threads: Thread[];
};

// eslint-disable-next-line
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

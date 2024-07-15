export interface Message {
  id?: string;
  createdAt?: string;
  updatedAt?: string;

  threadId: string;
  userId?: string;
  content: string;
  assetIds: string[];
}

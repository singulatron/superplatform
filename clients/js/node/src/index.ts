export * from "./chat";
export * from "./config";
export * from "./docker";
export * from "./download";
export * from "./generic";
export * from "./prompt";
export * from "./user";

import { ClientOptions } from "./util";
import { ChatService } from "./chat";

export class Client {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  chatService(): ChatService {
    return new ChatService(this.options);
  }
}

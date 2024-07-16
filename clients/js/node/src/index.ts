export * from "./chat";
export * from "./config";
export * from "./docker";
export * from "./download";
export * from "./generic";
export * from "./prompt";
export * from "./user";

import { ClientOptions } from "./util";
import { ChatService } from "./chat";
import { ConfigService } from "./config";
import { DockerService } from "./docker";
import { DownloadService } from "./download";
import { GenericService } from "./generic";
import { PromptService } from "./prompt";
import { UserService } from "./user";

export class Client {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  chatService(): ChatService {
    return new ChatService(this.options);
  }

  configService(): ConfigService {
    return new ConfigService(this.options);
  }

  dockerService(): DockerService {
    return new DockerService(this.options);
  }

  downloadService(): DownloadService {
    return new DownloadService(this.options);
  }

  genericService(): GenericService {
    return new GenericService(this.options);
  }

  promptService(): PromptService {
    return new PromptService(this.options);
  }

  userService(): UserService {
    return new UserService(this.options);
  }
}

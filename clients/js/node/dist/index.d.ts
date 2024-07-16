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
export declare class Client {
    private options;
    constructor(options: ClientOptions);
    chatService(): ChatService;
    configService(): ConfigService;
    dockerService(): DockerService;
    downloadService(): DownloadService;
    genericService(): GenericService;
    promptService(): PromptService;
    userService(): UserService;
}

import { ChatService } from './chat.mjs';
import { ConfigService } from './config.mjs';
import { DockerService } from './docker.mjs';
import { DownloadService } from './download.mjs';
import { GenericService } from './generic.mjs';
import { PromptService } from './prompt.mjs';
import { UserService } from './user.mjs';
import './util2.mjs';
import 'axios';
import '@singulatron/types';

class Client {
    constructor(options) {
        this.options = options;
    }
    chatService() {
        return new ChatService(this.options);
    }
    configService() {
        return new ConfigService(this.options);
    }
    dockerService() {
        return new DockerService(this.options);
    }
    downloadService() {
        return new DownloadService(this.options);
    }
    genericService() {
        return new GenericService(this.options);
    }
    promptService() {
        return new PromptService(this.options);
    }
    userService() {
        return new UserService(this.options);
    }
}

export { ChatService, Client, ConfigService, DockerService, DownloadService, GenericService, PromptService, UserService };

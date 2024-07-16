import { c as call, _ as __awaiter } from './util2.mjs';
import 'axios';

class ChatService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options.address, this.options.apiKey, endpoint, request);
    }
    chatMessageDelete(messageId) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { messageId: messageId };
            return this.call("/chat/message/delete", request);
        });
    }
    chatMessages(threadId) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { threadId: threadId };
            return this.call("/chat/messages", request);
        });
    }
    chatThread(threadId) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { threadId: threadId };
            return this.call("/chat/thread", request);
        });
    }
    chatThreadAdd(thread) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { thread: thread };
            return this.call("/chat/thread/add", request);
        });
    }
    chatThreadUpdate(thread) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { thread: thread };
            return this.call("/chat/thread/update", request);
        });
    }
    chatThreadDelete(threadId) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { threadId: threadId };
            return this.call("/chat/thread/delete", request);
        });
    }
    chatThreads() {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {};
            return this.call("/chat/threads", request);
        });
    }
}

export { ChatService };

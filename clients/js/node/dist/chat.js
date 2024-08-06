'use strict';

var util = require('./util2.js');
require('axios');

class ChatService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options, endpoint, request);
    }
    chatMessageDelete(messageId) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { messageId: messageId };
            return this.call("/chat-svc/message/delete", request);
        });
    }
    chatMessages(threadId) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { threadId: threadId };
            return this.call("/chat-svc/messages", request);
        });
    }
    chatThread(threadId) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { threadId: threadId };
            return this.call("/chat-svc/thread", request);
        });
    }
    chatThreadAdd(thread) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { thread: thread };
            return this.call("/chat-svc/thread/add", request);
        });
    }
    chatThreadUpdate(thread) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { thread: thread };
            return this.call("/chat-svc/thread/update", request);
        });
    }
    chatThreadDelete(threadId) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { threadId: threadId };
            return this.call("/chat-svc/thread/delete", request);
        });
    }
    chatThreads() {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = {};
            return this.call("/chat-svc/threads", request);
        });
    }
}

exports.ChatService = ChatService;

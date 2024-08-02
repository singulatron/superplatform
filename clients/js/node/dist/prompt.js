'use strict';

var util = require('./util2.js');
require('axios');

class PromptService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options, endpoint, request);
    }
    promptAdd(prompt) {
        return util.__awaiter(this, void 0, void 0, function* () {
            if (!prompt.id) {
                prompt.id = util.uuid();
            }
            return this.call("'/prompt-service/add", prompt);
        });
    }
    promptRemove(promptId) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { promptId: promptId };
            return this.call("'/prompt-service/remove", request);
        });
    }
    promptList(request) {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("'/prompt-service/list", request);
        });
    }
}

exports.PromptService = PromptService;

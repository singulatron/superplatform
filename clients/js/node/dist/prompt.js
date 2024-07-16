'use strict';

var util = require('./util2.js');
require('axios');

class PromptService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options.address, this.options.apiKey, endpoint, request);
    }
    promptAdd(prompt) {
        return util.__awaiter(this, void 0, void 0, function* () {
            if (!prompt.id) {
                prompt.id = util.uuid();
            }
            const request = { prompt: prompt };
            return this.call("/prompt/add", request);
        });
    }
    promptRemove(prompt) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = { prompt: prompt };
            return this.call("/prompt/remove", request);
        });
    }
    promptList(request) {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/prompt/list", request);
        });
    }
}

exports.PromptService = PromptService;

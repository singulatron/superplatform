import { c as call, _ as __awaiter, u as uuid } from './util2.mjs';
import 'axios';

class PromptService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options.address, this.options.apiKey, endpoint, request);
    }
    promptAdd(prompt) {
        return __awaiter(this, void 0, void 0, function* () {
            if (!prompt.id) {
                prompt.id = uuid();
            }
            const request = { prompt: prompt };
            return this.call("/prompt/add", request);
        });
    }
    promptRemove(prompt) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { prompt: prompt };
            return this.call("/prompt/remove", request);
        });
    }
    promptList(request) {
        return __awaiter(this, void 0, void 0, function* () {
            return this.call("/prompt/list", request);
        });
    }
}

export { PromptService };

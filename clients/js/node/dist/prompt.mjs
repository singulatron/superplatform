import { c as call, _ as __awaiter, u as uuid } from './util2.mjs';
import 'axios';

class PromptService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options, endpoint, request);
    }
    promptAdd(prompt) {
        return __awaiter(this, void 0, void 0, function* () {
            if (!prompt.id) {
                prompt.id = uuid();
            }
            return this.call("/prompt/add", prompt);
        });
    }
    promptRemove(promptId) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = { promptId: promptId };
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

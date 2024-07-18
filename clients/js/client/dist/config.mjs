import { c as call, _ as __awaiter } from './util2.mjs';
import 'axios';

class ConfigService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options, endpoint, request);
    }
    configGet() {
        return __awaiter(this, void 0, void 0, function* () {
            return yield this.call("/config/get", {});
        });
    }
}

export { ConfigService };

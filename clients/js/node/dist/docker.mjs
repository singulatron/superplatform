import { c as call, _ as __awaiter } from './util2.mjs';
import 'axios';

class DockerService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options.address, this.options.apiKey, endpoint, request);
    }
    dockerInfo() {
        return __awaiter(this, void 0, void 0, function* () {
            return this.call("/docker/info", {});
        });
    }
}

export { DockerService };

import { c as call, _ as __awaiter } from './util2.mjs';
import 'axios';

class DownloadService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options, endpoint, request);
    }
    do(url) {
        return __awaiter(this, void 0, void 0, function* () {
            this.call("/download/do", { url: url });
        });
    }
    pause(url) {
        return __awaiter(this, void 0, void 0, function* () {
            this.call("/download/pause", { url: url });
        });
    }
    list() {
        return __awaiter(this, void 0, void 0, function* () {
            return this.call("/download/list", {});
        });
    }
}

export { DownloadService };

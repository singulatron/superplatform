import { c as call, _ as __awaiter } from './util2.mjs';
import 'axios';

class GenericService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options, endpoint, request);
    }
    create(request) {
        return __awaiter(this, void 0, void 0, function* () {
            return this.call("/generic-svc/create", request);
        });
    }
    find(options) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = options;
            return this.call("/generic-svc/find", request);
        });
    }
    upsert(object) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                object: object,
            };
            return this.call("/generic-svc/upsert", request);
        });
    }
    update(table, conditions, object) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
                object: object,
            };
            return this.call("/generic-svc/update", request);
        });
    }
    delete(table, conditions) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
            };
            return this.call("/generic-svc/delete", request);
        });
    }
}

export { GenericService };

import { c as call, _ as __awaiter } from './util2.mjs';
import 'axios';

class GenericService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return call(this.options.address, this.options.apiKey, endpoint, request);
    }
    create(table, object) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                object: object,
            };
            return this.call("/generic/create", request);
        });
    }
    find(table, conditions) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
            };
            return this.call("/generic/find", request);
        });
    }
    upsert(table, object) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                object: object,
            };
            return this.call("/generic/upsert", request);
        });
    }
    update(table, conditions, object) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
                object: object,
            };
            return this.call("/generic/update", request);
        });
    }
    delete(table, conditions) {
        return __awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
            };
            return this.call("/generic/delete", request);
        });
    }
}

export { GenericService };

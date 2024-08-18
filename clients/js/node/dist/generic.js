'use strict';

var util = require('./util2.js');
require('axios');

class DynamicService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options, endpoint, request);
    }
    create(request) {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/dynamic-svc/create", request);
        });
    }
    find(options) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = options;
            return this.call("/dynamic-svc/find", request);
        });
    }
    upsert(object) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = {
                object: object,
            };
            return this.call("/dynamic-svc/upsert", request);
        });
    }
    update(table, conditions, object) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
                object: object,
            };
            return this.call("/dynamic-svc/update", request);
        });
    }
    delete(table, conditions) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
            };
            return this.call("/dynamic-svc/delete", request);
        });
    }
}

exports.DynamicService = DynamicService;

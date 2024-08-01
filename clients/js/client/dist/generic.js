'use strict';

var util = require('./util2.js');
require('axios');

class GenericService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options, endpoint, request);
    }
    create(request) {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/generic-service/create", request);
        });
    }
    find(options) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = options;
            return this.call("/generic-service/find", request);
        });
    }
    upsert(object) {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/generic-service/upsert", object);
        });
    }
    update(table, conditions, object) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
                object: object,
            };
            return this.call("/generic-service/update", request);
        });
    }
    delete(table, conditions) {
        return util.__awaiter(this, void 0, void 0, function* () {
            const request = {
                table: table,
                conditions: conditions,
            };
            return this.call("/generic-service/delete", request);
        });
    }
}

exports.GenericService = GenericService;

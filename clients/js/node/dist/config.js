'use strict';

var util = require('./util2.js');
require('axios');

class ConfigService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options.address, this.options.apiKey, endpoint, request);
    }
    configGet() {
        return util.__awaiter(this, void 0, void 0, function* () {
            return yield this.call("/config/get", {});
        });
    }
}

exports.ConfigService = ConfigService;

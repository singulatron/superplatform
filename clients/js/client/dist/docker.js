'use strict';

var util = require('./util2.js');
require('axios');

class DockerService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options, endpoint, request);
    }
    dockerInfo() {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/docker/info", {});
        });
    }
}

exports.DockerService = DockerService;

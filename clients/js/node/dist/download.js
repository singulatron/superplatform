'use strict';

var util = require('./util2.js');
require('axios');

class DownloadService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options, endpoint, request);
    }
    do(url) {
        return util.__awaiter(this, void 0, void 0, function* () {
            this.call("/download/do", { url: url });
        });
    }
    pause(url) {
        return util.__awaiter(this, void 0, void 0, function* () {
            this.call("/download/pause", { url: url });
        });
    }
    list() {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/download/list", {});
        });
    }
    get(url) {
        return util.__awaiter(this, void 0, void 0, function* () {
            return this.call("/download/get", {
                url: url,
            });
        });
    }
}

exports.DownloadService = DownloadService;

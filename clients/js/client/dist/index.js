'use strict';

var chat = require('./chat.js');
var config = require('./config.js');
var docker = require('./docker.js');
var download = require('./download.js');
var generic = require('./generic.js');
var prompt = require('./prompt.js');
var user = require('./user.js');
require('./util2.js');
require('axios');
require('@singulatron/types');

class Client {
    constructor(options) {
        this.options = options;
    }
    chatService() {
        return new chat.ChatService(this.options);
    }
    configService() {
        return new config.ConfigService(this.options);
    }
    dockerService() {
        return new docker.DockerService(this.options);
    }
    downloadService() {
        return new download.DownloadService(this.options);
    }
    genericService() {
        return new generic.GenericService(this.options);
    }
    promptService() {
        return new prompt.PromptService(this.options);
    }
    userService() {
        return new user.UserService(this.options);
    }
}

exports.ChatService = chat.ChatService;
exports.ConfigService = config.ConfigService;
exports.DockerService = docker.DockerService;
exports.DownloadService = download.DownloadService;
exports.GenericService = generic.GenericService;
exports.PromptService = prompt.PromptService;
exports.UserService = user.UserService;
exports.Client = Client;

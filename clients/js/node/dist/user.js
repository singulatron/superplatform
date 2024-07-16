'use strict';

var util = require('./util2.js');
var types = require('@singulatron/types');
require('axios');

class UserService {
    constructor(options) {
        this.options = options;
    }
    call(endpoint, request) {
        return util.call(this.options.address, this.options.apiKey, endpoint, request);
    }
    login(email, password) {
        return this.call("/user/login", {
            email: email,
            password: password,
        });
    }
    readUserByToken(token) {
        return this.call("/user/read-user-by-token", {
            token: token,
        });
    }
    getUsers(request) {
        return this.call("/user/get-users", request);
    }
    /** Save profile on behalf of a user */
    saveProfile(email, name) {
        const request = {
            email: email,
            name: name,
        };
        return this.call("/user/save-profile", request);
    }
    changePassword(email, currentPassword, newPassword) {
        const request = {
            email: email,
            currentPassword: currentPassword,
            newPassword: newPassword,
        };
        return this.call("/user/change-password", request);
    }
    changePasswordAdmin(email, newPassword) {
        const request = {
            email: email,
            newPassword: newPassword,
        };
        return this.call("/user/change-password-admin", request);
    }
    /** Create a user - alternative to registration
     */
    createUser(user, password, roleIds) {
        const request = {
            user: user,
            password: password,
            roleIds: roleIds,
        };
        return this.call("/user/create-user", request);
    }
    getRoles() {
        return this.call("/user/get-roles", {});
    }
    getPermissions() {
        return this.call("/user/get-permissions", {});
    }
    setRolePermissions(roleId, permissionIds) {
        const request = {
            roleId: roleId,
            permissionIds: permissionIds,
        };
        return this.call("/user/set-role-permissions", request);
    }
    deleteRole(roleId) {
        const request = {
            roleId: roleId,
        };
        return this.call("/user/delete-role", request);
    }
    deleteUser(userId) {
        const request = {
            userId: userId,
        };
        return this.call("/user/delete-user", request);
    }
    getUser(id) {
        return util.__awaiter(this, void 0, void 0, function* () {
            let rsp = yield this.getUsers({
                query: {
                    conditions: [types.equal(types.field("id"), id)],
                },
            });
            return rsp.users[0];
        });
    }
}

exports.UserService = UserService;

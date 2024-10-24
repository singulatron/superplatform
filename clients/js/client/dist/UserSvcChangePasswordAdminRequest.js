'use strict';

/* tslint:disable */
/* eslint-disable */
/**
 * Superplatform
 * On-premise AI platform and microservices ecosystem.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UserSvcChangePasswordAdminRequest interface.
 */
function instanceOfUserSvcChangePasswordAdminRequest(value) {
    return true;
}
function UserSvcChangePasswordAdminRequestFromJSON(json) {
    return UserSvcChangePasswordAdminRequestFromJSONTyped(json);
}
function UserSvcChangePasswordAdminRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'newPassword': json['newPassword'] == null ? undefined : json['newPassword'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}
function UserSvcChangePasswordAdminRequestToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'newPassword': value['newPassword'],
        'slug': value['slug'],
    };
}

exports.UserSvcChangePasswordAdminRequestFromJSON = UserSvcChangePasswordAdminRequestFromJSON;
exports.UserSvcChangePasswordAdminRequestFromJSONTyped = UserSvcChangePasswordAdminRequestFromJSONTyped;
exports.UserSvcChangePasswordAdminRequestToJSON = UserSvcChangePasswordAdminRequestToJSON;
exports.instanceOfUserSvcChangePasswordAdminRequest = instanceOfUserSvcChangePasswordAdminRequest;

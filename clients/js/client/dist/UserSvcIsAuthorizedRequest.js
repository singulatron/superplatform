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
 * Check if a given object implements the UserSvcIsAuthorizedRequest interface.
 */
function instanceOfUserSvcIsAuthorizedRequest(value) {
    return true;
}
function UserSvcIsAuthorizedRequestFromJSON(json) {
    return UserSvcIsAuthorizedRequestFromJSONTyped(json);
}
function UserSvcIsAuthorizedRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'contactsGranted': json['contactsGranted'] == null ? undefined : json['contactsGranted'],
        'slugsGranted': json['slugsGranted'] == null ? undefined : json['slugsGranted'],
    };
}
function UserSvcIsAuthorizedRequestToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'contactsGranted': value['contactsGranted'],
        'slugsGranted': value['slugsGranted'],
    };
}

exports.UserSvcIsAuthorizedRequestFromJSON = UserSvcIsAuthorizedRequestFromJSON;
exports.UserSvcIsAuthorizedRequestFromJSONTyped = UserSvcIsAuthorizedRequestFromJSONTyped;
exports.UserSvcIsAuthorizedRequestToJSON = UserSvcIsAuthorizedRequestToJSON;
exports.instanceOfUserSvcIsAuthorizedRequest = instanceOfUserSvcIsAuthorizedRequest;

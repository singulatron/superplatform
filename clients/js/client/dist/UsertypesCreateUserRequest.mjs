import { UsertypesUserFromJSON, UsertypesUserToJSON } from './UsertypesUser.mjs';

/* tslint:disable */
/* eslint-disable */
/**
 * Singulatron
 * Run and develop self-hosted AI apps. Your programmable in-house GPT. The Firebase for the AI age.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the UsertypesCreateUserRequest interface.
 */
function instanceOfUsertypesCreateUserRequest(value) {
    return true;
}
function UsertypesCreateUserRequestFromJSON(json) {
    return UsertypesCreateUserRequestFromJSONTyped(json);
}
function UsertypesCreateUserRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'password': json['password'] == null ? undefined : json['password'],
        'roleIds': json['roleIds'] == null ? undefined : json['roleIds'],
        'user': json['user'] == null ? undefined : UsertypesUserFromJSON(json['user']),
    };
}
function UsertypesCreateUserRequestToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'password': value['password'],
        'roleIds': value['roleIds'],
        'user': UsertypesUserToJSON(value['user']),
    };
}

export { UsertypesCreateUserRequestFromJSON, UsertypesCreateUserRequestFromJSONTyped, UsertypesCreateUserRequestToJSON, instanceOfUsertypesCreateUserRequest };
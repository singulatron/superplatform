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
 * Check if a given object implements the UsertypesIsAuthorizedResponse interface.
 */
function instanceOfUsertypesIsAuthorizedResponse(value) {
    return true;
}
function UsertypesIsAuthorizedResponseFromJSON(json) {
    return UsertypesIsAuthorizedResponseFromJSONTyped(json);
}
function UsertypesIsAuthorizedResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'authorized': json['authorized'] == null ? undefined : json['authorized'],
        'user': json['user'] == null ? undefined : UsertypesUserFromJSON(json['user']),
    };
}
function UsertypesIsAuthorizedResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'authorized': value['authorized'],
        'user': UsertypesUserToJSON(value['user']),
    };
}

export { UsertypesIsAuthorizedResponseFromJSON, UsertypesIsAuthorizedResponseFromJSONTyped, UsertypesIsAuthorizedResponseToJSON, instanceOfUsertypesIsAuthorizedResponse };
import { UsertypesPermissionFromJSON, UsertypesPermissionToJSON } from './UsertypesPermission.mjs';

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
 * Check if a given object implements the UsertypesUpserPermissionRequest interface.
 */
function instanceOfUsertypesUpserPermissionRequest(value) {
    return true;
}
function UsertypesUpserPermissionRequestFromJSON(json) {
    return UsertypesUpserPermissionRequestFromJSONTyped(json);
}
function UsertypesUpserPermissionRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'permission': json['permission'] == null ? undefined : UsertypesPermissionFromJSON(json['permission']),
    };
}
function UsertypesUpserPermissionRequestToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'permission': UsertypesPermissionToJSON(value['permission']),
    };
}

export { UsertypesUpserPermissionRequestFromJSON, UsertypesUpserPermissionRequestFromJSONTyped, UsertypesUpserPermissionRequestToJSON, instanceOfUsertypesUpserPermissionRequest };
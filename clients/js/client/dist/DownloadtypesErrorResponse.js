'use strict';

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
 * Check if a given object implements the DownloadtypesErrorResponse interface.
 */
function instanceOfDownloadtypesErrorResponse(value) {
    return true;
}
function DownloadtypesErrorResponseFromJSON(json) {
    return DownloadtypesErrorResponseFromJSONTyped(json);
}
function DownloadtypesErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
function DownloadtypesErrorResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}

exports.DownloadtypesErrorResponseFromJSON = DownloadtypesErrorResponseFromJSON;
exports.DownloadtypesErrorResponseFromJSONTyped = DownloadtypesErrorResponseFromJSONTyped;
exports.DownloadtypesErrorResponseToJSON = DownloadtypesErrorResponseToJSON;
exports.instanceOfDownloadtypesErrorResponse = instanceOfDownloadtypesErrorResponse;
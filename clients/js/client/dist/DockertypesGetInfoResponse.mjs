import { DockertypesDockerInfoFromJSON, DockertypesDockerInfoToJSON } from './DockertypesDockerInfo.mjs';

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
 * Check if a given object implements the DockertypesGetInfoResponse interface.
 */
function instanceOfDockertypesGetInfoResponse(value) {
    return true;
}
function DockertypesGetInfoResponseFromJSON(json) {
    return DockertypesGetInfoResponseFromJSONTyped(json);
}
function DockertypesGetInfoResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'info': json['info'] == null ? undefined : DockertypesDockerInfoFromJSON(json['info']),
    };
}
function DockertypesGetInfoResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'info': DockertypesDockerInfoToJSON(value['info']),
    };
}

export { DockertypesGetInfoResponseFromJSON, DockertypesGetInfoResponseFromJSONTyped, DockertypesGetInfoResponseToJSON, instanceOfDockertypesGetInfoResponse };
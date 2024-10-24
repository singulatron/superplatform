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
 * Check if a given object implements the ModelSvcModelStatus interface.
 */
export function instanceOfModelSvcModelStatus(value) {
    return true;
}
export function ModelSvcModelStatusFromJSON(json) {
    return ModelSvcModelStatusFromJSONTyped(json, false);
}
export function ModelSvcModelStatusFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'address': json['address'] == null ? undefined : json['address'],
        'assetsReady': json['assetsReady'] == null ? undefined : json['assetsReady'],
        'running': json['running'] == null ? undefined : json['running'],
    };
}
export function ModelSvcModelStatusToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'address': value['address'],
        'assetsReady': value['assetsReady'],
        'running': value['running'],
    };
}

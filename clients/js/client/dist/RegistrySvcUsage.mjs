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
 * Check if a given object implements the RegistrySvcUsage interface.
 */
function instanceOfRegistrySvcUsage(value) {
    return true;
}
function RegistrySvcUsageFromJSON(json) {
    return RegistrySvcUsageFromJSONTyped(json);
}
function RegistrySvcUsageFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'percent': json['percent'] == null ? undefined : json['percent'],
        'total': json['total'] == null ? undefined : json['total'],
        'used': json['used'] == null ? undefined : json['used'],
    };
}
function RegistrySvcUsageToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'percent': value['percent'],
        'total': value['total'],
        'used': value['used'],
    };
}

export { RegistrySvcUsageFromJSON, RegistrySvcUsageFromJSONTyped, RegistrySvcUsageToJSON, instanceOfRegistrySvcUsage };

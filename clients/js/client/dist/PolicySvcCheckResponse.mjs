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
 * Check if a given object implements the PolicySvcCheckResponse interface.
 */
function instanceOfPolicySvcCheckResponse(value) {
    if (!('allowed' in value) || value['allowed'] === undefined)
        return false;
    return true;
}
function PolicySvcCheckResponseFromJSON(json) {
    return PolicySvcCheckResponseFromJSONTyped(json);
}
function PolicySvcCheckResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'allowed': json['allowed'],
    };
}
function PolicySvcCheckResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'allowed': value['allowed'],
    };
}

export { PolicySvcCheckResponseFromJSON, PolicySvcCheckResponseFromJSONTyped, PolicySvcCheckResponseToJSON, instanceOfPolicySvcCheckResponse };

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
 * Check if a given object implements the PolicySvcErrorResponse interface.
 */
function instanceOfPolicySvcErrorResponse(value) {
    return true;
}
function PolicySvcErrorResponseFromJSON(json) {
    return PolicySvcErrorResponseFromJSONTyped(json);
}
function PolicySvcErrorResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'error': json['error'] == null ? undefined : json['error'],
    };
}
function PolicySvcErrorResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'error': value['error'],
    };
}

export { PolicySvcErrorResponseFromJSON, PolicySvcErrorResponseFromJSONTyped, PolicySvcErrorResponseToJSON, instanceOfPolicySvcErrorResponse };

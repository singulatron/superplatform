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
 * Check if a given object implements the ConfigSvcModelServiceConfig interface.
 */
export function instanceOfConfigSvcModelServiceConfig(value) {
    return true;
}
export function ConfigSvcModelServiceConfigFromJSON(json) {
    return ConfigSvcModelServiceConfigFromJSONTyped(json, false);
}
export function ConfigSvcModelServiceConfigFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'currentModelId': json['currentModelId'] == null ? undefined : json['currentModelId'],
    };
}
export function ConfigSvcModelServiceConfigToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'currentModelId': value['currentModelId'],
    };
}

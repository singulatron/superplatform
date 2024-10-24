'use strict';

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
function instanceOfConfigSvcModelServiceConfig(value) {
    return true;
}
function ConfigSvcModelServiceConfigFromJSON(json) {
    return ConfigSvcModelServiceConfigFromJSONTyped(json);
}
function ConfigSvcModelServiceConfigFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'currentModelId': json['currentModelId'] == null ? undefined : json['currentModelId'],
    };
}
function ConfigSvcModelServiceConfigToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'currentModelId': value['currentModelId'],
    };
}

exports.ConfigSvcModelServiceConfigFromJSON = ConfigSvcModelServiceConfigFromJSON;
exports.ConfigSvcModelServiceConfigFromJSONTyped = ConfigSvcModelServiceConfigFromJSONTyped;
exports.ConfigSvcModelServiceConfigToJSON = ConfigSvcModelServiceConfigToJSON;
exports.instanceOfConfigSvcModelServiceConfig = instanceOfConfigSvcModelServiceConfig;

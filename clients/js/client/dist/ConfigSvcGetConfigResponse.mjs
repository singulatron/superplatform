import { ConfigSvcConfigFromJSON, ConfigSvcConfigToJSON } from './ConfigSvcConfig.mjs';
import './ConfigSvcAppServiceConfig.mjs';
import './ConfigSvcModelServiceConfig.mjs';
import './ConfigSvcDownloadServiceConfig.mjs';

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
 * Check if a given object implements the ConfigSvcGetConfigResponse interface.
 */
function instanceOfConfigSvcGetConfigResponse(value) {
    return true;
}
function ConfigSvcGetConfigResponseFromJSON(json) {
    return ConfigSvcGetConfigResponseFromJSONTyped(json);
}
function ConfigSvcGetConfigResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'config': json['config'] == null ? undefined : ConfigSvcConfigFromJSON(json['config']),
    };
}
function ConfigSvcGetConfigResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'config': ConfigSvcConfigToJSON(value['config']),
    };
}

export { ConfigSvcGetConfigResponseFromJSON, ConfigSvcGetConfigResponseFromJSONTyped, ConfigSvcGetConfigResponseToJSON, instanceOfConfigSvcGetConfigResponse };

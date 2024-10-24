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
import { ConfigSvcAppServiceConfigFromJSON, ConfigSvcAppServiceConfigToJSON, } from './ConfigSvcAppServiceConfig';
import { ConfigSvcModelServiceConfigFromJSON, ConfigSvcModelServiceConfigToJSON, } from './ConfigSvcModelServiceConfig';
import { ConfigSvcDownloadServiceConfigFromJSON, ConfigSvcDownloadServiceConfigToJSON, } from './ConfigSvcDownloadServiceConfig';
/**
 * Check if a given object implements the ConfigSvcConfig interface.
 */
export function instanceOfConfigSvcConfig(value) {
    return true;
}
export function ConfigSvcConfigFromJSON(json) {
    return ConfigSvcConfigFromJSONTyped(json, false);
}
export function ConfigSvcConfigFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'app': json['app'] == null ? undefined : ConfigSvcAppServiceConfigFromJSON(json['app']),
        'directory': json['directory'] == null ? undefined : json['directory'],
        'download': json['download'] == null ? undefined : ConfigSvcDownloadServiceConfigFromJSON(json['download']),
        'isRuntimeInstalled': json['isRuntimeInstalled'] == null ? undefined : json['isRuntimeInstalled'],
        'model': json['model'] == null ? undefined : ConfigSvcModelServiceConfigFromJSON(json['model']),
    };
}
export function ConfigSvcConfigToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'app': ConfigSvcAppServiceConfigToJSON(value['app']),
        'directory': value['directory'],
        'download': ConfigSvcDownloadServiceConfigToJSON(value['download']),
        'isRuntimeInstalled': value['isRuntimeInstalled'],
        'model': ConfigSvcModelServiceConfigToJSON(value['model']),
    };
}

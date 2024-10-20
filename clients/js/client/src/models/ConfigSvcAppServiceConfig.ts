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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ConfigSvcAppServiceConfig
 */
export interface ConfigSvcAppServiceConfig {
    /**
     * 
     * @type {boolean}
     * @memberof ConfigSvcAppServiceConfig
     */
    loggingDisabled?: boolean;
}

/**
 * Check if a given object implements the ConfigSvcAppServiceConfig interface.
 */
export function instanceOfConfigSvcAppServiceConfig(value: object): value is ConfigSvcAppServiceConfig {
    return true;
}

export function ConfigSvcAppServiceConfigFromJSON(json: any): ConfigSvcAppServiceConfig {
    return ConfigSvcAppServiceConfigFromJSONTyped(json, false);
}

export function ConfigSvcAppServiceConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigSvcAppServiceConfig {
    if (json == null) {
        return json;
    }
    return {
        
        'loggingDisabled': json['loggingDisabled'] == null ? undefined : json['loggingDisabled'],
    };
}

export function ConfigSvcAppServiceConfigToJSON(value?: ConfigSvcAppServiceConfig | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'loggingDisabled': value['loggingDisabled'],
    };
}


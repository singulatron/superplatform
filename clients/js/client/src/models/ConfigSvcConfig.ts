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
import type { ConfigSvcAppServiceConfig } from './ConfigSvcAppServiceConfig';
import {
    ConfigSvcAppServiceConfigFromJSON,
    ConfigSvcAppServiceConfigFromJSONTyped,
    ConfigSvcAppServiceConfigToJSON,
} from './ConfigSvcAppServiceConfig';
import type { ConfigSvcModelServiceConfig } from './ConfigSvcModelServiceConfig';
import {
    ConfigSvcModelServiceConfigFromJSON,
    ConfigSvcModelServiceConfigFromJSONTyped,
    ConfigSvcModelServiceConfigToJSON,
} from './ConfigSvcModelServiceConfig';
import type { ConfigSvcDownloadServiceConfig } from './ConfigSvcDownloadServiceConfig';
import {
    ConfigSvcDownloadServiceConfigFromJSON,
    ConfigSvcDownloadServiceConfigFromJSONTyped,
    ConfigSvcDownloadServiceConfigToJSON,
} from './ConfigSvcDownloadServiceConfig';

/**
 * 
 * @export
 * @interface ConfigSvcConfig
 */
export interface ConfigSvcConfig {
    /**
     * 
     * @type {ConfigSvcAppServiceConfig}
     * @memberof ConfigSvcConfig
     */
    app?: ConfigSvcAppServiceConfig;
    /**
     * 
     * @type {string}
     * @memberof ConfigSvcConfig
     */
    directory?: string;
    /**
     * 
     * @type {ConfigSvcDownloadServiceConfig}
     * @memberof ConfigSvcConfig
     */
    download?: ConfigSvcDownloadServiceConfig;
    /**
     * * This flag drives a minor UX feature:
     * 	 * if the user has not installed the runtime we show an INSTALL
     * 	 * button, but if the user has already installed the runtime we show
     * 	 * we show a START runtime button.
     * 	 *
     * @type {boolean}
     * @memberof ConfigSvcConfig
     */
    isRuntimeInstalled?: boolean;
    /**
     * 
     * @type {ConfigSvcModelServiceConfig}
     * @memberof ConfigSvcConfig
     */
    model?: ConfigSvcModelServiceConfig;
}

/**
 * Check if a given object implements the ConfigSvcConfig interface.
 */
export function instanceOfConfigSvcConfig(value: object): value is ConfigSvcConfig {
    return true;
}

export function ConfigSvcConfigFromJSON(json: any): ConfigSvcConfig {
    return ConfigSvcConfigFromJSONTyped(json, false);
}

export function ConfigSvcConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigSvcConfig {
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

export function ConfigSvcConfigToJSON(value?: ConfigSvcConfig | null): any {
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


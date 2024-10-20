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
 *
 * @export
 * @interface ConfigSvcDownloadServiceConfig
 */
export interface ConfigSvcDownloadServiceConfig {
    /**
     *
     * @type {string}
     * @memberof ConfigSvcDownloadServiceConfig
     */
    downloadFolder?: string;
}
/**
 * Check if a given object implements the ConfigSvcDownloadServiceConfig interface.
 */
export declare function instanceOfConfigSvcDownloadServiceConfig(value: object): value is ConfigSvcDownloadServiceConfig;
export declare function ConfigSvcDownloadServiceConfigFromJSON(json: any): ConfigSvcDownloadServiceConfig;
export declare function ConfigSvcDownloadServiceConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigSvcDownloadServiceConfig;
export declare function ConfigSvcDownloadServiceConfigToJSON(value?: ConfigSvcDownloadServiceConfig | null): any;

/**
 * Singulatron
 * Run and develop self-hosted AI apps. Your programmable in-house GPT. The Firebase for the AI age.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { ConfigtypesConfig } from './ConfigtypesConfig';
/**
 *
 * @export
 * @interface ConfigtypesSaveConfigRequest
 */
export interface ConfigtypesSaveConfigRequest {
    /**
     *
     * @type {ConfigtypesConfig}
     * @memberof ConfigtypesSaveConfigRequest
     */
    config?: ConfigtypesConfig;
}
/**
 * Check if a given object implements the ConfigtypesSaveConfigRequest interface.
 */
export declare function instanceOfConfigtypesSaveConfigRequest(value: object): value is ConfigtypesSaveConfigRequest;
export declare function ConfigtypesSaveConfigRequestFromJSON(json: any): ConfigtypesSaveConfigRequest;
export declare function ConfigtypesSaveConfigRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConfigtypesSaveConfigRequest;
export declare function ConfigtypesSaveConfigRequestToJSON(value?: ConfigtypesSaveConfigRequest | null): any;
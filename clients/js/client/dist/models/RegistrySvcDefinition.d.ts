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
import type { RegistrySvcAPISpec } from './RegistrySvcAPISpec';
import type { RegistrySvcImageSpec } from './RegistrySvcImageSpec';
import type { RegistrySvcClient } from './RegistrySvcClient';
/**
 *
 * @export
 * @interface RegistrySvcDefinition
 */
export interface RegistrySvcDefinition {
    /**
     * API Specs such as OpenAPI definitions etc.
     * @type {Array<RegistrySvcAPISpec>}
     * @memberof RegistrySvcDefinition
     */
    apiSpecs?: Array<RegistrySvcAPISpec>;
    /**
     * Programming language clients such as on npm or GitHub.
     * @type {Array<RegistrySvcClient>}
     * @memberof RegistrySvcDefinition
     */
    clients?: Array<RegistrySvcClient>;
    /**
     *
     * @type {string}
     * @memberof RegistrySvcDefinition
     */
    id: string;
    /**
     * Container specifications for Docker, K8s, etc.
     * @type {RegistrySvcImageSpec}
     * @memberof RegistrySvcDefinition
     */
    image: RegistrySvcImageSpec;
}
/**
 * Check if a given object implements the RegistrySvcDefinition interface.
 */
export declare function instanceOfRegistrySvcDefinition(value: object): value is RegistrySvcDefinition;
export declare function RegistrySvcDefinitionFromJSON(json: any): RegistrySvcDefinition;
export declare function RegistrySvcDefinitionFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcDefinition;
export declare function RegistrySvcDefinitionToJSON(value?: RegistrySvcDefinition | null): any;

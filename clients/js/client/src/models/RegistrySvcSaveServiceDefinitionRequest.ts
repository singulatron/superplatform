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
import type { RegistrySvcServiceDefinition } from './RegistrySvcServiceDefinition';
import {
    RegistrySvcServiceDefinitionFromJSON,
    RegistrySvcServiceDefinitionFromJSONTyped,
    RegistrySvcServiceDefinitionToJSON,
} from './RegistrySvcServiceDefinition';

/**
 * 
 * @export
 * @interface RegistrySvcSaveServiceDefinitionRequest
 */
export interface RegistrySvcSaveServiceDefinitionRequest {
    /**
     * 
     * @type {RegistrySvcServiceDefinition}
     * @memberof RegistrySvcSaveServiceDefinitionRequest
     */
    serviceDefinition?: RegistrySvcServiceDefinition;
}

/**
 * Check if a given object implements the RegistrySvcSaveServiceDefinitionRequest interface.
 */
export function instanceOfRegistrySvcSaveServiceDefinitionRequest(value: object): value is RegistrySvcSaveServiceDefinitionRequest {
    return true;
}

export function RegistrySvcSaveServiceDefinitionRequestFromJSON(json: any): RegistrySvcSaveServiceDefinitionRequest {
    return RegistrySvcSaveServiceDefinitionRequestFromJSONTyped(json, false);
}

export function RegistrySvcSaveServiceDefinitionRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegistrySvcSaveServiceDefinitionRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'serviceDefinition': json['serviceDefinition'] == null ? undefined : RegistrySvcServiceDefinitionFromJSON(json['serviceDefinition']),
    };
}

export function RegistrySvcSaveServiceDefinitionRequestToJSON(value?: RegistrySvcSaveServiceDefinitionRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'serviceDefinition': RegistrySvcServiceDefinitionToJSON(value['serviceDefinition']),
    };
}

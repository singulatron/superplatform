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
import type { DeploySvcStrategyType } from './DeploySvcStrategyType';
import {
    DeploySvcStrategyTypeFromJSON,
    DeploySvcStrategyTypeFromJSONTyped,
    DeploySvcStrategyTypeToJSON,
} from './DeploySvcStrategyType';

/**
 * 
 * @export
 * @interface DeploySvcDeploymentStrategy
 */
export interface DeploySvcDeploymentStrategy {
    /**
     * Max extra replicas during update
     * @type {number}
     * @memberof DeploySvcDeploymentStrategy
     */
    maxSurge?: number;
    /**
     * Max unavailable replicas during update
     * @type {number}
     * @memberof DeploySvcDeploymentStrategy
     */
    maxUnavailable?: number;
    /**
     * Deployment strategy type (RollingUpdate, Recreate, etc.)
     * @type {DeploySvcStrategyType}
     * @memberof DeploySvcDeploymentStrategy
     */
    type?: DeploySvcStrategyType;
}



/**
 * Check if a given object implements the DeploySvcDeploymentStrategy interface.
 */
export function instanceOfDeploySvcDeploymentStrategy(value: object): value is DeploySvcDeploymentStrategy {
    return true;
}

export function DeploySvcDeploymentStrategyFromJSON(json: any): DeploySvcDeploymentStrategy {
    return DeploySvcDeploymentStrategyFromJSONTyped(json, false);
}

export function DeploySvcDeploymentStrategyFromJSONTyped(json: any, ignoreDiscriminator: boolean): DeploySvcDeploymentStrategy {
    if (json == null) {
        return json;
    }
    return {
        
        'maxSurge': json['maxSurge'] == null ? undefined : json['maxSurge'],
        'maxUnavailable': json['maxUnavailable'] == null ? undefined : json['maxUnavailable'],
        'type': json['type'] == null ? undefined : DeploySvcStrategyTypeFromJSON(json['type']),
    };
}

export function DeploySvcDeploymentStrategyToJSON(value?: DeploySvcDeploymentStrategy | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'maxSurge': value['maxSurge'],
        'maxUnavailable': value['maxUnavailable'],
        'type': DeploySvcStrategyTypeToJSON(value['type']),
    };
}

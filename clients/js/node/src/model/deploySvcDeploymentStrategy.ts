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

import { RequestFile } from './models';
import { DeploySvcStrategyType } from './deploySvcStrategyType';

export class DeploySvcDeploymentStrategy {
    /**
    * Max extra replicas during update
    */
    'maxSurge'?: number;
    /**
    * Max unavailable replicas during update
    */
    'maxUnavailable'?: number;
    /**
    * Deployment strategy type (RollingUpdate, Recreate, etc.)
    */
    'type'?: DeploySvcStrategyType;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "maxSurge",
            "baseName": "maxSurge",
            "type": "number"
        },
        {
            "name": "maxUnavailable",
            "baseName": "maxUnavailable",
            "type": "number"
        },
        {
            "name": "type",
            "baseName": "type",
            "type": "DeploySvcStrategyType"
        }    ];

    static getAttributeTypeMap() {
        return DeploySvcDeploymentStrategy.attributeTypeMap;
    }
}

export namespace DeploySvcDeploymentStrategy {
}
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
import { DeploySvcAutoScalingConfig } from './deploySvcAutoScalingConfig';
import { DeploySvcDeploymentStrategy } from './deploySvcDeploymentStrategy';
import { DeploySvcResourceLimits } from './deploySvcResourceLimits';
import { DeploySvcTargetRegion } from './deploySvcTargetRegion';

export class DeploySvcDeployment {
    /**
    * Optional: Auto-scaling rules
    */
    'autoScaling'?: DeploySvcAutoScalingConfig;
    /**
    * ID of the deployment (e.g., \"depl_dbOdi5eLQK\")
    */
    'id'?: string;
    /**
    * Number of container instances to run
    */
    'replicas'?: number;
    /**
    * Resource requirements for each replica
    */
    'resources'?: DeploySvcResourceLimits;
    /**
    * The User Svc slug of the service that is being deployed.
    */
    'serviceSlug': string;
    /**
    * Deployment strategy (e.g., rolling update)
    */
    'strategy'?: DeploySvcDeploymentStrategy;
    /**
    * Target deployment regions or clusters
    */
    'targetRegions'?: Array<DeploySvcTargetRegion>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "autoScaling",
            "baseName": "autoScaling",
            "type": "DeploySvcAutoScalingConfig"
        },
        {
            "name": "id",
            "baseName": "id",
            "type": "string"
        },
        {
            "name": "replicas",
            "baseName": "replicas",
            "type": "number"
        },
        {
            "name": "resources",
            "baseName": "resources",
            "type": "DeploySvcResourceLimits"
        },
        {
            "name": "serviceSlug",
            "baseName": "serviceSlug",
            "type": "string"
        },
        {
            "name": "strategy",
            "baseName": "strategy",
            "type": "DeploySvcDeploymentStrategy"
        },
        {
            "name": "targetRegions",
            "baseName": "targetRegions",
            "type": "Array<DeploySvcTargetRegion>"
        }    ];

    static getAttributeTypeMap() {
        return DeploySvcDeployment.attributeTypeMap;
    }
}

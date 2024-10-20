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
 * @interface DockerSvcContainerIsRunningResponse
 */
export interface DockerSvcContainerIsRunningResponse {
    /**
     * 
     * @type {boolean}
     * @memberof DockerSvcContainerIsRunningResponse
     */
    isRunning?: boolean;
}

/**
 * Check if a given object implements the DockerSvcContainerIsRunningResponse interface.
 */
export function instanceOfDockerSvcContainerIsRunningResponse(value: object): value is DockerSvcContainerIsRunningResponse {
    return true;
}

export function DockerSvcContainerIsRunningResponseFromJSON(json: any): DockerSvcContainerIsRunningResponse {
    return DockerSvcContainerIsRunningResponseFromJSONTyped(json, false);
}

export function DockerSvcContainerIsRunningResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DockerSvcContainerIsRunningResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isRunning': json['isRunning'] == null ? undefined : json['isRunning'],
    };
}

export function DockerSvcContainerIsRunningResponseToJSON(value?: DockerSvcContainerIsRunningResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'isRunning': value['isRunning'],
    };
}


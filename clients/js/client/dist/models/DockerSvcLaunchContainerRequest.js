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
import { DockerSvcLaunchContainerOptionsFromJSON, DockerSvcLaunchContainerOptionsToJSON, } from './DockerSvcLaunchContainerOptions';
/**
 * Check if a given object implements the DockerSvcLaunchContainerRequest interface.
 */
export function instanceOfDockerSvcLaunchContainerRequest(value) {
    if (!('image' in value) || value['image'] === undefined)
        return false;
    if (!('port' in value) || value['port'] === undefined)
        return false;
    return true;
}
export function DockerSvcLaunchContainerRequestFromJSON(json) {
    return DockerSvcLaunchContainerRequestFromJSONTyped(json, false);
}
export function DockerSvcLaunchContainerRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'hostPort': json['hostPort'] == null ? undefined : json['hostPort'],
        'image': json['image'],
        'options': json['options'] == null ? undefined : DockerSvcLaunchContainerOptionsFromJSON(json['options']),
        'port': json['port'],
    };
}
export function DockerSvcLaunchContainerRequestToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'hostPort': value['hostPort'],
        'image': value['image'],
        'options': DockerSvcLaunchContainerOptionsToJSON(value['options']),
        'port': value['port'],
    };
}

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
import { RegistrySvcGPUFromJSON, RegistrySvcGPUToJSON, } from './RegistrySvcGPU';
import { RegistrySvcResourceUsageFromJSON, RegistrySvcResourceUsageToJSON, } from './RegistrySvcResourceUsage';
/**
 * Check if a given object implements the RegistrySvcNode interface.
 */
export function instanceOfRegistrySvcNode(value) {
    return true;
}
export function RegistrySvcNodeFromJSON(json) {
    return RegistrySvcNodeFromJSONTyped(json, false);
}
export function RegistrySvcNodeFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'availabilityZone': json['availabilityZone'] == null ? undefined : json['availabilityZone'],
        'gpus': json['gpus'] == null ? undefined : (json['gpus'].map(RegistrySvcGPUFromJSON)),
        'lastHeartbeat': json['lastHeartbeat'] == null ? undefined : json['lastHeartbeat'],
        'region': json['region'] == null ? undefined : json['region'],
        'url': json['url'] == null ? undefined : json['url'],
        'usage': json['usage'] == null ? undefined : RegistrySvcResourceUsageFromJSON(json['usage']),
    };
}
export function RegistrySvcNodeToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'availabilityZone': value['availabilityZone'],
        'gpus': value['gpus'] == null ? undefined : (value['gpus'].map(RegistrySvcGPUToJSON)),
        'lastHeartbeat': value['lastHeartbeat'],
        'region': value['region'],
        'url': value['url'],
        'usage': RegistrySvcResourceUsageToJSON(value['usage']),
    };
}

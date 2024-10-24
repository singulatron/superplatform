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
import { RegistrySvcAPISpecFromJSON, RegistrySvcAPISpecToJSON, } from './RegistrySvcAPISpec';
import { RegistrySvcImageSpecFromJSON, RegistrySvcImageSpecToJSON, } from './RegistrySvcImageSpec';
import { RegistrySvcClientFromJSON, RegistrySvcClientToJSON, } from './RegistrySvcClient';
/**
 * Check if a given object implements the RegistrySvcDefinition interface.
 */
export function instanceOfRegistrySvcDefinition(value) {
    if (!('id' in value) || value['id'] === undefined)
        return false;
    if (!('image' in value) || value['image'] === undefined)
        return false;
    return true;
}
export function RegistrySvcDefinitionFromJSON(json) {
    return RegistrySvcDefinitionFromJSONTyped(json, false);
}
export function RegistrySvcDefinitionFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'apiSpecs': json['apiSpecs'] == null ? undefined : (json['apiSpecs'].map(RegistrySvcAPISpecFromJSON)),
        'clients': json['clients'] == null ? undefined : (json['clients'].map(RegistrySvcClientFromJSON)),
        'id': json['id'],
        'image': RegistrySvcImageSpecFromJSON(json['image']),
    };
}
export function RegistrySvcDefinitionToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'apiSpecs': value['apiSpecs'] == null ? undefined : (value['apiSpecs'].map(RegistrySvcAPISpecToJSON)),
        'clients': value['clients'] == null ? undefined : (value['clients'].map(RegistrySvcClientToJSON)),
        'id': value['id'],
        'image': RegistrySvcImageSpecToJSON(value['image']),
    };
}

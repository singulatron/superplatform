/* tslint:disable */
/* eslint-disable */
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
import { DatastoreConditionFromJSON, DatastoreConditionToJSON, } from './DatastoreCondition';
import { GenerictypesGenericObjectFromJSON, GenerictypesGenericObjectToJSON, } from './GenerictypesGenericObject';
/**
 * Check if a given object implements the GenerictypesUpdateRequest interface.
 */
export function instanceOfGenerictypesUpdateRequest(value) {
    return true;
}
export function GenerictypesUpdateRequestFromJSON(json) {
    return GenerictypesUpdateRequestFromJSONTyped(json, false);
}
export function GenerictypesUpdateRequestFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'conditions': json['conditions'] == null ? undefined : (json['conditions'].map(DatastoreConditionFromJSON)),
        'object': json['object'] == null ? undefined : GenerictypesGenericObjectFromJSON(json['object']),
        'table': json['table'] == null ? undefined : json['table'],
    };
}
export function GenerictypesUpdateRequestToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'conditions': value['conditions'] == null ? undefined : (value['conditions'].map(DatastoreConditionToJSON)),
        'object': GenerictypesGenericObjectToJSON(value['object']),
        'table': value['table'],
    };
}
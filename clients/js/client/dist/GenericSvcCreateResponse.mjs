import { GenericSvcGenericObjectFromJSON, GenericSvcGenericObjectToJSON } from './GenericSvcGenericObject.mjs';

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
/**
 * Check if a given object implements the GenericSvcCreateResponse interface.
 */
function instanceOfGenericSvcCreateResponse(value) {
    return true;
}
function GenericSvcCreateResponseFromJSON(json) {
    return GenericSvcCreateResponseFromJSONTyped(json);
}
function GenericSvcCreateResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'object': json['object'] == null ? undefined : GenericSvcGenericObjectFromJSON(json['object']),
    };
}
function GenericSvcCreateResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'object': GenericSvcGenericObjectToJSON(value['object']),
    };
}

export { GenericSvcCreateResponseFromJSON, GenericSvcCreateResponseFromJSONTyped, GenericSvcCreateResponseToJSON, instanceOfGenericSvcCreateResponse };
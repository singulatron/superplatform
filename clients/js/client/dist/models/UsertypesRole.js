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
 * Check if a given object implements the UsertypesRole interface.
 */
export function instanceOfUsertypesRole(value) {
    return true;
}
export function UsertypesRoleFromJSON(json) {
    return UsertypesRoleFromJSONTyped(json, false);
}
export function UsertypesRoleFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'description': json['description'] == null ? undefined : json['description'],
        'id': json['id'] == null ? undefined : json['id'],
        'name': json['name'] == null ? undefined : json['name'],
        'ownerId': json['ownerId'] == null ? undefined : json['ownerId'],
        'permissionIds': json['permissionIds'] == null ? undefined : json['permissionIds'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
export function UsertypesRoleToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'description': value['description'],
        'id': value['id'],
        'name': value['name'],
        'ownerId': value['ownerId'],
        'permissionIds': value['permissionIds'],
        'updatedAt': value['updatedAt'],
    };
}
'use strict';

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
 * Check if a given object implements the UsertypesUser interface.
 */
function instanceOfUsertypesUser(value) {
    return true;
}
function UsertypesUserFromJSON(json) {
    return UsertypesUserFromJSONTyped(json);
}
function UsertypesUserFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'createdAt': json['createdAt'] == null ? undefined : json['createdAt'],
        'deletedAt': json['deletedAt'] == null ? undefined : json['deletedAt'],
        'email': json['email'] == null ? undefined : json['email'],
        'id': json['id'] == null ? undefined : json['id'],
        'isService': json['isService'] == null ? undefined : json['isService'],
        'name': json['name'] == null ? undefined : json['name'],
        'passwordHash': json['passwordHash'] == null ? undefined : json['passwordHash'],
        'roleIds': json['roleIds'] == null ? undefined : json['roleIds'],
        'updatedAt': json['updatedAt'] == null ? undefined : json['updatedAt'],
    };
}
function UsertypesUserToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'createdAt': value['createdAt'],
        'deletedAt': value['deletedAt'],
        'email': value['email'],
        'id': value['id'],
        'isService': value['isService'],
        'name': value['name'],
        'passwordHash': value['passwordHash'],
        'roleIds': value['roleIds'],
        'updatedAt': value['updatedAt'],
    };
}

exports.UsertypesUserFromJSON = UsertypesUserFromJSON;
exports.UsertypesUserFromJSONTyped = UsertypesUserFromJSONTyped;
exports.UsertypesUserToJSON = UsertypesUserToJSON;
exports.instanceOfUsertypesUser = instanceOfUsertypesUser;
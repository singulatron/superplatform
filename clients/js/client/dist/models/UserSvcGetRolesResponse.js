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
import { UserSvcRoleFromJSON, UserSvcRoleToJSON, } from './UserSvcRole';
/**
 * Check if a given object implements the UserSvcGetRolesResponse interface.
 */
export function instanceOfUserSvcGetRolesResponse(value) {
    return true;
}
export function UserSvcGetRolesResponseFromJSON(json) {
    return UserSvcGetRolesResponseFromJSONTyped(json, false);
}
export function UserSvcGetRolesResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'roles': json['roles'] == null ? undefined : (json['roles'].map(UserSvcRoleFromJSON)),
    };
}
export function UserSvcGetRolesResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'roles': value['roles'] == null ? undefined : (value['roles'].map(UserSvcRoleToJSON)),
    };
}

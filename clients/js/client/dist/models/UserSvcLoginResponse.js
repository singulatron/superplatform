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
import { UserSvcAuthTokenFromJSON, UserSvcAuthTokenToJSON, } from './UserSvcAuthToken';
/**
 * Check if a given object implements the UserSvcLoginResponse interface.
 */
export function instanceOfUserSvcLoginResponse(value) {
    return true;
}
export function UserSvcLoginResponseFromJSON(json) {
    return UserSvcLoginResponseFromJSONTyped(json, false);
}
export function UserSvcLoginResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'token': json['token'] == null ? undefined : UserSvcAuthTokenFromJSON(json['token']),
    };
}
export function UserSvcLoginResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'token': UserSvcAuthTokenToJSON(value['token']),
    };
}

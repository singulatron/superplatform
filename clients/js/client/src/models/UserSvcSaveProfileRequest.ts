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
 * @interface UserSvcSaveProfileRequest
 */
export interface UserSvcSaveProfileRequest {
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveProfileRequest
     */
    name?: string;
    /**
     * 
     * @type {string}
     * @memberof UserSvcSaveProfileRequest
     */
    slug?: string;
}

/**
 * Check if a given object implements the UserSvcSaveProfileRequest interface.
 */
export function instanceOfUserSvcSaveProfileRequest(value: object): value is UserSvcSaveProfileRequest {
    return true;
}

export function UserSvcSaveProfileRequestFromJSON(json: any): UserSvcSaveProfileRequest {
    return UserSvcSaveProfileRequestFromJSONTyped(json, false);
}

export function UserSvcSaveProfileRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcSaveProfileRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'name': json['name'] == null ? undefined : json['name'],
        'slug': json['slug'] == null ? undefined : json['slug'],
    };
}

export function UserSvcSaveProfileRequestToJSON(value?: UserSvcSaveProfileRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'name': value['name'],
        'slug': value['slug'],
    };
}


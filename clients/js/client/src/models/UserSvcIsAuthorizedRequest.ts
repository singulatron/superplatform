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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UserSvcIsAuthorizedRequest
 */
export interface UserSvcIsAuthorizedRequest {
    /**
     * 
     * @type {Array<string>}
     * @memberof UserSvcIsAuthorizedRequest
     */
    emailsGranted?: Array<string>;
}

/**
 * Check if a given object implements the UserSvcIsAuthorizedRequest interface.
 */
export function instanceOfUserSvcIsAuthorizedRequest(value: object): value is UserSvcIsAuthorizedRequest {
    return true;
}

export function UserSvcIsAuthorizedRequestFromJSON(json: any): UserSvcIsAuthorizedRequest {
    return UserSvcIsAuthorizedRequestFromJSONTyped(json, false);
}

export function UserSvcIsAuthorizedRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcIsAuthorizedRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'emailsGranted': json['emailsGranted'] == null ? undefined : json['emailsGranted'],
    };
}

export function UserSvcIsAuthorizedRequestToJSON(value?: UserSvcIsAuthorizedRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'emailsGranted': value['emailsGranted'],
    };
}

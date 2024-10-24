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
/**
 *
 * @export
 * @interface UserSvcAddUserToOrganizationRequest
 */
export interface UserSvcAddUserToOrganizationRequest {
    /**
     *
     * @type {string}
     * @memberof UserSvcAddUserToOrganizationRequest
     */
    userId?: string;
}
/**
 * Check if a given object implements the UserSvcAddUserToOrganizationRequest interface.
 */
export declare function instanceOfUserSvcAddUserToOrganizationRequest(value: object): value is UserSvcAddUserToOrganizationRequest;
export declare function UserSvcAddUserToOrganizationRequestFromJSON(json: any): UserSvcAddUserToOrganizationRequest;
export declare function UserSvcAddUserToOrganizationRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserSvcAddUserToOrganizationRequest;
export declare function UserSvcAddUserToOrganizationRequestToJSON(value?: UserSvcAddUserToOrganizationRequest | null): any;

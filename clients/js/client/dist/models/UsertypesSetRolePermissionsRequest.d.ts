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
 *
 * @export
 * @interface UsertypesSetRolePermissionsRequest
 */
export interface UsertypesSetRolePermissionsRequest {
    /**
     *
     * @type {Array<string>}
     * @memberof UsertypesSetRolePermissionsRequest
     */
    permissionIds?: Array<string>;
}
/**
 * Check if a given object implements the UsertypesSetRolePermissionsRequest interface.
 */
export declare function instanceOfUsertypesSetRolePermissionsRequest(value: object): value is UsertypesSetRolePermissionsRequest;
export declare function UsertypesSetRolePermissionsRequestFromJSON(json: any): UsertypesSetRolePermissionsRequest;
export declare function UsertypesSetRolePermissionsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UsertypesSetRolePermissionsRequest;
export declare function UsertypesSetRolePermissionsRequestToJSON(value?: UsertypesSetRolePermissionsRequest | null): any;
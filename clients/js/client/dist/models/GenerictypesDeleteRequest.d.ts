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
import type { DatastoreCondition } from './DatastoreCondition';
/**
 *
 * @export
 * @interface GenerictypesDeleteRequest
 */
export interface GenerictypesDeleteRequest {
    /**
     *
     * @type {Array<DatastoreCondition>}
     * @memberof GenerictypesDeleteRequest
     */
    conditions?: Array<DatastoreCondition>;
    /**
     *
     * @type {string}
     * @memberof GenerictypesDeleteRequest
     */
    table?: string;
}
/**
 * Check if a given object implements the GenerictypesDeleteRequest interface.
 */
export declare function instanceOfGenerictypesDeleteRequest(value: object): value is GenerictypesDeleteRequest;
export declare function GenerictypesDeleteRequestFromJSON(json: any): GenerictypesDeleteRequest;
export declare function GenerictypesDeleteRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): GenerictypesDeleteRequest;
export declare function GenerictypesDeleteRequestToJSON(value?: GenerictypesDeleteRequest | null): any;
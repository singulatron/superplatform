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
 * @interface ChatAsset
 */
export interface ChatAsset {
    /**
     * Content is the base64 encoded binary file direcly embedded in the asset itself
     * @type {string}
     * @memberof ChatAsset
     */
    content?: string;
    /**
     *
     * @type {string}
     * @memberof ChatAsset
     */
    createdAt?: string;
    /**
     *
     * @type {string}
     * @memberof ChatAsset
     */
    description?: string;
    /**
     *
     * @type {string}
     * @memberof ChatAsset
     */
    id?: string;
    /**
     *
     * @type {string}
     * @memberof ChatAsset
     */
    type?: string;
    /**
     *
     * @type {string}
     * @memberof ChatAsset
     */
    updatedAt?: string;
    /**
     * Url of the asset where
     * @type {string}
     * @memberof ChatAsset
     */
    url?: string;
}
/**
 * Check if a given object implements the ChatAsset interface.
 */
export declare function instanceOfChatAsset(value: object): value is ChatAsset;
export declare function ChatAssetFromJSON(json: any): ChatAsset;
export declare function ChatAssetFromJSONTyped(json: any, ignoreDiscriminator: boolean): ChatAsset;
export declare function ChatAssetToJSON(value?: ChatAsset | null): any;
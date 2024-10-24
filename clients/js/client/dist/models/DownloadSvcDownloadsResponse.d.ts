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
import type { DownloadSvcDownloadDetails } from './DownloadSvcDownloadDetails';
/**
 *
 * @export
 * @interface DownloadSvcDownloadsResponse
 */
export interface DownloadSvcDownloadsResponse {
    /**
     *
     * @type {Array<DownloadSvcDownloadDetails>}
     * @memberof DownloadSvcDownloadsResponse
     */
    downloads?: Array<DownloadSvcDownloadDetails>;
}
/**
 * Check if a given object implements the DownloadSvcDownloadsResponse interface.
 */
export declare function instanceOfDownloadSvcDownloadsResponse(value: object): value is DownloadSvcDownloadsResponse;
export declare function DownloadSvcDownloadsResponseFromJSON(json: any): DownloadSvcDownloadsResponse;
export declare function DownloadSvcDownloadsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): DownloadSvcDownloadsResponse;
export declare function DownloadSvcDownloadsResponseToJSON(value?: DownloadSvcDownloadsResponse | null): any;

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
import { DownloadSvcDownloadDetailsFromJSON, DownloadSvcDownloadDetailsToJSON, } from './DownloadSvcDownloadDetails';
/**
 * Check if a given object implements the DownloadSvcGetDownloadResponse interface.
 */
export function instanceOfDownloadSvcGetDownloadResponse(value) {
    return true;
}
export function DownloadSvcGetDownloadResponseFromJSON(json) {
    return DownloadSvcGetDownloadResponseFromJSONTyped(json, false);
}
export function DownloadSvcGetDownloadResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'download': json['download'] == null ? undefined : DownloadSvcDownloadDetailsFromJSON(json['download']),
        '_exists': json['exists'] == null ? undefined : json['exists'],
    };
}
export function DownloadSvcGetDownloadResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'download': DownloadSvcDownloadDetailsToJSON(value['download']),
        'exists': value['_exists'],
    };
}

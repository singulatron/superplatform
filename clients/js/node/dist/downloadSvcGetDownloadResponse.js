'use strict';

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
class DownloadSvcGetDownloadResponse {
    static getAttributeTypeMap() {
        return DownloadSvcGetDownloadResponse.attributeTypeMap;
    }
}
DownloadSvcGetDownloadResponse.discriminator = undefined;
DownloadSvcGetDownloadResponse.attributeTypeMap = [
    {
        "name": "download",
        "baseName": "download",
        "type": "DownloadSvcDownloadDetails"
    },
    {
        "name": "exists",
        "baseName": "exists",
        "type": "boolean"
    }
];

exports.DownloadSvcGetDownloadResponse = DownloadSvcGetDownloadResponse;

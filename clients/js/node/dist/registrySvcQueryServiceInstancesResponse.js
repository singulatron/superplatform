'use strict';

/**
 * Superplatform
 * AI management and development platform.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
class RegistrySvcQueryServiceInstancesResponse {
    static getAttributeTypeMap() {
        return RegistrySvcQueryServiceInstancesResponse.attributeTypeMap;
    }
}
RegistrySvcQueryServiceInstancesResponse.discriminator = undefined;
RegistrySvcQueryServiceInstancesResponse.attributeTypeMap = [
    {
        "name": "instances",
        "baseName": "instances",
        "type": "Array<RegistrySvcServiceInstance>"
    }
];

exports.RegistrySvcQueryServiceInstancesResponse = RegistrySvcQueryServiceInstancesResponse;

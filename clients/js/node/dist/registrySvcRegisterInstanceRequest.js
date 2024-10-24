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
class RegistrySvcRegisterInstanceRequest {
    static getAttributeTypeMap() {
        return RegistrySvcRegisterInstanceRequest.attributeTypeMap;
    }
}
RegistrySvcRegisterInstanceRequest.discriminator = undefined;
RegistrySvcRegisterInstanceRequest.attributeTypeMap = [
    {
        "name": "definitionId",
        "baseName": "definitionId",
        "type": "string"
    },
    {
        "name": "host",
        "baseName": "host",
        "type": "string"
    },
    {
        "name": "ip",
        "baseName": "ip",
        "type": "string"
    },
    {
        "name": "path",
        "baseName": "path",
        "type": "string"
    },
    {
        "name": "port",
        "baseName": "port",
        "type": "number"
    },
    {
        "name": "scheme",
        "baseName": "scheme",
        "type": "string"
    },
    {
        "name": "url",
        "baseName": "url",
        "type": "string"
    }
];

exports.RegistrySvcRegisterInstanceRequest = RegistrySvcRegisterInstanceRequest;

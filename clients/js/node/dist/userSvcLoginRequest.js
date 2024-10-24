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
class UserSvcLoginRequest {
    static getAttributeTypeMap() {
        return UserSvcLoginRequest.attributeTypeMap;
    }
}
UserSvcLoginRequest.discriminator = undefined;
UserSvcLoginRequest.attributeTypeMap = [
    {
        "name": "contact",
        "baseName": "contact",
        "type": "string"
    },
    {
        "name": "password",
        "baseName": "password",
        "type": "string"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    }
];

exports.UserSvcLoginRequest = UserSvcLoginRequest;

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
class RegistrySvcListNodesResponse {
    static getAttributeTypeMap() {
        return RegistrySvcListNodesResponse.attributeTypeMap;
    }
}
RegistrySvcListNodesResponse.discriminator = undefined;
RegistrySvcListNodesResponse.attributeTypeMap = [
    {
        "name": "nodes",
        "baseName": "nodes",
        "type": "Array<RegistrySvcNode>"
    }
];

export { RegistrySvcListNodesResponse };

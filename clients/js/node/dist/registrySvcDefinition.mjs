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
class RegistrySvcDefinition {
    static getAttributeTypeMap() {
        return RegistrySvcDefinition.attributeTypeMap;
    }
}
RegistrySvcDefinition.discriminator = undefined;
RegistrySvcDefinition.attributeTypeMap = [
    {
        "name": "apiSpecs",
        "baseName": "apiSpecs",
        "type": "Array<RegistrySvcAPISpec>"
    },
    {
        "name": "clients",
        "baseName": "clients",
        "type": "Array<RegistrySvcClient>"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "image",
        "baseName": "image",
        "type": "RegistrySvcImageSpec"
    }
];

export { RegistrySvcDefinition };

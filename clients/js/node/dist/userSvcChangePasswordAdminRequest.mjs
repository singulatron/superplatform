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
class UserSvcChangePasswordAdminRequest {
    static getAttributeTypeMap() {
        return UserSvcChangePasswordAdminRequest.attributeTypeMap;
    }
}
UserSvcChangePasswordAdminRequest.discriminator = undefined;
UserSvcChangePasswordAdminRequest.attributeTypeMap = [
    {
        "name": "newPassword",
        "baseName": "newPassword",
        "type": "string"
    },
    {
        "name": "slug",
        "baseName": "slug",
        "type": "string"
    }
];

export { UserSvcChangePasswordAdminRequest };
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
class ModelSvcModel {
    static getAttributeTypeMap() {
        return ModelSvcModel.attributeTypeMap;
    }
}
ModelSvcModel.discriminator = undefined;
ModelSvcModel.attributeTypeMap = [
    {
        "name": "assets",
        "baseName": "assets",
        "type": "{ [key: string]: string; }"
    },
    {
        "name": "bits",
        "baseName": "bits",
        "type": "number"
    },
    {
        "name": "description",
        "baseName": "description",
        "type": "string"
    },
    {
        "name": "extension",
        "baseName": "extension",
        "type": "string"
    },
    {
        "name": "flavour",
        "baseName": "flavour",
        "type": "string"
    },
    {
        "name": "fullName",
        "baseName": "full_name",
        "type": "string"
    },
    {
        "name": "id",
        "baseName": "id",
        "type": "string"
    },
    {
        "name": "maxBits",
        "baseName": "max_bits",
        "type": "number"
    },
    {
        "name": "maxRam",
        "baseName": "max_ram",
        "type": "number"
    },
    {
        "name": "mirrors",
        "baseName": "mirrors",
        "type": "Array<string>"
    },
    {
        "name": "name",
        "baseName": "name",
        "type": "string"
    },
    {
        "name": "parameters",
        "baseName": "parameters",
        "type": "string"
    },
    {
        "name": "platformId",
        "baseName": "platformId",
        "type": "string"
    },
    {
        "name": "promptTemplate",
        "baseName": "prompt_template",
        "type": "string"
    },
    {
        "name": "quality",
        "baseName": "quality",
        "type": "string"
    },
    {
        "name": "quantComment",
        "baseName": "quant_comment",
        "type": "string"
    },
    {
        "name": "size",
        "baseName": "size",
        "type": "number"
    },
    {
        "name": "tags",
        "baseName": "tags",
        "type": "Array<string>"
    },
    {
        "name": "uncensored",
        "baseName": "uncensored",
        "type": "boolean"
    },
    {
        "name": "version",
        "baseName": "version",
        "type": "string"
    }
];

export { ModelSvcModel };
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

import { RequestFile } from './models';
import { ModelSvcModel } from './modelSvcModel';
import { ModelSvcPlatform } from './modelSvcPlatform';

export class ModelSvcGetModelResponse {
    'exists'?: boolean;
    'model'?: ModelSvcModel;
    'platform'?: ModelSvcPlatform;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "exists",
            "baseName": "exists",
            "type": "boolean"
        },
        {
            "name": "model",
            "baseName": "model",
            "type": "ModelSvcModel"
        },
        {
            "name": "platform",
            "baseName": "platform",
            "type": "ModelSvcPlatform"
        }    ];

    static getAttributeTypeMap() {
        return ModelSvcGetModelResponse.attributeTypeMap;
    }
}

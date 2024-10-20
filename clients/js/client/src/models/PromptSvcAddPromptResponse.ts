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

import { mapValues } from '../runtime';
import type { PromptSvcPrompt } from './PromptSvcPrompt';
import {
    PromptSvcPromptFromJSON,
    PromptSvcPromptFromJSONTyped,
    PromptSvcPromptToJSON,
} from './PromptSvcPrompt';

/**
 * 
 * @export
 * @interface PromptSvcAddPromptResponse
 */
export interface PromptSvcAddPromptResponse {
    /**
     * 
     * @type {string}
     * @memberof PromptSvcAddPromptResponse
     */
    answer?: string;
    /**
     * 
     * @type {PromptSvcPrompt}
     * @memberof PromptSvcAddPromptResponse
     */
    prompt?: PromptSvcPrompt;
}

/**
 * Check if a given object implements the PromptSvcAddPromptResponse interface.
 */
export function instanceOfPromptSvcAddPromptResponse(value: object): value is PromptSvcAddPromptResponse {
    return true;
}

export function PromptSvcAddPromptResponseFromJSON(json: any): PromptSvcAddPromptResponse {
    return PromptSvcAddPromptResponseFromJSONTyped(json, false);
}

export function PromptSvcAddPromptResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PromptSvcAddPromptResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'answer': json['answer'] == null ? undefined : json['answer'],
        'prompt': json['prompt'] == null ? undefined : PromptSvcPromptFromJSON(json['prompt']),
    };
}

export function PromptSvcAddPromptResponseToJSON(value?: PromptSvcAddPromptResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'answer': value['answer'],
        'prompt': PromptSvcPromptToJSON(value['prompt']),
    };
}


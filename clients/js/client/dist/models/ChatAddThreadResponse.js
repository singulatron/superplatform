/* tslint:disable */
/* eslint-disable */
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
import { ChatThreadFromJSON, ChatThreadToJSON, } from './ChatThread';
/**
 * Check if a given object implements the ChatAddThreadResponse interface.
 */
export function instanceOfChatAddThreadResponse(value) {
    return true;
}
export function ChatAddThreadResponseFromJSON(json) {
    return ChatAddThreadResponseFromJSONTyped(json, false);
}
export function ChatAddThreadResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'thread': json['thread'] == null ? undefined : ChatThreadFromJSON(json['thread']),
    };
}
export function ChatAddThreadResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'thread': ChatThreadToJSON(value['thread']),
    };
}
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
import { ChatSvcThreadFromJSON, ChatSvcThreadToJSON, } from './ChatSvcThread';
/**
 * Check if a given object implements the ChatSvcAddThreadResponse interface.
 */
export function instanceOfChatSvcAddThreadResponse(value) {
    return true;
}
export function ChatSvcAddThreadResponseFromJSON(json) {
    return ChatSvcAddThreadResponseFromJSONTyped(json, false);
}
export function ChatSvcAddThreadResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'thread': json['thread'] == null ? undefined : ChatSvcThreadFromJSON(json['thread']),
    };
}
export function ChatSvcAddThreadResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'thread': ChatSvcThreadToJSON(value['thread']),
    };
}

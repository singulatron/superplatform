'use strict';

/* tslint:disable */
/* eslint-disable */
/**
 * Singulatron
 * AI management and development platform.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Check if a given object implements the DockerSvcLaunchContainerOptions interface.
 */
function instanceOfDockerSvcLaunchContainerOptions(value) {
    return true;
}
function DockerSvcLaunchContainerOptionsFromJSON(json) {
    return DockerSvcLaunchContainerOptionsFromJSONTyped(json);
}
function DockerSvcLaunchContainerOptionsFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'assets': json['assets'] == null ? undefined : json['assets'],
        'envs': json['envs'] == null ? undefined : json['envs'],
        'gpuEnabled': json['gpuEnabled'] == null ? undefined : json['gpuEnabled'],
        'hash': json['hash'] == null ? undefined : json['hash'],
        'labels': json['labels'] == null ? undefined : json['labels'],
        'name': json['name'] == null ? undefined : json['name'],
        'persistentPaths': json['persistentPaths'] == null ? undefined : json['persistentPaths'],
    };
}
function DockerSvcLaunchContainerOptionsToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'assets': value['assets'],
        'envs': value['envs'],
        'gpuEnabled': value['gpuEnabled'],
        'hash': value['hash'],
        'labels': value['labels'],
        'name': value['name'],
        'persistentPaths': value['persistentPaths'],
    };
}

exports.DockerSvcLaunchContainerOptionsFromJSON = DockerSvcLaunchContainerOptionsFromJSON;
exports.DockerSvcLaunchContainerOptionsFromJSONTyped = DockerSvcLaunchContainerOptionsFromJSONTyped;
exports.DockerSvcLaunchContainerOptionsToJSON = DockerSvcLaunchContainerOptionsToJSON;
exports.instanceOfDockerSvcLaunchContainerOptions = instanceOfDockerSvcLaunchContainerOptions;
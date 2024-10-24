import { DeploySvcDeploymentFromJSON, DeploySvcDeploymentToJSON } from './DeploySvcDeployment.mjs';
import './DeploySvcDeploymentStrategy.mjs';
import './DeploySvcStrategyType.mjs';
import './DeploySvcAutoScalingConfig.mjs';
import './DeploySvcDeploymentStatus.mjs';
import './DeploySvcTargetRegion.mjs';
import './DeploySvcResourceLimits.mjs';

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
/**
 * Check if a given object implements the DeploySvcListDeploymentsResponse interface.
 */
function instanceOfDeploySvcListDeploymentsResponse(value) {
    return true;
}
function DeploySvcListDeploymentsResponseFromJSON(json) {
    return DeploySvcListDeploymentsResponseFromJSONTyped(json);
}
function DeploySvcListDeploymentsResponseFromJSONTyped(json, ignoreDiscriminator) {
    if (json == null) {
        return json;
    }
    return {
        'deployments': json['deployments'] == null ? undefined : (json['deployments'].map(DeploySvcDeploymentFromJSON)),
    };
}
function DeploySvcListDeploymentsResponseToJSON(value) {
    if (value == null) {
        return value;
    }
    return {
        'deployments': value['deployments'] == null ? undefined : (value['deployments'].map(DeploySvcDeploymentToJSON)),
    };
}

export { DeploySvcListDeploymentsResponseFromJSON, DeploySvcListDeploymentsResponseFromJSONTyped, DeploySvcListDeploymentsResponseToJSON, instanceOfDeploySvcListDeploymentsResponse };

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
import type { DockerSvcLaunchContainerOptions } from './DockerSvcLaunchContainerOptions';
/**
 *
 * @export
 * @interface DockerSvcLaunchContainerRequest
 */
export interface DockerSvcLaunchContainerRequest {
    /**
     * HostPort is the port on the host machine that will be mapped to the container's port
     * @type {number}
     * @memberof DockerSvcLaunchContainerRequest
     */
    hostPort?: number;
    /**
     * Image is the Docker image to use for the container
     * @type {string}
     * @memberof DockerSvcLaunchContainerRequest
     */
    image: string;
    /**
     * Options provides additional options for launching the container
     * @type {DockerSvcLaunchContainerOptions}
     * @memberof DockerSvcLaunchContainerRequest
     */
    options?: DockerSvcLaunchContainerOptions;
    /**
     * Port is the port number that the container will expose
     * @type {number}
     * @memberof DockerSvcLaunchContainerRequest
     */
    port: number;
}
/**
 * Check if a given object implements the DockerSvcLaunchContainerRequest interface.
 */
export declare function instanceOfDockerSvcLaunchContainerRequest(value: object): value is DockerSvcLaunchContainerRequest;
export declare function DockerSvcLaunchContainerRequestFromJSON(json: any): DockerSvcLaunchContainerRequest;
export declare function DockerSvcLaunchContainerRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): DockerSvcLaunchContainerRequest;
export declare function DockerSvcLaunchContainerRequestToJSON(value?: DockerSvcLaunchContainerRequest | null): any;

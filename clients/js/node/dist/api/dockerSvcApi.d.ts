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
import http from 'http';
import { DockerSvcContainerIsRunningResponse } from '../model/dockerSvcContainerIsRunningResponse';
import { DockerSvcGetContainerSummaryResponse } from '../model/dockerSvcGetContainerSummaryResponse';
import { DockerSvcGetDockerHostResponse } from '../model/dockerSvcGetDockerHostResponse';
import { DockerSvcGetInfoResponse } from '../model/dockerSvcGetInfoResponse';
import { DockerSvcLaunchContainerRequest } from '../model/dockerSvcLaunchContainerRequest';
import { DockerSvcLaunchContainerResponse } from '../model/dockerSvcLaunchContainerResponse';
import { Authentication, Interceptor } from '../model/models';
import { ApiKeyAuth } from '../model/models';
export declare enum DockerSvcApiApiKeys {
    BearerAuth = 0
}
export declare class DockerSvcApi {
    protected _basePath: string;
    protected _defaultHeaders: any;
    protected _useQuerystring: boolean;
    protected authentications: {
        default: Authentication;
        BearerAuth: ApiKeyAuth;
    };
    protected interceptors: Interceptor[];
    constructor(basePath?: string);
    set useQuerystring(value: boolean);
    set basePath(basePath: string);
    set defaultHeaders(defaultHeaders: any);
    get defaultHeaders(): any;
    get basePath(): string;
    setDefaultAuthentication(auth: Authentication): void;
    setApiKey(key: DockerSvcApiApiKeys, value: string): void;
    addInterceptor(interceptor: Interceptor): void;
    /**
     * Get a summary of the Docker container identified by the hash, limited to a specified number of lines
     * @summary Get Container Summary
     * @param hash Container Hash
     * @param numberOfLines Number of Lines
     */
    getContainerSummary(hash: string, numberOfLines: number, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DockerSvcGetContainerSummaryResponse;
    }>;
    /**
     * Retrieve information about the Docker host
     * @summary Get Docker Host
     */
    getHost(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DockerSvcGetDockerHostResponse;
    }>;
    /**
     * Retrieve detailed information about the Docker service
     * @summary Get Docker Service Information
     */
    getInfo(options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DockerSvcGetInfoResponse;
    }>;
    /**
     * Check if a Docker container identified by the hash is running
     * @summary Check If a Container Is Running
     * @param hash Container Hash
     */
    isRunning(hash: string, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DockerSvcContainerIsRunningResponse;
    }>;
    /**
     * Launches a Docker container with the specified parameters.  Requires the `docker-svc:docker:create` permission.
     * @summary Launch a Docker Container
     * @param request Launch Container Request
     */
    launchContainer(request: DockerSvcLaunchContainerRequest, options?: {
        headers: {
            [name: string]: string;
        };
    }): Promise<{
        response: http.IncomingMessage;
        body: DockerSvcLaunchContainerResponse;
    }>;
}
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


import * as runtime from '../runtime';
import type {
  DockerSvcContainerIsRunningResponse,
  DockerSvcErrorResponse,
  DockerSvcGetContainerSummaryResponse,
  DockerSvcGetDockerHostResponse,
  DockerSvcGetInfoResponse,
  DockerSvcLaunchContainerRequest,
  DockerSvcLaunchContainerResponse,
} from '../models/index';
import {
    DockerSvcContainerIsRunningResponseFromJSON,
    DockerSvcContainerIsRunningResponseToJSON,
    DockerSvcErrorResponseFromJSON,
    DockerSvcErrorResponseToJSON,
    DockerSvcGetContainerSummaryResponseFromJSON,
    DockerSvcGetContainerSummaryResponseToJSON,
    DockerSvcGetDockerHostResponseFromJSON,
    DockerSvcGetDockerHostResponseToJSON,
    DockerSvcGetInfoResponseFromJSON,
    DockerSvcGetInfoResponseToJSON,
    DockerSvcLaunchContainerRequestFromJSON,
    DockerSvcLaunchContainerRequestToJSON,
    DockerSvcLaunchContainerResponseFromJSON,
    DockerSvcLaunchContainerResponseToJSON,
} from '../models/index';

export interface GetContainerSummaryRequest {
    hash: string;
    numberOfLines: number;
}

export interface IsRunningRequest {
    hash: string;
}

export interface LaunchContainerRequest {
    request: DockerSvcLaunchContainerRequest;
}

/**
 * 
 */
export class DockerSvcApi extends runtime.BaseAPI {

    /**
     * Get a summary of the Docker container identified by the hash, limited to a specified number of lines
     * Get Container Summary
     */
    async getContainerSummaryRaw(requestParameters: GetContainerSummaryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DockerSvcGetContainerSummaryResponse>> {
        if (requestParameters['hash'] == null) {
            throw new runtime.RequiredError(
                'hash',
                'Required parameter "hash" was null or undefined when calling getContainerSummary().'
            );
        }

        if (requestParameters['numberOfLines'] == null) {
            throw new runtime.RequiredError(
                'numberOfLines',
                'Required parameter "numberOfLines" was null or undefined when calling getContainerSummary().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/docker-svc/container/{hash}/summary/{numberOfLines}`.replace(`{${"hash"}}`, encodeURIComponent(String(requestParameters['hash']))).replace(`{${"numberOfLines"}}`, encodeURIComponent(String(requestParameters['numberOfLines']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DockerSvcGetContainerSummaryResponseFromJSON(jsonValue));
    }

    /**
     * Get a summary of the Docker container identified by the hash, limited to a specified number of lines
     * Get Container Summary
     */
    async getContainerSummary(requestParameters: GetContainerSummaryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DockerSvcGetContainerSummaryResponse> {
        const response = await this.getContainerSummaryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve information about the Docker host
     * Get Docker Host
     */
    async getHostRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DockerSvcGetDockerHostResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/docker-svc/host`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DockerSvcGetDockerHostResponseFromJSON(jsonValue));
    }

    /**
     * Retrieve information about the Docker host
     * Get Docker Host
     */
    async getHost(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DockerSvcGetDockerHostResponse> {
        const response = await this.getHostRaw(initOverrides);
        return await response.value();
    }

    /**
     * Retrieve detailed information about the Docker service
     * Get Docker Service Information
     */
    async getInfoRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DockerSvcGetInfoResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/docker-svc/info`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DockerSvcGetInfoResponseFromJSON(jsonValue));
    }

    /**
     * Retrieve detailed information about the Docker service
     * Get Docker Service Information
     */
    async getInfo(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DockerSvcGetInfoResponse> {
        const response = await this.getInfoRaw(initOverrides);
        return await response.value();
    }

    /**
     * Check if a Docker container identified by the hash is running
     * Check If a Container Is Running
     */
    async isRunningRaw(requestParameters: IsRunningRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DockerSvcContainerIsRunningResponse>> {
        if (requestParameters['hash'] == null) {
            throw new runtime.RequiredError(
                'hash',
                'Required parameter "hash" was null or undefined when calling isRunning().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/docker-svc/container/{hash}/is-running`.replace(`{${"hash"}}`, encodeURIComponent(String(requestParameters['hash']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DockerSvcContainerIsRunningResponseFromJSON(jsonValue));
    }

    /**
     * Check if a Docker container identified by the hash is running
     * Check If a Container Is Running
     */
    async isRunning(requestParameters: IsRunningRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DockerSvcContainerIsRunningResponse> {
        const response = await this.isRunningRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Launches a Docker container with the specified parameters.  Requires the `docker-svc:docker:create` permission.
     * Launch a Docker Container
     */
    async launchContainerRaw(requestParameters: LaunchContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<DockerSvcLaunchContainerResponse>> {
        if (requestParameters['request'] == null) {
            throw new runtime.RequiredError(
                'request',
                'Required parameter "request" was null or undefined when calling launchContainer().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = await this.configuration.apiKey("Authorization"); // BearerAuth authentication
        }

        const response = await this.request({
            path: `/docker-svc/container`,
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: DockerSvcLaunchContainerRequestToJSON(requestParameters['request']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => DockerSvcLaunchContainerResponseFromJSON(jsonValue));
    }

    /**
     * Launches a Docker container with the specified parameters.  Requires the `docker-svc:docker:create` permission.
     * Launch a Docker Container
     */
    async launchContainer(requestParameters: LaunchContainerRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<DockerSvcLaunchContainerResponse> {
        const response = await this.launchContainerRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
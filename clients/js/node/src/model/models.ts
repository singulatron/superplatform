import localVarRequest from 'request';

export * from './chatSvcAddMessageRequest';
export * from './chatSvcAddThreadRequest';
export * from './chatSvcAddThreadResponse';
export * from './chatSvcAsset';
export * from './chatSvcGetMessagesResponse';
export * from './chatSvcGetThreadResponse';
export * from './chatSvcGetThreadsResponse';
export * from './chatSvcMessage';
export * from './chatSvcThread';
export * from './chatSvcUpdateThreadRequest';
export * from './configSvcAppServiceConfig';
export * from './configSvcConfig';
export * from './configSvcDownloadServiceConfig';
export * from './configSvcGetConfigResponse';
export * from './configSvcModelServiceConfig';
export * from './configSvcSaveConfigRequest';
export * from './datastoreCondition';
export * from './datastoreContainsCondition';
export * from './datastoreEqualCondition';
export * from './datastoreFieldSelector';
export * from './datastoreOrderBy';
export * from './datastoreQuery';
export * from './datastoreStartsWithCondition';
export * from './dockerSvcContainerIsRunningResponse';
export * from './dockerSvcDockerInfo';
export * from './dockerSvcErrorResponse';
export * from './dockerSvcGetContainerSummaryResponse';
export * from './dockerSvcGetDockerHostResponse';
export * from './dockerSvcGetInfoResponse';
export * from './dockerSvcLaunchContainerRequest';
export * from './dockerSvcLaunchContainerResponse';
export * from './dockerSvcLaunchInfo';
export * from './dockerSvcLaunchOptions';
export * from './downloadSvcDownloadDetails';
export * from './downloadSvcDownloadRequest';
export * from './downloadSvcDownloadsResponse';
export * from './downloadSvcErrorResponse';
export * from './downloadSvcGetDownloadResponse';
export * from './firehoseSvcErrorResponse';
export * from './firehoseSvcEvent';
export * from './firehoseSvcPublishRequest';
export * from './genericSvcCreateObjectRequest';
export * from './genericSvcCreateObjectResponse';
export * from './genericSvcDeleteObjectRequest';
export * from './genericSvcErrorResponse';
export * from './genericSvcGenericObject';
export * from './genericSvcGenericObjectCreateFields';
export * from './genericSvcQueryRequest';
export * from './genericSvcQueryResponse';
export * from './genericSvcUpdateObjectRequest';
export * from './genericSvcUpsertObjectRequest';
export * from './genericSvcUpsertObjectResponse';
export * from './modelSvcArchitectures';
export * from './modelSvcContainer';
export * from './modelSvcErrorResponse';
export * from './modelSvcGetModelResponse';
export * from './modelSvcListResponse';
export * from './modelSvcModel';
export * from './modelSvcModelStatus';
export * from './modelSvcPlatform';
export * from './modelSvcStatusResponse';
export * from './promptSvcAddPromptRequest';
export * from './promptSvcAddPromptResponse';
export * from './promptSvcErrorResponse';
export * from './promptSvcListPromptsRequest';
export * from './promptSvcListPromptsResponse';
export * from './promptSvcPrompt';
export * from './promptSvcPromptStatus';
export * from './promptSvcRemovePromptRequest';
export * from './userSvcAuthToken';
export * from './userSvcChangePasswordAdminRequest';
export * from './userSvcChangePasswordRequest';
export * from './userSvcCreateRoleRequest';
export * from './userSvcCreateRoleResponse';
export * from './userSvcCreateUserRequest';
export * from './userSvcErrorResponse';
export * from './userSvcGetPermissionsResponse';
export * from './userSvcGetPublicKeyResponse';
export * from './userSvcGetRolesResponse';
export * from './userSvcGetUsersRequest';
export * from './userSvcGetUsersResponse';
export * from './userSvcIsAuthorizedRequest';
export * from './userSvcIsAuthorizedResponse';
export * from './userSvcLoginRequest';
export * from './userSvcLoginResponse';
export * from './userSvcPermission';
export * from './userSvcReadUserByTokenRequest';
export * from './userSvcReadUserByTokenResponse';
export * from './userSvcRegisterRequest';
export * from './userSvcRole';
export * from './userSvcSaveProfileRequest';
export * from './userSvcSetRolePermissionsRequest';
export * from './userSvcUpserPermissionRequest';
export * from './userSvcUser';

import * as fs from 'fs';

export interface RequestDetailedFile {
    value: Buffer;
    options?: {
        filename?: string;
        contentType?: string;
    }
}

export type RequestFile = string | Buffer | fs.ReadStream | RequestDetailedFile;


import { ChatSvcAddMessageRequest } from './chatSvcAddMessageRequest';
import { ChatSvcAddThreadRequest } from './chatSvcAddThreadRequest';
import { ChatSvcAddThreadResponse } from './chatSvcAddThreadResponse';
import { ChatSvcAsset } from './chatSvcAsset';
import { ChatSvcGetMessagesResponse } from './chatSvcGetMessagesResponse';
import { ChatSvcGetThreadResponse } from './chatSvcGetThreadResponse';
import { ChatSvcGetThreadsResponse } from './chatSvcGetThreadsResponse';
import { ChatSvcMessage } from './chatSvcMessage';
import { ChatSvcThread } from './chatSvcThread';
import { ChatSvcUpdateThreadRequest } from './chatSvcUpdateThreadRequest';
import { ConfigSvcAppServiceConfig } from './configSvcAppServiceConfig';
import { ConfigSvcConfig } from './configSvcConfig';
import { ConfigSvcDownloadServiceConfig } from './configSvcDownloadServiceConfig';
import { ConfigSvcGetConfigResponse } from './configSvcGetConfigResponse';
import { ConfigSvcModelServiceConfig } from './configSvcModelServiceConfig';
import { ConfigSvcSaveConfigRequest } from './configSvcSaveConfigRequest';
import { DatastoreCondition } from './datastoreCondition';
import { DatastoreContainsCondition } from './datastoreContainsCondition';
import { DatastoreEqualCondition } from './datastoreEqualCondition';
import { DatastoreFieldSelector } from './datastoreFieldSelector';
import { DatastoreOrderBy } from './datastoreOrderBy';
import { DatastoreQuery } from './datastoreQuery';
import { DatastoreStartsWithCondition } from './datastoreStartsWithCondition';
import { DockerSvcContainerIsRunningResponse } from './dockerSvcContainerIsRunningResponse';
import { DockerSvcDockerInfo } from './dockerSvcDockerInfo';
import { DockerSvcErrorResponse } from './dockerSvcErrorResponse';
import { DockerSvcGetContainerSummaryResponse } from './dockerSvcGetContainerSummaryResponse';
import { DockerSvcGetDockerHostResponse } from './dockerSvcGetDockerHostResponse';
import { DockerSvcGetInfoResponse } from './dockerSvcGetInfoResponse';
import { DockerSvcLaunchContainerRequest } from './dockerSvcLaunchContainerRequest';
import { DockerSvcLaunchContainerResponse } from './dockerSvcLaunchContainerResponse';
import { DockerSvcLaunchInfo } from './dockerSvcLaunchInfo';
import { DockerSvcLaunchOptions } from './dockerSvcLaunchOptions';
import { DownloadSvcDownloadDetails } from './downloadSvcDownloadDetails';
import { DownloadSvcDownloadRequest } from './downloadSvcDownloadRequest';
import { DownloadSvcDownloadsResponse } from './downloadSvcDownloadsResponse';
import { DownloadSvcErrorResponse } from './downloadSvcErrorResponse';
import { DownloadSvcGetDownloadResponse } from './downloadSvcGetDownloadResponse';
import { FirehoseSvcErrorResponse } from './firehoseSvcErrorResponse';
import { FirehoseSvcEvent } from './firehoseSvcEvent';
import { FirehoseSvcPublishRequest } from './firehoseSvcPublishRequest';
import { GenericSvcCreateObjectRequest } from './genericSvcCreateObjectRequest';
import { GenericSvcCreateObjectResponse } from './genericSvcCreateObjectResponse';
import { GenericSvcDeleteObjectRequest } from './genericSvcDeleteObjectRequest';
import { GenericSvcErrorResponse } from './genericSvcErrorResponse';
import { GenericSvcGenericObject } from './genericSvcGenericObject';
import { GenericSvcGenericObjectCreateFields } from './genericSvcGenericObjectCreateFields';
import { GenericSvcQueryRequest } from './genericSvcQueryRequest';
import { GenericSvcQueryResponse } from './genericSvcQueryResponse';
import { GenericSvcUpdateObjectRequest } from './genericSvcUpdateObjectRequest';
import { GenericSvcUpsertObjectRequest } from './genericSvcUpsertObjectRequest';
import { GenericSvcUpsertObjectResponse } from './genericSvcUpsertObjectResponse';
import { ModelSvcArchitectures } from './modelSvcArchitectures';
import { ModelSvcContainer } from './modelSvcContainer';
import { ModelSvcErrorResponse } from './modelSvcErrorResponse';
import { ModelSvcGetModelResponse } from './modelSvcGetModelResponse';
import { ModelSvcListResponse } from './modelSvcListResponse';
import { ModelSvcModel } from './modelSvcModel';
import { ModelSvcModelStatus } from './modelSvcModelStatus';
import { ModelSvcPlatform } from './modelSvcPlatform';
import { ModelSvcStatusResponse } from './modelSvcStatusResponse';
import { PromptSvcAddPromptRequest } from './promptSvcAddPromptRequest';
import { PromptSvcAddPromptResponse } from './promptSvcAddPromptResponse';
import { PromptSvcErrorResponse } from './promptSvcErrorResponse';
import { PromptSvcListPromptsRequest } from './promptSvcListPromptsRequest';
import { PromptSvcListPromptsResponse } from './promptSvcListPromptsResponse';
import { PromptSvcPrompt } from './promptSvcPrompt';
import { PromptSvcPromptStatus } from './promptSvcPromptStatus';
import { PromptSvcRemovePromptRequest } from './promptSvcRemovePromptRequest';
import { UserSvcAuthToken } from './userSvcAuthToken';
import { UserSvcChangePasswordAdminRequest } from './userSvcChangePasswordAdminRequest';
import { UserSvcChangePasswordRequest } from './userSvcChangePasswordRequest';
import { UserSvcCreateRoleRequest } from './userSvcCreateRoleRequest';
import { UserSvcCreateRoleResponse } from './userSvcCreateRoleResponse';
import { UserSvcCreateUserRequest } from './userSvcCreateUserRequest';
import { UserSvcErrorResponse } from './userSvcErrorResponse';
import { UserSvcGetPermissionsResponse } from './userSvcGetPermissionsResponse';
import { UserSvcGetPublicKeyResponse } from './userSvcGetPublicKeyResponse';
import { UserSvcGetRolesResponse } from './userSvcGetRolesResponse';
import { UserSvcGetUsersRequest } from './userSvcGetUsersRequest';
import { UserSvcGetUsersResponse } from './userSvcGetUsersResponse';
import { UserSvcIsAuthorizedRequest } from './userSvcIsAuthorizedRequest';
import { UserSvcIsAuthorizedResponse } from './userSvcIsAuthorizedResponse';
import { UserSvcLoginRequest } from './userSvcLoginRequest';
import { UserSvcLoginResponse } from './userSvcLoginResponse';
import { UserSvcPermission } from './userSvcPermission';
import { UserSvcReadUserByTokenRequest } from './userSvcReadUserByTokenRequest';
import { UserSvcReadUserByTokenResponse } from './userSvcReadUserByTokenResponse';
import { UserSvcRegisterRequest } from './userSvcRegisterRequest';
import { UserSvcRole } from './userSvcRole';
import { UserSvcSaveProfileRequest } from './userSvcSaveProfileRequest';
import { UserSvcSetRolePermissionsRequest } from './userSvcSetRolePermissionsRequest';
import { UserSvcUpserPermissionRequest } from './userSvcUpserPermissionRequest';
import { UserSvcUser } from './userSvcUser';

/* tslint:disable:no-unused-variable */
let primitives = [
                    "string",
                    "boolean",
                    "double",
                    "integer",
                    "long",
                    "float",
                    "number",
                    "any"
                 ];

let enumsMap: {[index: string]: any} = {
        "PromptSvcPromptStatus": PromptSvcPromptStatus,
}

let typeMap: {[index: string]: any} = {
    "ChatSvcAddMessageRequest": ChatSvcAddMessageRequest,
    "ChatSvcAddThreadRequest": ChatSvcAddThreadRequest,
    "ChatSvcAddThreadResponse": ChatSvcAddThreadResponse,
    "ChatSvcAsset": ChatSvcAsset,
    "ChatSvcGetMessagesResponse": ChatSvcGetMessagesResponse,
    "ChatSvcGetThreadResponse": ChatSvcGetThreadResponse,
    "ChatSvcGetThreadsResponse": ChatSvcGetThreadsResponse,
    "ChatSvcMessage": ChatSvcMessage,
    "ChatSvcThread": ChatSvcThread,
    "ChatSvcUpdateThreadRequest": ChatSvcUpdateThreadRequest,
    "ConfigSvcAppServiceConfig": ConfigSvcAppServiceConfig,
    "ConfigSvcConfig": ConfigSvcConfig,
    "ConfigSvcDownloadServiceConfig": ConfigSvcDownloadServiceConfig,
    "ConfigSvcGetConfigResponse": ConfigSvcGetConfigResponse,
    "ConfigSvcModelServiceConfig": ConfigSvcModelServiceConfig,
    "ConfigSvcSaveConfigRequest": ConfigSvcSaveConfigRequest,
    "DatastoreCondition": DatastoreCondition,
    "DatastoreContainsCondition": DatastoreContainsCondition,
    "DatastoreEqualCondition": DatastoreEqualCondition,
    "DatastoreFieldSelector": DatastoreFieldSelector,
    "DatastoreOrderBy": DatastoreOrderBy,
    "DatastoreQuery": DatastoreQuery,
    "DatastoreStartsWithCondition": DatastoreStartsWithCondition,
    "DockerSvcContainerIsRunningResponse": DockerSvcContainerIsRunningResponse,
    "DockerSvcDockerInfo": DockerSvcDockerInfo,
    "DockerSvcErrorResponse": DockerSvcErrorResponse,
    "DockerSvcGetContainerSummaryResponse": DockerSvcGetContainerSummaryResponse,
    "DockerSvcGetDockerHostResponse": DockerSvcGetDockerHostResponse,
    "DockerSvcGetInfoResponse": DockerSvcGetInfoResponse,
    "DockerSvcLaunchContainerRequest": DockerSvcLaunchContainerRequest,
    "DockerSvcLaunchContainerResponse": DockerSvcLaunchContainerResponse,
    "DockerSvcLaunchInfo": DockerSvcLaunchInfo,
    "DockerSvcLaunchOptions": DockerSvcLaunchOptions,
    "DownloadSvcDownloadDetails": DownloadSvcDownloadDetails,
    "DownloadSvcDownloadRequest": DownloadSvcDownloadRequest,
    "DownloadSvcDownloadsResponse": DownloadSvcDownloadsResponse,
    "DownloadSvcErrorResponse": DownloadSvcErrorResponse,
    "DownloadSvcGetDownloadResponse": DownloadSvcGetDownloadResponse,
    "FirehoseSvcErrorResponse": FirehoseSvcErrorResponse,
    "FirehoseSvcEvent": FirehoseSvcEvent,
    "FirehoseSvcPublishRequest": FirehoseSvcPublishRequest,
    "GenericSvcCreateObjectRequest": GenericSvcCreateObjectRequest,
    "GenericSvcCreateObjectResponse": GenericSvcCreateObjectResponse,
    "GenericSvcDeleteObjectRequest": GenericSvcDeleteObjectRequest,
    "GenericSvcErrorResponse": GenericSvcErrorResponse,
    "GenericSvcGenericObject": GenericSvcGenericObject,
    "GenericSvcGenericObjectCreateFields": GenericSvcGenericObjectCreateFields,
    "GenericSvcQueryRequest": GenericSvcQueryRequest,
    "GenericSvcQueryResponse": GenericSvcQueryResponse,
    "GenericSvcUpdateObjectRequest": GenericSvcUpdateObjectRequest,
    "GenericSvcUpsertObjectRequest": GenericSvcUpsertObjectRequest,
    "GenericSvcUpsertObjectResponse": GenericSvcUpsertObjectResponse,
    "ModelSvcArchitectures": ModelSvcArchitectures,
    "ModelSvcContainer": ModelSvcContainer,
    "ModelSvcErrorResponse": ModelSvcErrorResponse,
    "ModelSvcGetModelResponse": ModelSvcGetModelResponse,
    "ModelSvcListResponse": ModelSvcListResponse,
    "ModelSvcModel": ModelSvcModel,
    "ModelSvcModelStatus": ModelSvcModelStatus,
    "ModelSvcPlatform": ModelSvcPlatform,
    "ModelSvcStatusResponse": ModelSvcStatusResponse,
    "PromptSvcAddPromptRequest": PromptSvcAddPromptRequest,
    "PromptSvcAddPromptResponse": PromptSvcAddPromptResponse,
    "PromptSvcErrorResponse": PromptSvcErrorResponse,
    "PromptSvcListPromptsRequest": PromptSvcListPromptsRequest,
    "PromptSvcListPromptsResponse": PromptSvcListPromptsResponse,
    "PromptSvcPrompt": PromptSvcPrompt,
    "PromptSvcRemovePromptRequest": PromptSvcRemovePromptRequest,
    "UserSvcAuthToken": UserSvcAuthToken,
    "UserSvcChangePasswordAdminRequest": UserSvcChangePasswordAdminRequest,
    "UserSvcChangePasswordRequest": UserSvcChangePasswordRequest,
    "UserSvcCreateRoleRequest": UserSvcCreateRoleRequest,
    "UserSvcCreateRoleResponse": UserSvcCreateRoleResponse,
    "UserSvcCreateUserRequest": UserSvcCreateUserRequest,
    "UserSvcErrorResponse": UserSvcErrorResponse,
    "UserSvcGetPermissionsResponse": UserSvcGetPermissionsResponse,
    "UserSvcGetPublicKeyResponse": UserSvcGetPublicKeyResponse,
    "UserSvcGetRolesResponse": UserSvcGetRolesResponse,
    "UserSvcGetUsersRequest": UserSvcGetUsersRequest,
    "UserSvcGetUsersResponse": UserSvcGetUsersResponse,
    "UserSvcIsAuthorizedRequest": UserSvcIsAuthorizedRequest,
    "UserSvcIsAuthorizedResponse": UserSvcIsAuthorizedResponse,
    "UserSvcLoginRequest": UserSvcLoginRequest,
    "UserSvcLoginResponse": UserSvcLoginResponse,
    "UserSvcPermission": UserSvcPermission,
    "UserSvcReadUserByTokenRequest": UserSvcReadUserByTokenRequest,
    "UserSvcReadUserByTokenResponse": UserSvcReadUserByTokenResponse,
    "UserSvcRegisterRequest": UserSvcRegisterRequest,
    "UserSvcRole": UserSvcRole,
    "UserSvcSaveProfileRequest": UserSvcSaveProfileRequest,
    "UserSvcSetRolePermissionsRequest": UserSvcSetRolePermissionsRequest,
    "UserSvcUpserPermissionRequest": UserSvcUpserPermissionRequest,
    "UserSvcUser": UserSvcUser,
}

export class ObjectSerializer {
    public static findCorrectType(data: any, expectedType: string) {
        if (data == undefined) {
            return expectedType;
        } else if (primitives.indexOf(expectedType.toLowerCase()) !== -1) {
            return expectedType;
        } else if (expectedType === "Date") {
            return expectedType;
        } else {
            if (enumsMap[expectedType]) {
                return expectedType;
            }

            if (!typeMap[expectedType]) {
                return expectedType; // w/e we don't know the type
            }

            // Check the discriminator
            let discriminatorProperty = typeMap[expectedType].discriminator;
            if (discriminatorProperty == null) {
                return expectedType; // the type does not have a discriminator. use it.
            } else {
                if (data[discriminatorProperty]) {
                    var discriminatorType = data[discriminatorProperty];
                    if(typeMap[discriminatorType]){
                        return discriminatorType; // use the type given in the discriminator
                    } else {
                        return expectedType; // discriminator did not map to a type
                    }
                } else {
                    return expectedType; // discriminator was not present (or an empty string)
                }
            }
        }
    }

    public static serialize(data: any, type: string) {
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType: string = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData: any[] = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.serialize(datum, subType));
            }
            return transformedData;
        } else if (type === "Date") {
            return data.toISOString();
        } else {
            if (enumsMap[type]) {
                return data;
            }
            if (!typeMap[type]) { // in case we dont know the type
                return data;
            }

            // Get the actual type of this object
            type = this.findCorrectType(data, type);

            // get the map for the correct type.
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            let instance: {[index: string]: any} = {};
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.baseName] = ObjectSerializer.serialize(data[attributeType.name], attributeType.type);
            }
            return instance;
        }
    }

    public static deserialize(data: any, type: string) {
        // polymorphism may change the actual type.
        type = ObjectSerializer.findCorrectType(data, type);
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType: string = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData: any[] = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.deserialize(datum, subType));
            }
            return transformedData;
        } else if (type === "Date") {
            return new Date(data);
        } else {
            if (enumsMap[type]) {// is Enum
                return data;
            }

            if (!typeMap[type]) { // dont know the type
                return data;
            }
            let instance = new typeMap[type]();
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.name] = ObjectSerializer.deserialize(data[attributeType.baseName], attributeType.type);
            }
            return instance;
        }
    }
}

export interface Authentication {
    /**
    * Apply authentication settings to header and query params.
    */
    applyToRequest(requestOptions: localVarRequest.Options): Promise<void> | void;
}

export class HttpBasicAuth implements Authentication {
    public username: string = '';
    public password: string = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        requestOptions.auth = {
            username: this.username, password: this.password
        }
    }
}

export class HttpBearerAuth implements Authentication {
    public accessToken: string | (() => string) = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (requestOptions && requestOptions.headers) {
            const accessToken = typeof this.accessToken === 'function'
                            ? this.accessToken()
                            : this.accessToken;
            requestOptions.headers["Authorization"] = "Bearer " + accessToken;
        }
    }
}

export class ApiKeyAuth implements Authentication {
    public apiKey: string = '';

    constructor(private location: string, private paramName: string) {
    }

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (this.location == "query") {
            (<any>requestOptions.qs)[this.paramName] = this.apiKey;
        } else if (this.location == "header" && requestOptions && requestOptions.headers) {
            requestOptions.headers[this.paramName] = this.apiKey;
        } else if (this.location == 'cookie' && requestOptions && requestOptions.headers) {
            if (requestOptions.headers['Cookie']) {
                requestOptions.headers['Cookie'] += '; ' + this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
            else {
                requestOptions.headers['Cookie'] = this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
        }
    }
}

export class OAuth implements Authentication {
    public accessToken: string = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (requestOptions && requestOptions.headers) {
            requestOptions.headers["Authorization"] = "Bearer " + this.accessToken;
        }
    }
}

export class VoidAuth implements Authentication {
    public username: string = '';
    public password: string = '';

    applyToRequest(_: localVarRequest.Options): void {
        // Do nothing
    }
}

export type Interceptor = (requestOptions: localVarRequest.Options) => (Promise<void> | void);

import { ChatSvcAddMessageRequest } from './chatSvcAddMessageRequest.mjs';
import { ChatSvcAddThreadRequest } from './chatSvcAddThreadRequest.mjs';
import { ChatSvcAddThreadResponse } from './chatSvcAddThreadResponse.mjs';
import { ChatSvcAsset } from './chatSvcAsset.mjs';
import { ChatSvcGetMessagesResponse } from './chatSvcGetMessagesResponse.mjs';
import { ChatSvcGetThreadResponse } from './chatSvcGetThreadResponse.mjs';
import { ChatSvcGetThreadsResponse } from './chatSvcGetThreadsResponse.mjs';
import { ChatSvcMessage } from './chatSvcMessage.mjs';
import { ChatSvcThread } from './chatSvcThread.mjs';
import { ChatSvcUpdateThreadRequest } from './chatSvcUpdateThreadRequest.mjs';
import { ConfigSvcAppServiceConfig } from './configSvcAppServiceConfig.mjs';
import { ConfigSvcConfig } from './configSvcConfig.mjs';
import { ConfigSvcDownloadServiceConfig } from './configSvcDownloadServiceConfig.mjs';
import { ConfigSvcGetConfigResponse } from './configSvcGetConfigResponse.mjs';
import { ConfigSvcModelServiceConfig } from './configSvcModelServiceConfig.mjs';
import { ConfigSvcSaveConfigRequest } from './configSvcSaveConfigRequest.mjs';
import { DatastoreCondition } from './datastoreCondition.mjs';
import { DatastoreContainsCondition } from './datastoreContainsCondition.mjs';
import { DatastoreEqualCondition } from './datastoreEqualCondition.mjs';
import { DatastoreFieldSelector } from './datastoreFieldSelector.mjs';
import { DatastoreOrderBy } from './datastoreOrderBy.mjs';
import { DatastoreQuery } from './datastoreQuery.mjs';
import { DatastoreStartsWithCondition } from './datastoreStartsWithCondition.mjs';
import { DockerSvcContainerIsRunningResponse } from './dockerSvcContainerIsRunningResponse.mjs';
import { DockerSvcDockerInfo } from './dockerSvcDockerInfo.mjs';
import { DockerSvcErrorResponse } from './dockerSvcErrorResponse.mjs';
import { DockerSvcGetContainerSummaryResponse } from './dockerSvcGetContainerSummaryResponse.mjs';
import { DockerSvcGetDockerHostResponse } from './dockerSvcGetDockerHostResponse.mjs';
import { DockerSvcGetInfoResponse } from './dockerSvcGetInfoResponse.mjs';
import { DockerSvcLaunchContainerRequest } from './dockerSvcLaunchContainerRequest.mjs';
import { DockerSvcLaunchContainerResponse } from './dockerSvcLaunchContainerResponse.mjs';
import { DockerSvcLaunchInfo } from './dockerSvcLaunchInfo.mjs';
import { DockerSvcLaunchOptions } from './dockerSvcLaunchOptions.mjs';
import { DownloadSvcDownloadDetails } from './downloadSvcDownloadDetails.mjs';
import { DownloadSvcDownloadRequest } from './downloadSvcDownloadRequest.mjs';
import { DownloadSvcDownloadsResponse } from './downloadSvcDownloadsResponse.mjs';
import { DownloadSvcErrorResponse } from './downloadSvcErrorResponse.mjs';
import { DownloadSvcGetDownloadResponse } from './downloadSvcGetDownloadResponse.mjs';
import { FirehoseSvcErrorResponse } from './firehoseSvcErrorResponse.mjs';
import { FirehoseSvcEvent } from './firehoseSvcEvent.mjs';
import { FirehoseSvcPublishRequest } from './firehoseSvcPublishRequest.mjs';
import { GenericSvcCreateObjectRequest } from './genericSvcCreateObjectRequest.mjs';
import { GenericSvcCreateObjectResponse } from './genericSvcCreateObjectResponse.mjs';
import { GenericSvcDeleteObjectRequest } from './genericSvcDeleteObjectRequest.mjs';
import { GenericSvcErrorResponse } from './genericSvcErrorResponse.mjs';
import { GenericSvcGenericObject } from './genericSvcGenericObject.mjs';
import { GenericSvcGenericObjectCreateFields } from './genericSvcGenericObjectCreateFields.mjs';
import { GenericSvcQueryRequest } from './genericSvcQueryRequest.mjs';
import { GenericSvcQueryResponse } from './genericSvcQueryResponse.mjs';
import { GenericSvcUpdateObjectRequest } from './genericSvcUpdateObjectRequest.mjs';
import { GenericSvcUpsertObjectRequest } from './genericSvcUpsertObjectRequest.mjs';
import { GenericSvcUpsertObjectResponse } from './genericSvcUpsertObjectResponse.mjs';
import { ModelSvcArchitectures } from './modelSvcArchitectures.mjs';
import { ModelSvcContainer } from './modelSvcContainer.mjs';
import { ModelSvcErrorResponse } from './modelSvcErrorResponse.mjs';
import { ModelSvcGetModelResponse } from './modelSvcGetModelResponse.mjs';
import { ModelSvcListResponse } from './modelSvcListResponse.mjs';
import { ModelSvcModel } from './modelSvcModel.mjs';
import { ModelSvcModelStatus } from './modelSvcModelStatus.mjs';
import { ModelSvcPlatform } from './modelSvcPlatform.mjs';
import { ModelSvcStatusResponse } from './modelSvcStatusResponse.mjs';
import { NodeSvcErrorResponse } from './nodeSvcErrorResponse.mjs';
import { NodeSvcGPU } from './nodeSvcGPU.mjs';
import { NodeSvcListNodesResponse } from './nodeSvcListNodesResponse.mjs';
import { NodeSvcNode } from './nodeSvcNode.mjs';
import { NodeSvcProcess } from './nodeSvcProcess.mjs';
import { PromptSvcAddPromptRequest } from './promptSvcAddPromptRequest.mjs';
import { PromptSvcAddPromptResponse } from './promptSvcAddPromptResponse.mjs';
import { PromptSvcErrorResponse } from './promptSvcErrorResponse.mjs';
import { PromptSvcListPromptsRequest } from './promptSvcListPromptsRequest.mjs';
import { PromptSvcListPromptsResponse } from './promptSvcListPromptsResponse.mjs';
import { PromptSvcPrompt } from './promptSvcPrompt.mjs';
import { PromptSvcPromptStatus } from './promptSvcPromptStatus.mjs';
import { PromptSvcRemovePromptRequest } from './promptSvcRemovePromptRequest.mjs';
import { UserSvcAddUserToOrganizationRequest } from './userSvcAddUserToOrganizationRequest.mjs';
import { UserSvcAuthToken } from './userSvcAuthToken.mjs';
import { UserSvcChangePasswordAdminRequest } from './userSvcChangePasswordAdminRequest.mjs';
import { UserSvcChangePasswordRequest } from './userSvcChangePasswordRequest.mjs';
import { UserSvcContact } from './userSvcContact.mjs';
import { UserSvcCreateOrganizationRequest } from './userSvcCreateOrganizationRequest.mjs';
import { UserSvcCreateRoleRequest } from './userSvcCreateRoleRequest.mjs';
import { UserSvcCreateRoleResponse } from './userSvcCreateRoleResponse.mjs';
import { UserSvcCreateUserRequest } from './userSvcCreateUserRequest.mjs';
import { UserSvcErrorResponse } from './userSvcErrorResponse.mjs';
import { UserSvcGetPermissionsResponse } from './userSvcGetPermissionsResponse.mjs';
import { UserSvcGetPublicKeyResponse } from './userSvcGetPublicKeyResponse.mjs';
import { UserSvcGetRolesResponse } from './userSvcGetRolesResponse.mjs';
import { UserSvcGetUsersRequest } from './userSvcGetUsersRequest.mjs';
import { UserSvcGetUsersResponse } from './userSvcGetUsersResponse.mjs';
import { UserSvcIsAuthorizedRequest } from './userSvcIsAuthorizedRequest.mjs';
import { UserSvcIsAuthorizedResponse } from './userSvcIsAuthorizedResponse.mjs';
import { UserSvcLoginRequest } from './userSvcLoginRequest.mjs';
import { UserSvcLoginResponse } from './userSvcLoginResponse.mjs';
import { UserSvcPermission } from './userSvcPermission.mjs';
import { UserSvcReadUserByTokenRequest } from './userSvcReadUserByTokenRequest.mjs';
import { UserSvcReadUserByTokenResponse } from './userSvcReadUserByTokenResponse.mjs';
import { UserSvcRegisterRequest } from './userSvcRegisterRequest.mjs';
import { UserSvcRole } from './userSvcRole.mjs';
import { UserSvcSaveProfileRequest } from './userSvcSaveProfileRequest.mjs';
import { UserSvcSetRolePermissionsRequest } from './userSvcSetRolePermissionsRequest.mjs';
import { UserSvcUpserPermissionRequest } from './userSvcUpserPermissionRequest.mjs';
import { UserSvcUser } from './userSvcUser.mjs';

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
let enumsMap = {
    "PromptSvcPromptStatus": PromptSvcPromptStatus,
};
let typeMap = {
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
    "NodeSvcErrorResponse": NodeSvcErrorResponse,
    "NodeSvcGPU": NodeSvcGPU,
    "NodeSvcListNodesResponse": NodeSvcListNodesResponse,
    "NodeSvcNode": NodeSvcNode,
    "NodeSvcProcess": NodeSvcProcess,
    "PromptSvcAddPromptRequest": PromptSvcAddPromptRequest,
    "PromptSvcAddPromptResponse": PromptSvcAddPromptResponse,
    "PromptSvcErrorResponse": PromptSvcErrorResponse,
    "PromptSvcListPromptsRequest": PromptSvcListPromptsRequest,
    "PromptSvcListPromptsResponse": PromptSvcListPromptsResponse,
    "PromptSvcPrompt": PromptSvcPrompt,
    "PromptSvcRemovePromptRequest": PromptSvcRemovePromptRequest,
    "UserSvcAddUserToOrganizationRequest": UserSvcAddUserToOrganizationRequest,
    "UserSvcAuthToken": UserSvcAuthToken,
    "UserSvcChangePasswordAdminRequest": UserSvcChangePasswordAdminRequest,
    "UserSvcChangePasswordRequest": UserSvcChangePasswordRequest,
    "UserSvcContact": UserSvcContact,
    "UserSvcCreateOrganizationRequest": UserSvcCreateOrganizationRequest,
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
};
class ObjectSerializer {
    static findCorrectType(data, expectedType) {
        if (data == undefined) {
            return expectedType;
        }
        else if (primitives.indexOf(expectedType.toLowerCase()) !== -1) {
            return expectedType;
        }
        else if (expectedType === "Date") {
            return expectedType;
        }
        else {
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
            }
            else {
                if (data[discriminatorProperty]) {
                    var discriminatorType = data[discriminatorProperty];
                    if (typeMap[discriminatorType]) {
                        return discriminatorType; // use the type given in the discriminator
                    }
                    else {
                        return expectedType; // discriminator did not map to a type
                    }
                }
                else {
                    return expectedType; // discriminator was not present (or an empty string)
                }
            }
        }
    }
    static serialize(data, type) {
        if (data == undefined) {
            return data;
        }
        else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        }
        else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.serialize(datum, subType));
            }
            return transformedData;
        }
        else if (type === "Date") {
            return data.toISOString();
        }
        else {
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
            let instance = {};
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.baseName] = ObjectSerializer.serialize(data[attributeType.name], attributeType.type);
            }
            return instance;
        }
    }
    static deserialize(data, type) {
        // polymorphism may change the actual type.
        type = ObjectSerializer.findCorrectType(data, type);
        if (data == undefined) {
            return data;
        }
        else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        }
        else if (type.lastIndexOf("Array<", 0) === 0) { // string.startsWith pre es6
            let subType = type.replace("Array<", ""); // Array<Type> => Type>
            subType = subType.substring(0, subType.length - 1); // Type> => Type
            let transformedData = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.deserialize(datum, subType));
            }
            return transformedData;
        }
        else if (type === "Date") {
            return new Date(data);
        }
        else {
            if (enumsMap[type]) { // is Enum
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
class HttpBasicAuth {
    constructor() {
        this.username = '';
        this.password = '';
    }
    applyToRequest(requestOptions) {
        requestOptions.auth = {
            username: this.username, password: this.password
        };
    }
}
class HttpBearerAuth {
    constructor() {
        this.accessToken = '';
    }
    applyToRequest(requestOptions) {
        if (requestOptions && requestOptions.headers) {
            const accessToken = typeof this.accessToken === 'function'
                ? this.accessToken()
                : this.accessToken;
            requestOptions.headers["Authorization"] = "Bearer " + accessToken;
        }
    }
}
class ApiKeyAuth {
    constructor(location, paramName) {
        this.location = location;
        this.paramName = paramName;
        this.apiKey = '';
    }
    applyToRequest(requestOptions) {
        if (this.location == "query") {
            requestOptions.qs[this.paramName] = this.apiKey;
        }
        else if (this.location == "header" && requestOptions && requestOptions.headers) {
            requestOptions.headers[this.paramName] = this.apiKey;
        }
        else if (this.location == 'cookie' && requestOptions && requestOptions.headers) {
            if (requestOptions.headers['Cookie']) {
                requestOptions.headers['Cookie'] += '; ' + this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
            else {
                requestOptions.headers['Cookie'] = this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
        }
    }
}
class OAuth {
    constructor() {
        this.accessToken = '';
    }
    applyToRequest(requestOptions) {
        if (requestOptions && requestOptions.headers) {
            requestOptions.headers["Authorization"] = "Bearer " + this.accessToken;
        }
    }
}
class VoidAuth {
    constructor() {
        this.username = '';
        this.password = '';
    }
    applyToRequest(_) {
        // Do nothing
    }
}

export { ApiKeyAuth, ChatSvcAddMessageRequest, ChatSvcAddThreadRequest, ChatSvcAddThreadResponse, ChatSvcAsset, ChatSvcGetMessagesResponse, ChatSvcGetThreadResponse, ChatSvcGetThreadsResponse, ChatSvcMessage, ChatSvcThread, ChatSvcUpdateThreadRequest, ConfigSvcAppServiceConfig, ConfigSvcConfig, ConfigSvcDownloadServiceConfig, ConfigSvcGetConfigResponse, ConfigSvcModelServiceConfig, ConfigSvcSaveConfigRequest, DatastoreCondition, DatastoreContainsCondition, DatastoreEqualCondition, DatastoreFieldSelector, DatastoreOrderBy, DatastoreQuery, DatastoreStartsWithCondition, DockerSvcContainerIsRunningResponse, DockerSvcDockerInfo, DockerSvcErrorResponse, DockerSvcGetContainerSummaryResponse, DockerSvcGetDockerHostResponse, DockerSvcGetInfoResponse, DockerSvcLaunchContainerRequest, DockerSvcLaunchContainerResponse, DockerSvcLaunchInfo, DockerSvcLaunchOptions, DownloadSvcDownloadDetails, DownloadSvcDownloadRequest, DownloadSvcDownloadsResponse, DownloadSvcErrorResponse, DownloadSvcGetDownloadResponse, FirehoseSvcErrorResponse, FirehoseSvcEvent, FirehoseSvcPublishRequest, GenericSvcCreateObjectRequest, GenericSvcCreateObjectResponse, GenericSvcDeleteObjectRequest, GenericSvcErrorResponse, GenericSvcGenericObject, GenericSvcGenericObjectCreateFields, GenericSvcQueryRequest, GenericSvcQueryResponse, GenericSvcUpdateObjectRequest, GenericSvcUpsertObjectRequest, GenericSvcUpsertObjectResponse, HttpBasicAuth, HttpBearerAuth, ModelSvcArchitectures, ModelSvcContainer, ModelSvcErrorResponse, ModelSvcGetModelResponse, ModelSvcListResponse, ModelSvcModel, ModelSvcModelStatus, ModelSvcPlatform, ModelSvcStatusResponse, NodeSvcErrorResponse, NodeSvcGPU, NodeSvcListNodesResponse, NodeSvcNode, NodeSvcProcess, OAuth, ObjectSerializer, PromptSvcAddPromptRequest, PromptSvcAddPromptResponse, PromptSvcErrorResponse, PromptSvcListPromptsRequest, PromptSvcListPromptsResponse, PromptSvcPrompt, PromptSvcPromptStatus, PromptSvcRemovePromptRequest, UserSvcAddUserToOrganizationRequest, UserSvcAuthToken, UserSvcChangePasswordAdminRequest, UserSvcChangePasswordRequest, UserSvcContact, UserSvcCreateOrganizationRequest, UserSvcCreateRoleRequest, UserSvcCreateRoleResponse, UserSvcCreateUserRequest, UserSvcErrorResponse, UserSvcGetPermissionsResponse, UserSvcGetPublicKeyResponse, UserSvcGetRolesResponse, UserSvcGetUsersRequest, UserSvcGetUsersResponse, UserSvcIsAuthorizedRequest, UserSvcIsAuthorizedResponse, UserSvcLoginRequest, UserSvcLoginResponse, UserSvcPermission, UserSvcReadUserByTokenRequest, UserSvcReadUserByTokenResponse, UserSvcRegisterRequest, UserSvcRole, UserSvcSaveProfileRequest, UserSvcSetRolePermissionsRequest, UserSvcUpserPermissionRequest, UserSvcUser, VoidAuth };
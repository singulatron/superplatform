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
import type { UsertypesChangePasswordAdminRequest, UsertypesChangePasswordRequest, UsertypesCreateRoleRequest, UsertypesCreateRoleResponse, UsertypesCreateUserRequest, UsertypesGetPermissionsResponse, UsertypesGetPublicKeyResponse, UsertypesGetRolesResponse, UsertypesGetUsersRequest, UsertypesGetUsersResponse, UsertypesIsAuthorizedRequest, UsertypesIsAuthorizedResponse, UsertypesLoginRequest, UsertypesLoginResponse, UsertypesReadUserByTokenRequest, UsertypesReadUserByTokenResponse, UsertypesRegisterRequest, UsertypesSetRolePermissionsRequest, UsertypesUpserPermissionRequest } from '../models/index';
export interface AddPermissionToRoleRequest {
    roleId: string;
    permissionId: string;
}
export interface ChangePasswordRequest {
    request: UsertypesChangePasswordRequest;
}
export interface ChangePasswordAdminRequest {
    request: UsertypesChangePasswordAdminRequest;
}
export interface CreateRoleRequest {
    request: UsertypesCreateRoleRequest;
}
export interface CreateUserRequest {
    request: UsertypesCreateUserRequest;
}
export interface DeleteRoleRequest {
    roleId: string;
}
export interface DeleteUserRequest {
    userId: string;
}
export interface GetPermissionsByRoleRequest {
    roleId: number;
}
export interface GetUserByTokenRequest {
    body: UsertypesReadUserByTokenRequest;
}
export interface GetUsersRequest {
    request?: UsertypesGetUsersRequest;
}
export interface IsAuthorizedRequest {
    permissionId: string;
    body: UsertypesIsAuthorizedRequest;
}
export interface LoginRequest {
    request: UsertypesLoginRequest;
}
export interface RegisterRequest {
    body: UsertypesRegisterRequest;
}
export interface SetRolePermissionRequest {
    roleId: string;
    body: UsertypesSetRolePermissionsRequest;
}
export interface UpsertPermissionRequest {
    permissionId: string;
    requestBody: UsertypesUpserPermissionRequest;
}
/**
 *
 */
export declare class UserServiceApi extends runtime.BaseAPI {
    /**
     * Adds a specific permission to a role identified by roleId.  Requires the `user-svc:permission:assign` permission.
     * Add Permission to Role
     */
    addPermissionToRoleRaw(requestParameters: AddPermissionToRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Adds a specific permission to a role identified by roleId.  Requires the `user-svc:permission:assign` permission.
     * Add Permission to Role
     */
    addPermissionToRole(requestParameters: AddPermissionToRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Allows an authenticated user to change their own password.
     * Change User Password
     */
    changePasswordRaw(requestParameters: ChangePasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an authenticated user to change their own password.
     * Change User Password
     */
    changePassword(requestParameters: ChangePasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Allows an administrator to change a user\'s password.
     * Change User Password (Admin)
     */
    changePasswordAdminRaw(requestParameters: ChangePasswordAdminRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an administrator to change a user\'s password.
     * Change User Password (Admin)
     */
    changePasswordAdmin(requestParameters: ChangePasswordAdminRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Create a new role. <b>The role ID must be prefixed by the callers username (email).</b> Eg. if the owner\'s email/username is `petstore-svc` the role should look like `petstore-svc:admin`. The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.  Requires the `user-svc:role:create` permission.
     * Create a New Role
     */
    createRoleRaw(requestParameters: CreateRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesCreateRoleResponse>>;
    /**
     * Create a new role. <b>The role ID must be prefixed by the callers username (email).</b> Eg. if the owner\'s email/username is `petstore-svc` the role should look like `petstore-svc:admin`. The user account who creates the role will become the owner of that role, and only the owner will be able to edit the role.  Requires the `user-svc:role:create` permission.
     * Create a New Role
     */
    createRole(requestParameters: CreateRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesCreateRoleResponse>;
    /**
     * Allows an authenticated administrator to create a new user with specified details.
     * Create a New User
     */
    createUserRaw(requestParameters: CreateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Allows an authenticated administrator to create a new user with specified details.
     * Create a New User
     */
    createUser(requestParameters: CreateUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Delete a role based on the role ID.
     * Delete a Role
     */
    deleteRoleRaw(requestParameters: DeleteRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Delete a role based on the role ID.
     * Delete a Role
     */
    deleteRole(requestParameters: DeleteRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Delete a user based on the user ID.
     * Delete a User
     */
    deleteUserRaw(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Delete a user based on the user ID.
     * Delete a User
     */
    deleteUser(requestParameters: DeleteUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Retrieve permissions associated with a specific role ID.
     * Get Permissions by Role
     */
    getPermissionsByRoleRaw(requestParameters: GetPermissionsByRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesGetPermissionsResponse>>;
    /**
     * Retrieve permissions associated with a specific role ID.
     * Get Permissions by Role
     */
    getPermissionsByRole(requestParameters: GetPermissionsByRoleRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesGetPermissionsResponse>;
    /**
     * Get the public key to descrypt the JWT.
     * Ge Public Key
     */
    getPublicKeyRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesGetPublicKeyResponse>>;
    /**
     * Get the public key to descrypt the JWT.
     * Ge Public Key
     */
    getPublicKey(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesGetPublicKeyResponse>;
    /**
     * Retrieve all roles from the user service.
     * Get all Roles
     */
    getRolesRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesGetRolesResponse>>;
    /**
     * Retrieve all roles from the user service.
     * Get all Roles
     */
    getRoles(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesGetRolesResponse>;
    /**
     * Retrieve user information based on an authentication token.
     * Read User by Token
     */
    getUserByTokenRaw(requestParameters: GetUserByTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesReadUserByTokenResponse>>;
    /**
     * Retrieve user information based on an authentication token.
     * Read User by Token
     */
    getUserByToken(requestParameters: GetUserByTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesReadUserByTokenResponse>;
    /**
     * Fetches a list of users with optional query filters and pagination.
     * List Users
     */
    getUsersRaw(requestParameters: GetUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesGetUsersResponse>>;
    /**
     * Fetches a list of users with optional query filters and pagination.
     * List Users
     */
    getUsers(requestParameters?: GetUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesGetUsersResponse>;
    /**
     * Check if a user is authorized for a specific permission.
     * Is Authorized
     */
    isAuthorizedRaw(requestParameters: IsAuthorizedRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesIsAuthorizedResponse>>;
    /**
     * Check if a user is authorized for a specific permission.
     * Is Authorized
     */
    isAuthorized(requestParameters: IsAuthorizedRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesIsAuthorizedResponse>;
    /**
     * Authenticates a user and returns a token.
     * Login
     */
    loginRaw(requestParameters: LoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsertypesLoginResponse>>;
    /**
     * Authenticates a user and returns a token.
     * Login
     */
    login(requestParameters: LoginRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsertypesLoginResponse>;
    /**
     * Register a new user with a name, email, and password.
     * Register a New User
     */
    registerRaw(requestParameters: RegisterRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Register a new user with a name, email, and password.
     * Register a New User
     */
    register(requestParameters: RegisterRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Set permissions for a specified role. The caller can add permissions it owns to any role. If the caller tries to add a permission it doesn\'t own to a role, `StatusBadRequest` will be returned.
     * Set Role Permissions
     */
    setRolePermissionRaw(requestParameters: SetRolePermissionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Set permissions for a specified role. The caller can add permissions it owns to any role. If the caller tries to add a permission it doesn\'t own to a role, `StatusBadRequest` will be returned.
     * Set Role Permissions
     */
    setRolePermission(requestParameters: SetRolePermissionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
    /**
     * Creates or updates a permission. <b>The permission ID must be prefixed by the callers username (email).</b> Eg. if the owner\'s email/username is `petstore-svc` the permission should look like `petstore-svc:pet:edit`.  Requires the `user-svc:permission:create` permission.
     * Upsert a Permission
     */
    upsertPermissionRaw(requestParameters: UpsertPermissionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<object>>;
    /**
     * Creates or updates a permission. <b>The permission ID must be prefixed by the callers username (email).</b> Eg. if the owner\'s email/username is `petstore-svc` the permission should look like `petstore-svc:pet:edit`.  Requires the `user-svc:permission:create` permission.
     * Upsert a Permission
     */
    upsertPermission(requestParameters: UpsertPermissionRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<object>;
}
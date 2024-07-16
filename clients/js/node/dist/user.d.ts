import { Query } from "./generic";
export interface User {
    id?: string;
    createdAt?: Date;
    updatedAt?: Date;
    deletedAt?: Date | undefined;
    name?: string;
    email?: string;
    passwordHash?: string;
    roleIds?: string[];
    authTokenIds?: string[];
}
export interface RegisterRequest {
    name?: string;
    email?: string;
}
export interface RegisterResponse {
}
export interface LoginRequest {
    name?: string;
    email?: string;
}
export interface LoginResponse {
    token: AuthToken;
}
export interface SaveProfileRequest {
    name?: string;
    email?: string;
}
export interface SaveProfileResponse {
}
export interface ChangePasswordRequest {
    email?: string;
    currentPassword?: string;
    newPassword?: string;
}
export interface ChangePasswordResponse {
}
export interface ChangePasswordAdminRequest {
    email?: string;
    newPassword?: string;
}
export interface ChangePasswordAdminResponse {
}
export declare const RoleAdmin: Role;
export declare const RoleUser: Role;
export interface Role {
    id?: string;
    createdAt?: Date;
    updatedAt?: Date;
    name: string;
    description?: string;
    permissionIds: string[];
}
export interface CreateRoleRequest {
    name: string;
    description: string;
    permissionIds: string[];
}
export interface CreateRoleResponse {
    role?: Role;
}
export interface DeleteRoleRequest {
    roleId?: string;
}
export interface DeleteRoleResponse {
}
export interface RemoveRoleRequest {
    userId?: string;
    roleId?: string;
}
export interface RemoveRoleResponse {
}
export interface Permission {
    id?: string;
    createdAt?: Date;
    updatedAt?: Date;
    name: string;
    description: string;
}
export interface CreatePermissionRequest {
    name: string;
    description: string;
}
export interface CreatePermissionResponse {
}
export interface AuthToken {
    id?: string;
    createdAt?: Date;
    updatedAt?: Date;
    deletedAt?: Date | undefined;
    userId?: string;
    token?: string;
}
export interface ReadUserByTokenResponse {
    user: User;
}
export interface GetUsersRequest {
    query: Query;
}
export interface GetUsersResponse {
    users: User[];
    after: string;
    count?: number;
}
export interface CreateUserRequest {
    user: User;
    password: string;
    roleIds: string[];
}
export interface CreateUserResponse {
}
export interface GetRolesRequest {
}
export interface GetRolesResposne {
    roles: Role[];
}
export interface GetPermissionsRequest {
}
export interface GetPermissionsResposne {
    permissions: Permission[];
}
export interface SetRolePermissionsRequest {
}
export interface SetRolePermissionsResponse {
}
export interface DeleteUserRequest {
    userId: string;
}
export interface DeleteUserResponse {
}

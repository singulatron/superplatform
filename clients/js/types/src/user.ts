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

// eslint-disable-next-line
export interface RegisterResponse {}

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

// eslint-disable-next-line
export interface SaveProfileResponse {}

export interface ChangePasswordRequest {
  email?: string;
  currentPassword?: string;
  newPassword?: string;
}

// eslint-disable-next-line
export interface ChangePasswordResponse {}

export interface ChangePasswordAdminRequest {
  email?: string;
  newPassword?: string;
}

// eslint-disable-next-line
export interface ChangePasswordAdminResponse {}

export const RoleAdmin: Role = {
  id: "role.admin",
  name: "Admin Role",
  permissionIds: [],
};

export const RoleUser: Role = {
  id: "role.user",
  name: "User Role",
  permissionIds: [],
};

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

// eslint-disable-next-line
export interface DeleteRoleResponse {}

export interface RemoveRoleRequest {
  userId?: string;
  roleId?: string;
}

// eslint-disable-next-line
export interface RemoveRoleResponse {}

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

// eslint-disable-next-line
export interface CreatePermissionResponse {}

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

// eslint-disable-next-line
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

// eslint-disable-next-line
export interface CreateUserResponse {}

// eslint-disable-next-line
export interface GetRolesRequest {}

export interface GetRolesResposne {
  roles: Role[];
}

// eslint-disable-next-line
export interface GetPermissionsRequest {}

export interface GetPermissionsResposne {
  permissions: Permission[];
}

// eslint-disable-next-line
export interface SetRolePermissionsRequest {}

// eslint-disable-next-line
export interface SetRolePermissionsResponse {}

export interface DeleteUserRequest {
  userId: string;
}
// eslint-disable-next-line
export interface DeleteUserResponse {}

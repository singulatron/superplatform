import { ClientOptions, call } from "./util";
import * as user from "@singulatron/types";
import { equal, field } from "@singulatron/types";
import { Method } from "axios";

export class UserService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any, method: Method = "POST"): Promise<any> {
    return call(this.options, endpoint, request, method);
  }

  login(email: string, password: string): Promise<user.LoginResponse> {
    return this.call("/user-service/login", {
      email: email,
      password: password,
    });
  }

  readUserByToken(token: string): Promise<user.ReadUserByTokenResponse> {
    return this.call("/user-service/user/by-token", {
      token: token,
    });
  }

  getUsers(request: user.GetUsersRequest): Promise<user.GetUsersResponse> {
    return this.call("/user-service/users", request);
  }

  /** Save profile on behalf of a user */
  saveProfile(email: string, name: string): Promise<user.SaveProfileResponse> {
    const request: user.SaveProfileRequest = {
      email: email,
      name: name,
    };
    return this.call(`/user-service/user/anyIDTodo`, request, "PUT");
  }

  changePassword(
    email: string,
    currentPassword: string,
    newPassword: string
  ): Promise<user.ChangePasswordResponse> {
    const request: user.ChangePasswordRequest = {
      email: email,
      currentPassword: currentPassword,
      newPassword: newPassword,
    };
    return this.call("/user-service/change-password", request);
  }

  changePasswordAdmin(
    email: string,
    newPassword: string
  ): Promise<user.ChangePasswordAdminResponse> {
    const request: user.ChangePasswordAdminRequest = {
      email: email,
      newPassword: newPassword,
    };
    return this.call("/user-service/change-password-admin", request);
  }

  /** Create a user - alternative to registration
   */
  createUser(
    user: user.User,
    password: string,
    roleIds: string[]
  ): Promise<user.CreateUserResponse> {
    const request: user.CreateUserRequest = {
      user: user,
      password: password,
      roleIds: roleIds,
    };
    return this.call("/user-service/create-user", request);
  }

  getRoles(): Promise<user.GetRolesResposne> {
    return this.call("/user-service/get-roles", {});
  }

  getPermissions(): Promise<user.GetPermissionsResposne> {
    return this.call("/user-service/get-permissions", {});
  }

  setRolePermissions(
    roleId: string,
    permissionIds: string[]
  ): Promise<user.SetRolePermissionsResponse> {
    const request: user.SetRolePermissionsRequest = {
      roleId: roleId,
      permissionIds: permissionIds,
    };
    return this.call("/user-service/set-role-permissions", request);
  }

  deleteRole(roleId: string): Promise<user.DeleteRoleResponse> {
    const request: user.DeleteRoleRequest = {
      roleId: roleId,
    };
    return this.call(`/user-service/role/${roleId}`, request, "DELETE");
  }

  deleteUser(userId: string): Promise<user.DeleteUserResponse> {
    const request: user.DeleteUserRequest = {
      userId: userId,
    };
    return this.call(`/user-service/user/${userId}`, request, "DELETE");
  }

  async getUser(id: string): Promise<user.User | undefined> {
    let rsp = await this.getUsers({
      query: {
        conditions: [equal(field("id"), id)],
      },
    });
    return rsp.users[0];
  }
}

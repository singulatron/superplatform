import { ClientOptions } from "./util";
import * as user from "@singulatron/types";
export declare class UserService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    login(email: string, password: string): Promise<user.LoginResponse>;
    readUserByToken(token: string): Promise<user.ReadUserByTokenResponse>;
    getUsers(request: user.GetUsersRequest): Promise<user.GetUsersResponse>;
    /** Save profile on behalf of a user */
    saveProfile(email: string, name: string): Promise<user.SaveProfileResponse>;
    changePassword(email: string, currentPassword: string, newPassword: string): Promise<user.ChangePasswordResponse>;
    changePasswordAdmin(email: string, newPassword: string): Promise<user.ChangePasswordAdminResponse>;
    /** Create a user - alternative to registration
     */
    createUser(user: user.User, password: string, roleIds: string[]): Promise<user.CreateUserResponse>;
    getRoles(): Promise<user.GetRolesResposne>;
    getPermissions(): Promise<user.GetPermissionsResposne>;
    setRolePermissions(roleId: string, permissionIds: string[]): Promise<user.SetRolePermissionsResponse>;
    deleteRole(roleId: string): Promise<user.DeleteRoleResponse>;
    deleteUser(userId: string): Promise<user.DeleteUserResponse>;
    getUser(id: string): Promise<user.User | undefined>;
}

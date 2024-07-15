/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { CookieService } from 'ngx-cookie-service';
import {
	ReplaySubject,
	Observable,
	firstValueFrom,
	first,
	BehaviorSubject,
} from 'rxjs';
import { Router } from '@angular/router';
import { Query, equal, field } from './generic.service';

@Injectable({
	providedIn: 'root',
})
export class UserService {
	private token: string = '';

	private userSubject = new ReplaySubject<User>(1);
	/** Current logged in user */
	public user$ = this.userSubject.asObservable();

	private userCache: { [id: string]: BehaviorSubject<User | undefined> } = {};

	constructor(
		private localtron: LocaltronService,
		private cookieService: CookieService,
		private router: Router
	) {
		this.init();
	}

	noop() {}

	async init() {
		this.getToken();

		if (!this.hasToken()) {
			try {
				const rsp = await this.login('singulatron', 'changeme');
				this.setToken(rsp.token.token as string);
			} catch {
				console.error('Login with default credentials failed');
				this.router.navigateByUrl('/login');
				return;
			}

			if (!this.hasToken()) {
				console.error('Something is wrong with the setting of cookies');
				this.router.navigateByUrl('/login');
				return;
			}
		}

		try {
			const rsp = await this.readUserByToken(this.token);
			this.userSubject.next(rsp.user);
		} catch {
			console.error('Cannot read user even with a token');
			this.router.navigateByUrl('/login');
		}
	}

	getToken(): string {
		if (this.token) {
			return this.token;
		}
		this.token = this.cookieService.get('the_token');
		return this.token;
	}

	setToken(token: string) {
		this.token = token;
		this.cookieService.set(
			'the_token',
			this.token,
			3650,
			'/',
			'',
			this.localtron.config.env.production ? true : false
		);
	}

	removeToken() {
		this.cookieService.delete(
			'the_token',
			'/',
			'',
			this.localtron.config.env.production ? true : false
		);
	}

	hasToken(): boolean {
		const t = this.cookieService.get('the_token');
		return !!t;
	}

	login(email: string, password: string): Promise<LoginResponse> {
		return this.localtron.call('/user/login', {
			email: email,
			password: password,
		});
	}

	readUserByToken(token: string): Promise<ReadUserByTokenResponse> {
		return this.localtron.call('/user/read-user-by-token', {
			token: token,
		});
	}

	getUsers(request: GetUsersRequest): Promise<GetUsersResponse> {
		return this.localtron.call('/user/get-users', request);
	}

	/** Save profile on behalf of a user */
	saveProfile(email: string, name: string): Promise<SaveProfileResponse> {
		const request: SaveProfileRequest = {
			email: email,
			name: name,
		};
		return this.localtron.call('/user/save-profile', request);
	}

	changePassword(
		email: string,
		currentPassword: string,
		newPassword: string
	): Promise<ChangePasswordResponse> {
		const request: ChangePasswordRequest = {
			email: email,
			currentPassword: currentPassword,
			newPassword: newPassword,
		};
		return this.localtron.call('/user/change-password', request);
	}

	changePasswordAdmin(
		email: string,
		newPassword: string
	): Promise<ChangePasswordAdminResponse> {
		const request: ChangePasswordAdminRequest = {
			email: email,
			newPassword: newPassword,
		};
		return this.localtron.call('/user/change-password-admin', request);
	}

	/** Create a user - alternative to registration
	 */
	createUser(
		user: User,
		password: string,
		roleIds: string[]
	): Promise<CreateUserResponse> {
		const request: CreateUserRequest = {
			user: user,
			password: password,
			roleIds: roleIds,
		};
		return this.localtron.call('/user/create-user', request);
	}

	getRoles(): Promise<GetRolesResposne> {
		return this.localtron.call('/user/get-roles', {});
	}

	getPermissions(): Promise<GetPermissionsResposne> {
		return this.localtron.call('/user/get-permissions', {});
	}

	setRolePermissions(
		roleId: string,
		permissionIds: string[]
	): Promise<SetRolePermissionsResponse> {
		const request: SetRolePermissionsRequest = {
			roleId: roleId,
			permissionIds: permissionIds,
		};
		return this.localtron.call('/user/set-role-permissions', request);
	}

	deleteRole(roleId: string): Promise<DeleteRoleResponse> {
		const request: DeleteRoleRequest = {
			roleId: roleId,
		};
		return this.localtron.call('/user/delete-role', request);
	}

	deleteUser(userId: string): Promise<DeleteUserResponse> {
		const request: DeleteUserRequest = {
			userId: userId,
		};
		return this.localtron.call('/user/delete-user', request);
	}

	async getUserId(): Promise<string> {
		try {
			const user = await firstValueFrom(this.user$.pipe(first()));
			return user.id as string;
		} catch (error) {
			console.error('Error getting user ID:', error);
			throw error;
		}
	}

	getUser(id: string): Observable<User | undefined> {
		if (!this.userCache[id]) {
			this.userCache[id] = new BehaviorSubject<User | undefined>(undefined);
			this.getUsers({
				query: {
					conditions: [equal(field('id'), id)],
				},
			}).then((rsp) => {
				this.userCache[id].next(rsp.users[0]);
			});
		}

		return this.userCache[id].asObservable();
	}
}

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
	id: 'role.admin',
	name: 'Admin Role',
	permissionIds: [],
};

export const RoleUser: Role = {
	id: 'role.user',
	name: 'User Role',
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

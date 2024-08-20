/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
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
import { equal, field } from '@singulatron/types';
import * as user from '@singulatron/types';
import {
	Configuration,
	UserSvcApi,
	UserSvcGetPermissionsResponse,
	UserSvcGetRolesResponse,
	UserSvcGetUsersResponse,
	UserSvcLoginResponse,
	UserSvcReadUserByTokenResponse,
	UserSvcUser,
} from '@singulatron/client';

@Injectable({
	providedIn: 'root',
})
export class UserService {
	private userService!: UserSvcApi;

	private token: string = '';

	private userSubject = new ReplaySubject<UserSvcUser>(1);
	/** Current logged in user */
	public user$ = this.userSubject.asObservable();

	private userCache: {
		[id: string]: BehaviorSubject<UserSvcUser | undefined>;
	} = {};

	constructor(
		private localtron: LocaltronService,
		private cookieService: CookieService,
		private router: Router
	) {
		this.userService = new UserSvcApi(
			new Configuration({
				basePath: this.localtron.addr(),
				apiKey: this.localtron.token(),
			})
		);
		this.init();
	}

	noop() {}

	async init() {
		this.getToken();

		if (!this.hasToken()) {
			try {
				const rsp = await this.login('singulatron', 'changeme');
				this.setToken(rsp.token!.token as string);
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
			this.userSubject.next(rsp.user!);
		} catch (error) {
			console.error('Cannot read user even with a token', error);
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

	login(slug: string, password: string): Promise<UserSvcLoginResponse> {
		return this.userService.login({
			request: { slug: slug, password: password },
		});
	}

	readUserByToken(token: string): Promise<UserSvcReadUserByTokenResponse> {
		return this.userService.readUserByToken({
			body: {
				token: token,
			},
		});
	}

	getUsers(request: user.GetUsersRequest): Promise<UserSvcGetUsersResponse> {
		return this.userService.getUsers({
			request: request,
		});
	}

	/** Save profile on behalf of a user */
	saveProfile(email: string, name: string): Promise<object> {
		const request: user.SaveProfileRequest = {
			email: email,
			name: name,
		};
		return this.userService.saveUserProfile({
			userId: '@thisIsFakeYetItWorks',
			body: request,
		});
	}

	changePassword(
		email: string,
		currentPassword: string,
		newPassword: string
	): Promise<object> {
		const request: user.ChangePasswordRequest = {
			email: email,
			currentPassword: currentPassword,
			newPassword: newPassword,
		};
		return this.userService.changePassword({
			request: request,
		});
	}

	changePasswordAdmin(
		email: string,
		newPassword: string
	): Promise<user.ChangePasswordAdminResponse> {
		const request: user.ChangePasswordAdminRequest = {
			email: email,
			newPassword: newPassword,
		};
		return this.userService.changePasswordAdmin({
			request: request,
		});
	}

	/** Create a user - alternative to registration
	 */
	createUser(
		user: UserSvcUser,
		password: string,
		roleIds: string[]
	): Promise<user.CreateUserResponse> {
		return this.userService.createUser({
			request: {
				user: user,
				password: password,
				roleIds: roleIds,
			},
		});
	}

	getRoles(): Promise<UserSvcGetRolesResponse> {
		return this.userService.getRoles();
	}

	getPermissions(roleId: string): Promise<UserSvcGetPermissionsResponse> {
		return this.userService.getPermissionsByRole({
			roleId: roleId,
		});
	}

	setRolePermissions(roleId: string, permissionIds: string[]): Promise<object> {
		const request: user.SetRolePermissionsRequest = {
			permissionIds: permissionIds,
		};
		return this.userService.setRolePermission({
			roleId: roleId,
			body: request,
		});
	}

	deleteRole(roleId: string): Promise<user.DeleteRoleResponse> {
		return this.userService.deleteRole({
			roleId: roleId,
		});
	}

	deleteUser(userId: string): Promise<user.DeleteUserResponse> {
		return this.userService.deleteUser({
			userId: userId,
		});
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

	getUser(id: string): Observable<UserSvcUser | undefined> {
		if (!this.userCache[id]) {
			this.userCache[id] = new BehaviorSubject<UserSvcUser | undefined>(
				undefined
			);
			this.getUsers({
				query: {
					conditions: [equal(field('id'), id)],
				},
			}).then((rsp) => {
				this.userCache[id].next(rsp.users![0]);
			});
		}

		return this.userCache[id].asObservable();
	}
}

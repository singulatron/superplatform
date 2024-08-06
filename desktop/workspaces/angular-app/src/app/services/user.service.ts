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

@Injectable({
	providedIn: 'root',
})
export class UserService {
	private token: string = '';

	private userSubject = new ReplaySubject<user.User>(1);
	/** Current logged in user */
	public user$ = this.userSubject.asObservable();

	private userCache: { [id: string]: BehaviorSubject<user.User | undefined> } =
		{};

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

	login(email: string, password: string): Promise<user.LoginResponse> {
		return this.localtron.post('/user-svc/login', {
			email: email,
			password: password,
		});
	}

	readUserByToken(token: string): Promise<user.ReadUserByTokenResponse> {
		return this.localtron.post('/user-svc/user/by-token', {
			token: token,
		});
	}

	getUsers(request: user.GetUsersRequest): Promise<user.GetUsersResponse> {
		return this.localtron.post('/user-svc/users', request);
	}

	/** Save profile on behalf of a user */
	saveProfile(email: string, name: string): Promise<user.SaveProfileResponse> {
		const request: user.SaveProfileRequest = {
			email: email,
			name: name,
		};
		return this.localtron.put(`/user-svc/user/${email}`, request);
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
		return this.localtron.post('/user-serviec/change-password', request);
	}

	changePasswordAdmin(
		email: string,
		newPassword: string
	): Promise<user.ChangePasswordAdminResponse> {
		const request: user.ChangePasswordAdminRequest = {
			email: email,
			newPassword: newPassword,
		};
		return this.localtron.post('/user-svc/change-password-admin', request);
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
		return this.localtron.post('/user-svc/user', request);
	}

	getRoles(): Promise<user.GetRolesResposne> {
		return this.localtron.get('/user-svc/roles');
	}

	getPermissions(roleId: string): Promise<user.GetPermissionsResposne> {
		return this.localtron.get(`/user-svc/role/${roleId}/permissions`);
	}

	setRolePermissions(
		roleId: string,
		permissionIds: string[]
	): Promise<user.SetRolePermissionsResponse> {
		const request: user.SetRolePermissionsRequest = {
			permissionIds: permissionIds,
		};
		return this.localtron.put(
			`/user-svc/role/${roleId}/permissions`,
			request
		);
	}

	deleteRole(roleId: string): Promise<user.DeleteRoleResponse> {
		return this.localtron.delete(`/user-svc/role/${roleId}`);
	}

	deleteUser(userId: string): Promise<user.DeleteUserResponse> {
		return this.localtron.delete(`/user-svc/user/${userId}`);
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

	getUser(id: string): Observable<user.User | undefined> {
		if (!this.userCache[id]) {
			this.userCache[id] = new BehaviorSubject<user.User | undefined>(
				undefined
			);
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

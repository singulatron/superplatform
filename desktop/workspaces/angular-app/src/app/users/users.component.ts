/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component } from '@angular/core';
import {
	FormBuilder,
	FormGroup,
	Validators,
	FormsModule,
	ReactiveFormsModule,
} from '@angular/forms';
import { UserService } from '../services/user.service';
import { User, GetUsersRequest } from '@singulatron/types';
import { first } from 'rxjs';
import { ToastController, IonicModule } from '@ionic/angular';
import { TranslatePipe } from '../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { NgFor, NgIf } from '@angular/common';
import { CenteredComponent } from '../components/centered/centered.component';
import { ChangeDetectorRef, ChangeDetectionStrategy } from '@angular/core';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { Router, ActivatedRoute } from '@angular/router';
import { QueryParser } from '../services/query.service';
import {
	fields,
	conditionsToKeyValue,
	contains,
	Condition,
	conditionFieldIs,
} from '@singulatron/types';

interface UserVisible extends User {
	visible?: boolean;
}

@Component({
	selector: 'app-users',
	templateUrl: './users.component.html',
	styleUrls: ['./users.component.scss'],
	standalone: true,
	imports: [
		PageComponent,
		IconMenuComponent,
		CenteredComponent,
		IonicModule,
		NgFor,
		FormsModule,
		ReactiveFormsModule,
		NgIf,
		TranslateModule,
		TranslatePipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UsersComponent {
	users: UserVisible[] = [];
	after: any;
	private userForms: Map<string, FormGroup> = new Map();

	count = 0;
	searchTerm = '';
	queryParser: QueryParser;

	constructor(
		private fb: FormBuilder,
		private userService: UserService,
		private toast: ToastController,
		private cd: ChangeDetectorRef,
		private router: Router,
		private activatedRoute: ActivatedRoute
	) {
		this.userForms = new Map();

		this.queryParser = new QueryParser();
		this.queryParser.defaultConditionFunc = (value: any): Condition => {
			return contains(fields(['name', 'email']), value);
		};

		this.userService.user$.pipe(first()).subscribe(() => {
			this.initializeOnLogin();
		});
	}

	private async initializeOnLogin() {
		this.activatedRoute.queryParams.subscribe(async (parameters) => {
			this.searchTerm =
				this.queryParser.convertQueryParamsToSearchTerm(parameters);

			await this.fetchUsers();
			this.cd.markForCheck();

			this.userForms?.clear();

			for (const user of this.users) {
				this.userForms.set(user.id!, this.createUserForm(user));
			}
		});
	}

	public redirect() {
		const query = this.queryParser.parse(this.searchTerm);

		const kv = conditionsToKeyValue(
			query.conditions
				? query.conditions.filter((v) => {
						return (
							!conditionFieldIs(v, 'name') && !conditionFieldIs(v, 'email')
						);
					})
				: []
		);

		if (Object.keys(kv)?.length) {
			this.router.navigate([], {
				queryParams: kv,
			});
			return;
		}

		if (this.searchTerm) {
			this.router.navigate([], {
				queryParams: { search: this.searchTerm },
			});
			return;
		}

		this.router.navigate([], {
			queryParams: {},
		});
	}

	public async fetchUsers() {
		const query = this.queryParser.parse(this.searchTerm);
		query.count = true;
		query.conditions = query.conditions || [];

		const request: GetUsersRequest = {
			query: query,
		};

		if (this.after) {
			request.query!.after = [this.after];
		}

		const response = await this.userService.getUsers(request);

		// eslint-disable-next-line
		if (response.users && this.after) {
			this.users = [...this.users, ...response.users];
		} else if (response.users) {
			this.users = response.users;
		} else {
			this.users = [];
		}

		this.count = response.count || 0;

		// eslint-disable-next-line
		if (response.after && response.after != `0001-01-01T00:00:00Z`) {
			this.after = response.after;
		} else {
			this.after = undefined;
		}

		this.cd.markForCheck();
	}

	createUserForm(user: UserVisible): FormGroup {
		return this.fb.group({
			name: [user.name, Validators.required],
			email: [user.email, [Validators.required]],
			password: [''],
			passwordConfirmation: [''],
			createdAt: [{ value: user.createdAt, disabled: true }],
			updatedAt: [{ value: user.updatedAt, disabled: true }],
			visible: [user.visible || false],
		});
	}

	getUserForm(userId: string): FormGroup {
		return this.userForms.get(userId)!;
	}

	async saveUser(userId: string) {
		const userForm = this.getUserForm(userId);
		if (userForm.invalid) {
			return;
		}

		const { name, email, password, passwordConfirmation } = userForm.value;

		if (password && password !== passwordConfirmation) {
			const toast = await this.toast.create({
				message: 'Passwords do not match',
				duration: 5000,
				color: 'danger',
				cssClass: 'white-text',
				position: 'middle',
			});
			toast.present();
			return;
		}

		try {
			let toastMessage = `Profile ${name} saved`;
			await this.userService.saveProfile(email, name);

			if (password) {
				toastMessage += ' and password changed';
				await this.userService.changePasswordAdmin(email, password);
			}

			const toast = await this.toast.create({
				color: 'secondary',
				message: toastMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();

			this.initializeOnLogin();
		} catch (error) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(error as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
				cssClass: 'white-text',
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}

	async deleteUser($event: any, userId: string) {
		$event.stopPropagation();

		try {
			await this.userService.deleteUser(userId);

			const toastMessage = `User ${name} deleted`;

			const toast = await this.toast.create({
				color: 'secondary',
				message: toastMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();

			this.initializeOnLogin();
		} catch (error) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(error as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
				cssClass: 'white-text',
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}

	trackById(_: number, message: { id?: string }): string {
		return message.id || '';
	}

	async loadMoreData() {
		if (!this.after) {
			console.log('No more users to load');
			return;
		}
		await this.fetchUsers();
	}
}

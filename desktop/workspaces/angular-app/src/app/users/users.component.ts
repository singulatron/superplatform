/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
 */
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormArray, Validators, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { User, UserService } from '../services/user.service';
import { first } from 'rxjs';
import { ToastController, IonicModule } from '@ionic/angular';
import { TranslatePipe } from '../../../shared/stdlib/translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { NgFor, NgIf } from '@angular/common';
import { CenteredComponent } from '../../../shared/stdlib/components/centered/centered.component';
import { SidebarPageComponent } from '../../../shared/stdlib/components/sidebar-page/sidebar-page.component';

interface UserVisible extends User {
	visible?: boolean;
}

@Component({
    selector: 'app-users',
    templateUrl: './users.component.html',
    styleUrls: ['./users.component.scss'],
    standalone: true,
    imports: [
        SidebarPageComponent,
        CenteredComponent,
        IonicModule,
        NgFor,
        FormsModule,
        ReactiveFormsModule,
        NgIf,
        TranslateModule,
        TranslatePipe,
    ],
})
export class UsersComponent implements OnInit {
	users: UserVisible[] = [];
	userForms: FormArray;

	constructor(
		private fb: FormBuilder,
		private userService: UserService,
		private toast: ToastController
	) {
		this.userForms = this.fb.array([]);
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	ngOnInit() {
		this.loggedInInit();
	}

	async loggedInInit() {
		let rsp = await this.userService.getUsers();
		this.users = rsp.users;
		this.users.forEach((user) => {
			this.userForms.push(this.createUserForm(user));
		});
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

	getUserForm(index: number): FormGroup {
		return this.userForms.at(index) as FormGroup;
	}

	async saveUser(index: number) {
		const userForm = this.getUserForm(index);
		if (userForm.invalid) {
			return;
		}

		const { name, email, password, passwordConfirmation } = userForm.value;

		if (password && password !== passwordConfirmation) {
			const toast = await this.toast.create({
				message: 'Passwords do not match',
				duration: 5000,
				color: 'danger',
				position: 'middle',
			});
			toast.present();
			return;
		}

		try {
			let toastMsg = `Profile ${name} saved`;
			await this.userService.saveProfile(email, name);

			if (password) {
				toastMsg += ' and password changed';
				await this.userService.changePassword(email, '', password);
			}

			const toast = await this.toast.create({
				color: 'secondary',
				message: toastMsg,
				duration: 5000,
				position: 'middle',
			});
			toast.present();

			this.loggedInInit();
		} catch (err) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(err as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
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

			let toastMsg = `User ${name} deleted`;

			const toast = await this.toast.create({
				color: 'secondary',
				message: toastMsg,
				duration: 5000,
				position: 'middle',
			});
			toast.present();

			this.loggedInInit();
		} catch (err) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(err as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				color: 'danger',
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}
}

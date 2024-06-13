import { Component } from '@angular/core';
import { User, UserService, Role } from '../../services/user.service';
import { first } from 'rxjs';
import { ToastController } from '@ionic/angular';

@Component({
	selector: 'app-add-user',
	templateUrl: './add-user.component.html',
	styleUrl: './add-user.component.scss',
})
export class AddUserComponent {
	password = '';
	passwordConfirmation = '';
	user: User = {};
	roles: Role[] = [];

	constructor(
		private userService: UserService,
		private toast: ToastController
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		let rsp = await this.userService.getRoles();
		this.roles = rsp.roles;
	}

	async createUser() {
		try {
			await this.userService.createUser(
				this.user,
				this.password,
				this.roles.map((v) => v.id as string)
			);
			const toast = await this.toast.create({
				message: 'User saved',
				duration: 5000,
				position: 'middle',
			});
			toast.present();
			this.user = {};
		} catch (err) {
			let errorMessage = 'An unexpected error occurred';
			try {
				errorMessage = (JSON.parse(err as any) as any)?.error;
			} catch {}

			const toast = await this.toast.create({
				message: errorMessage,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
		}
	}
}

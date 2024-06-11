import { Component } from '@angular/core';
import { User, UserService } from '../services/user.service';
import { first } from 'rxjs';
import { ToastController } from '@ionic/angular';

interface UserVisible extends User {
	visible?: boolean;
}

@Component({
	selector: 'app-users',
	templateUrl: './users.component.html',
	styleUrl: './users.component.scss',
})
export class UsersComponent {
	users: UserVisible[] = [];
	password = '';

	constructor(
		private userService: UserService,
		private toast: ToastController
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		let rsp = await this.userService.getUsers();
		this.users = rsp.users;
	}

	async saveUser(user: User) {
		await this.userService.saveProfile(
			user.email as string,
			user.name as string
		);
		let toast = await this.toast.create({
			message: 'Profile saved',
			duration: 5000,
			position: 'middle',
		});
		toast.present();
		if (!this.password) {
			return;
		}
		await this.userService.changePassword(user.email as string, '', this.password);
		toast = await this.toast.create({
			message: 'Password changed',
			duration: 5000,
			position: 'middle',
		});
		toast.present();
		return;
	}
}

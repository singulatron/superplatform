import { Component } from '@angular/core';
import { User, UserService } from '../services/user.service';
import { first } from 'rxjs';

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

	constructor(private userService: UserService) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		let rsp = await this.userService.getUsers();
		this.users = rsp.users;
	}

	async saveUser(user: User) {
		console.log(user);
	}
}

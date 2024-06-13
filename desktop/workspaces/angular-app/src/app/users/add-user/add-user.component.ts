import { Component } from '@angular/core';
import { User } from '../../services/user.service';

@Component({
	selector: 'app-add-user',
	templateUrl: './add-user.component.html',
	styleUrl: './add-user.component.scss',
})
export class AddUserComponent {
	password = '';
	passwordConfirmation = ''
	user: User = {};

	saveUser() {}
}

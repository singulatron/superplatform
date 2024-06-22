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
import { Component } from '@angular/core';
import { UserService, LoginResponse } from '../services/user.service';
import { ToastController, IonicModule } from '@ionic/angular';
import { FormsModule } from '@angular/forms';
import { NgIf } from '@angular/common';
import { CenteredComponent } from '../../../shared/stdlib/components/centered/centered.component';
import { SidebarPageComponent } from '../../../shared/stdlib/components/sidebar-page/sidebar-page.component';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
    styleUrl: './login.component.css',
    standalone: true,
    imports: [
        SidebarPageComponent,
        CenteredComponent,
        IonicModule,
        NgIf,
        FormsModule,
    ],
})
export class LoginComponent {
	email: string = '';
	name: string = '';
	password: string = '';
	passwordConfirmation: string = '';
	loginButtonDisabled: boolean = false;
	registerButtonDisabled: boolean = false;
	selectedSegment: string = 'login';

	constructor(
		private userService: UserService,
		private toast: ToastController
	) {}

	async login() {
		this.loginButtonDisabled = true;
		let rsp: LoginResponse;
		try {
			rsp = await this.userService.login(this.email, this.password);
		} catch (error) {
			const toast = await this.toast.create({
				message: (error as any)?.error,
				duration: 5000,
				position: 'middle',
			});
			toast.present();
			return;
		} finally {
			this.loginButtonDisabled = false;
		}
		if (!rsp?.token.token) {
			const toast = await this.toast.create({
				message: 'Login failure: no token in response',
				duration: 5000,
				position: 'middle',
			});
			toast.present();
			return;
		}

		this.userService.setToken(rsp?.token.token);
		window.location.href = '/';
	}
}

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
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

import { IonicModule } from '@ionic/angular';
import { TranslatePipe } from '../../translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { NgFor, NgIf } from '@angular/common';
import { CenteredComponent } from '../../components/centered/centered.component';
import { ChangeDetectionStrategy } from '@angular/core';
import { PageComponent } from '../../components/page/page.component';
import { IconMenuComponent } from '../../components/icon-menu/icon-menu.component';
import { UserService } from '../../services/user.service';

@Component({
	selector: 'app-logout',
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
	templateUrl: './logout.component.html',
	styleUrl: './logout.component.scss',
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LogoutComponent {
	constructor(private userService: UserService) {}

	logout() {
		this.userService.removeToken();
		window.location.pathname = '/';
	}
}

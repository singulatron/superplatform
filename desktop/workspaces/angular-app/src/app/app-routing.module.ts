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
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { StartupComponent } from './startup/startup.component';
import { ChatComponent } from './chat/chat.component';
import { ModelExplorerComponent } from './model-explorer/model-explorer.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { UsersComponent } from './users/users.component';
import { AddUserComponent } from './users/add-user/add-user.component';
import { RolesComponent } from './users/roles/roles.component';

const routes: Routes = [
	{
		path: '',
		component: HomeComponent,
	},
	{
		path: 'startup',
		component: StartupComponent,
	},
	{
		path: 'chat',
		component: ChatComponent,
	},
	{
		path: 'model-explorer',
		component: ModelExplorerComponent,
	},
	{
		path: 'users',
		component: UsersComponent,
	},
	{
		path: 'add-user',
		component: AddUserComponent,
	},
	{
		path: 'roles',
		component: RolesComponent,
	},
	{
		path: 'login',
		component: LoginComponent,
	},
];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule],
})
export class AppRoutingModule {}

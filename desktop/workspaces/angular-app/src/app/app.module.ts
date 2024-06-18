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
import {
	HttpClient,
	provideHttpClient,
	withInterceptorsFromDi,
} from '@angular/common/http';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { TranslateLoader, TranslateModule } from '@ngx-translate/core';
import { TranslateHttpLoader } from '@ngx-translate/http-loader';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { IonicModule } from '@ionic/angular';
import { StartupComponent } from './startup/startup.component';
import { StdlibModule } from '../../shared/stdlib/stdlib.module';
import { environment } from '../environments/environment';
import { AiModule } from '../../shared/ai/ai.module';
import { MarkdownModule } from 'ngx-markdown';
import { ChatComponent } from './chat/chat.component';
import { ModelExplorerComponent } from './model-explorer/model-explorer.component';
import { DownloadingComponent } from './downloading/downloading.component';
import { AdvancedModelExplorerComponent } from './model-explorer/advanced-model-explorer/advanced-model-explorer.component';
import { DefaultModelExplorerComponent } from './model-explorer/default-model-explorer/default-model-explorer.component';
import {
	LOCALTRON_SERVICE_CONFIG,
	LocaltronService,
} from './services/localtron.service';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { UsersComponent } from './users/users.component';
import { AddUserComponent } from './users/add-user/add-user.component';
import { RolesComponent } from './users/roles/roles.component';

// AoT requires an exported function for factories
export function HttpLoaderFactory(http: HttpClient): TranslateHttpLoader {
	return new TranslateHttpLoader(http, './assets/i18n/', '.json');
}

@NgModule({
	declarations: [
		AppComponent,
		StartupComponent,
		ChatComponent,
		ModelExplorerComponent,
		AdvancedModelExplorerComponent,
		DefaultModelExplorerComponent,
		DownloadingComponent,
		HomeComponent,
		LoginComponent,
		UsersComponent,
		AddUserComponent,
		RolesComponent,
	],
	schemas: [CUSTOM_ELEMENTS_SCHEMA],
	bootstrap: [AppComponent],
	imports: [
		BrowserModule,
		BrowserAnimationsModule,
		AppRoutingModule,
		ReactiveFormsModule,
		IonicModule.forRoot({
			// force Android mode across all platforms
			mode: 'md',
		}),
		TranslateModule.forRoot({
			defaultLanguage: 'en',
			loader: {
				provide: TranslateLoader,
				useFactory: HttpLoaderFactory,
				deps: [HttpClient],
			},
		}),
		StdlibModule.forRoot({
			apiServiceConfig: {
				env: environment,
			},
		}),
		AiModule,
		MarkdownModule.forRoot(),
	],
	providers: [
		{
			provide: LOCALTRON_SERVICE_CONFIG,
			useValue: { env: environment },
		},
		LocaltronService,
		provideHttpClient(withInterceptorsFromDi()),
	],
})
export class AppModule {}

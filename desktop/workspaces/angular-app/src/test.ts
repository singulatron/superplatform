import { getTestBed } from '@angular/core/testing';
import { platformBrowserDynamicTesting } from '@angular/platform-browser-dynamic/testing';

import { NgModule } from '@angular/core';
import { BrowserDynamicTestingModule } from '@angular/platform-browser-dynamic/testing';
import {
	provideExperimentalZonelessChangeDetection,
	importProvidersFrom,
} from '@angular/core';

import { HttpLoaderFactory } from './app/app.module';
import { environment } from './environments/environment';
import { MarkdownModule } from 'ngx-markdown';
import { AiModule } from './app/ai.module';
import { StdlibModule } from './app/stdlib.module';
import { TranslateModule, TranslateLoader } from '@ngx-translate/core';
import { IonicModule } from '@ionic/angular';
import { ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app/app-routing.module';
import { provideAnimations } from '@angular/platform-browser/animations';
import { BrowserModule } from '@angular/platform-browser';
import {
	provideHttpClient,
	withInterceptorsFromDi,
	HttpClient,
} from '@angular/common/http';
import {
	LOCALTRON_SERVICE_CONFIG,
	LocaltronService,
} from './app/services/localtron.service';
import { MobileService } from './app/services/mobile.service';
import { FooterService } from './app/services/footer.service';

@NgModule({
	imports: [BrowserDynamicTestingModule],
	providers: [
		provideExperimentalZonelessChangeDetection(),
		MobileService,
		FooterService,
		importProvidersFrom(
			BrowserModule,
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
			MarkdownModule.forRoot()
		),
		{
			provide: LOCALTRON_SERVICE_CONFIG,
			useValue: { env: environment },
		},
		LocaltronService,
		provideHttpClient(withInterceptorsFromDi()),
		provideAnimations(),
	],
})
export class CustomTestModule {}

getTestBed().initTestEnvironment(
	CustomTestModule,
	platformBrowserDynamicTesting(),
	{
		teardown: { destroyAfterEach: false },
	}
);

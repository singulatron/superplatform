/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable, Inject, InjectionToken } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { firstValueFrom, map, throwError } from 'rxjs';
import { CookieService } from 'ngx-cookie-service';
import { catchError } from 'rxjs/operators';

export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
	backendAddress: string;
	localPromptAddress: string;
	localtronAddress: string;
}

export interface LocaltronServiceConfig {
	env: Environment;
}
export const LOCALTRON_SERVICE_CONFIG =
	new InjectionToken<LocaltronServiceConfig>('LocaltronServiceConfig');

@Injectable({
	providedIn: 'root',
})
export class LocaltronService {
	private headers: HttpHeaders;
	public config: LocaltronServiceConfig;

	constructor(
		private http: HttpClient,
		private cs: CookieService,
		@Inject(LOCALTRON_SERVICE_CONFIG) config: LocaltronServiceConfig
	) {
		this.config = config;
		this.headers = new HttpHeaders();
	}

	call(path: string, request: any): Promise<any> {
		if (!this.config.env.localtronAddress) {
			console.log('Localtron address is not set', {
				config: this.config,
			});
			throw 'Localtron address seems to be empty';
		}

		const uri = this.config.env.localtronAddress + path;

		const body = JSON.stringify(request);

		// @todo get this from the user service - import cycle currently
		const headers = this.headers.set(
			'Authorization',
			'Bearer ' + this.cs.get('the_token')
		);

		return firstValueFrom(
			this.http
				.post<any>(uri, body, {
					headers: headers,
					responseType: 'text' as 'json',
				})
				.pipe(
					map((response) => {
						return JSON.parse(response);
					}),
					catchError((error) => {
						if (error.status >= 400) {
							throw error.error;
						}
						return throwError(error);
					})
				)
		);
	}

	uuid() {
		return (
			generateSegment(8) +
			'-' +
			generateSegment(4) +
			'-' +
			generateSegment(4) +
			'-' +
			generateSegment(4) +
			'-' +
			generateSegment(12)
		);
	}
}

function generateSegment(length: number) {
	return Array.from({ length: length }, () =>
		Math.floor(Math.random() * 16).toString(16)
	).join('');
}

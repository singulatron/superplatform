/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable, Inject, InjectionToken } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { firstValueFrom, map } from 'rxjs';
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

	private prepareHeaders(): HttpHeaders {
		return this.headers.set(
			'Authorization',
			'Bearer ' + this.cs.get('the_token')
		);
	}

	private handleError(error: any): Promise<any> {
		if (error.status >= 400) {
			throw error.error;
		}
		return Promise.reject(error);
	}

	get(path: string): Promise<any> {
		const uri = this.config.env.localtronAddress + path;
		const headers = this.prepareHeaders();

		return firstValueFrom(
			this.http.get<any>(uri, { headers, responseType: 'text' as 'json' }).pipe(
				map((response) => JSON.parse(response)),
				catchError((error) => this.handleError(error))
			)
		);
	}

	post(path: string, request: any): Promise<any> {
		const uri = this.config.env.localtronAddress + path;
		const body = JSON.stringify(request);
		const headers = this.prepareHeaders();

		return firstValueFrom(
			this.http
				.post<any>(uri, body, { headers, responseType: 'text' as 'json' })
				.pipe(
					map((response) => JSON.parse(response)),
					catchError((error) => this.handleError(error))
				)
		);
	}

	put(path: string, request: any): Promise<any> {
		const uri = this.config.env.localtronAddress + path;
		const body = JSON.stringify(request);
		const headers = this.prepareHeaders();

		return firstValueFrom(
			this.http
				.put<any>(uri, body, { headers, responseType: 'text' as 'json' })
				.pipe(
					map((response) => JSON.parse(response)),
					catchError((error) => this.handleError(error))
				)
		);
	}

	delete(path: string): Promise<any> {
		const uri = this.config.env.localtronAddress + path;
		const headers = this.prepareHeaders();

		return firstValueFrom(
			this.http
				.delete<any>(uri, { headers, responseType: 'text' as 'json' })
				.pipe(
					map((response) => JSON.parse(response)),
					catchError((error) => this.handleError(error))
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

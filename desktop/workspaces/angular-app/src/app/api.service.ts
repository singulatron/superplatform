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
import { throwError } from 'rxjs';
import { BehaviorSubject } from 'rxjs';

export interface Event {
	type: string;
}

export interface FilterSaveEvent extends Event {
	type: 'filter-save';
	filters: string[];
}

export interface Environment {
	production: boolean;
	brandName: string;
	shortBrandName: string;
	backendAddress: string;
}

export interface ApiServiceConfig {
	env: Environment;
}
export const API_SERVICE_CONFIG = new InjectionToken<ApiServiceConfig>(
	'ApiServiceConfig'
);

export interface Thread {
	id: string;
	name: string;
	messages?: Array<Message>;
}

export const defaultThreadName = 'New chat';

export interface Message {
	id: string;
	content?: string;
	userId?: string;
}

@Injectable({
	providedIn: 'root',
})
export class ApiService {
	private locale = 'en';

	public firehose: BehaviorSubject<Event> = new BehaviorSubject<Event>({
		type: 'noop',
	});

	private headers: HttpHeaders;
	private config: ApiServiceConfig;

	constructor(
		private http: HttpClient,
		private cs: CookieService,
		@Inject(API_SERVICE_CONFIG) config: ApiServiceConfig
	) {
		this.headers = new HttpHeaders();
		this.config = config;
	}

	public setLocale(s: string) {
		this.locale = s;
	}

	public getLocale(): string {
		return this.locale;
	}

	call(path: string, request: any): Promise<any> {
		const uri = this.config.env.backendAddress + path;

		const body = JSON.stringify(request);

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

	getModelsFromServer(): Promise<any> {
		return this.call('/app/models', {});
	}

	getVersion(): Promise<VersionResponse> {
		return this.call('/app/version', {});
	}




}

export interface ReadByWebsitesRequest {
	host: string;
}

export interface VersionResponse {
	windows?: Version;
	linux?: Version;
	mac?: Version;
}

export interface Version {
	version?: string;
	downloadPageURL?: string;
	downloadURL?: string;
	releaseDate?: Date | string;
	changeLog?: string;
}

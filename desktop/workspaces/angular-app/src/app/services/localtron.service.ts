/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable, Inject, InjectionToken } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import Sonyflake from 'sonyflake';

const base62 = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
const sonyflake = new Sonyflake();

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
	public config: LocaltronServiceConfig;

	constructor(
		private cs: CookieService,
		@Inject(LOCALTRON_SERVICE_CONFIG) config: LocaltronServiceConfig
	) {
		this.config = config;
	}

	token(): string {
		return this.cs.get('the_token');
	}

	addr(): string {
		return this.config.env.localtronAddress;
	}

	id(prefix: string): string {
		const numberStr = sonyflake.nextId();
		const number = BigInt(numberStr);

		if (number === BigInt(0)) {
			return `${prefix}_0`;
		}

		let result = '';
		let num = number;

		while (num > 0) {
			const remainder = Number(num % BigInt(62));
			num = num / BigInt(62);
			result = base62[remainder] + result;
		}

		return `${prefix}_${result}`;
	}

	
}

function generateSegment(length: number) {
	return Array.from({ length: length }, () =>
		Math.floor(Math.random() * 16).toString(16)
	).join('');
}

/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service.js';
import { FirehoseService } from './firehose.service.js';
import { ReplaySubject } from 'rxjs';
import { UserService } from './user.service.js';
import { first } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class ConfigService {
	lastConfig!: Config;

	onConfigUpdateSubject = new ReplaySubject<Config>(1);
	/** Config emitted whenever it's loaded (on startup) or saved */
	onConfigUpdate$ = this.onConfigUpdateSubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.init();
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'configUpdate': {
					const rsp = await this.configGet();
					this.onConfigUpdateSubject.next(rsp.config);
					break;
				}
			}
		});
	}

	async loggedInInit() {
		try {
			const rsp = await this.configGet();
			this.lastConfig = rsp?.config;
			this.onConfigUpdateSubject.next(rsp?.config as Config);
		} catch (error) {
			console.error('Error in pollConfig', {
				error: JSON.stringify(error),
			});
		}
	}

	async configGet(): Promise<ConfigGetResponse> {
		return await this.localtron.call('/config/get', {});
	}
}

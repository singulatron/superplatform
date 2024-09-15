/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { FirehoseService } from './firehose.service';
import { ReplaySubject } from 'rxjs';
import { UserService } from './user.service';
import { first } from 'rxjs';
import {
	ConfigSvcApi,
	ConfigSvcGetConfigResponse,
	Configuration,
	ConfigSvcConfig as Config
} from '@singulatron/client';

@Injectable({
	providedIn: 'root',
})
export class ConfigService {
	private configService!: ConfigSvcApi;

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
			this.configService = new ConfigSvcApi(
				new Configuration({
					basePath: this.localtron.addr(),
					apiKey: this.localtron.token(),
				})
			);
			this.loggedInInit();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'configUpdate': {
					const rsp = await this.configGet();
					this.onConfigUpdateSubject.next(rsp.config!);
					break;
				}
			}
		});
	}

	async loggedInInit() {
		try {
			const rsp = await this.configGet();
			this.lastConfig = rsp?.config || {};
			this.onConfigUpdateSubject.next(rsp?.config as Config);
		} catch (error) {
			console.log(error);
			console.error('Error in pollConfig', {
				error: JSON.stringify(error),
			});
		}
	}

	async configGet(): Promise<ConfigSvcGetConfigResponse> {
		return await this.configService.getConfig();
	}
}

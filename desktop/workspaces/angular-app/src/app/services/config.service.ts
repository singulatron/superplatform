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
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { FirehoseService } from './firehose.service';
import { ReplaySubject } from 'rxjs';
import { UserService } from './user.service';
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
				case 'configUpdate':
					let rsp1 = await this.configGet();
					this.onConfigUpdateSubject.next(rsp1.config);
					break;
			}
		});
	}

	async loggedInInit() {
		try {
			let rsp = await this.configGet();
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

export interface Config {
	download?: {
		downloadFolder?: string;
	};

	model?: {
		currentModelId?: string;
	};

	/** This flag drives a minor UX feature:
	 * if the user has not installed the runtime we show an INSTALL
	 * button, but if the user has already installed the runtime we show
	 * we show a START runtime button.
	 * */
	isRuntimeInstalled?: boolean;
}

export interface ConfigGetRequest {}

export interface ConfigGetResponse {
	config: Config;
}

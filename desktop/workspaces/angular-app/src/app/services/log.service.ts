/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';

@Injectable({
	providedIn: 'root',
})
export class LogService {
	constructor(private localtron: LocaltronService) {}

	async logDisable(): Promise<void> {
		return this.localtron.call('/app/log/disable', {});
	}

	async logEnable(): Promise<void> {
		return this.localtron.call('/app/log/enable', {});
	}

	async logStatus(): Promise<LoggingStatus> {
		return this.localtron.call('/app/log/status', {});
	}
}

interface LoggingStatus {
	enabled: boolean;
}

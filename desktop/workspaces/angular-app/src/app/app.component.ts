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
import { Component } from '@angular/core';
import { ElectronIpcService } from './services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { Log } from '../../shared/backend-api/app';
import { RouterOutlet } from '@angular/router';

let loggingEnabled = true;

export function enableLogging() {
	loggingEnabled = true;
}

export function disableLogging() {
	loggingEnabled = false;
}

const originalConsole = {
	log: console.log,
	error: console.error,
	warn: console.warn,
	info: console.info,
	debug: console.debug,
	trace: console.trace,
};

function overrideConsole(ipcService: ElectronIpcService) {
	for (let mn of ['log', 'error', 'warn', 'info', 'debug', 'trace']) {
		const methodName: keyof Console = mn as any;

		console[methodName] = ((...args: any[]) => {
			if (!loggingEnabled) {
				return;
			}
			(originalConsole as any)[methodName](...args);
			try {
				let req: Log = {
					level: methodName,
					message: args[0],
					source: 'frontend',
					fields: args.length > 1 ? args[1] : {},
				};
				ipcService.send(WindowApiConst.LOG_REQUEST, req);
			} catch (err) {
				originalConsole.error('Cannot send log to IPC', err);
			}
		}) as any;
	};
}

@Component({
	selector: 'app-root',
	templateUrl: './app.component.html',
	styleUrls: ['./app.component.scss'],
	standalone: true,
	imports: [RouterOutlet],
})
export class AppComponent {
	title = 'singulatron-angular-app';

	constructor(private ipcService: ElectronIpcService) {
		overrideConsole(this.ipcService);
	}

	ngOnInit(): void {}
}

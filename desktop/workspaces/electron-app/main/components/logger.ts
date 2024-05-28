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
import * as fs from 'fs';
import * as util from 'util';
import { homedir } from 'os';
import { join } from 'path';
import { sendLogToLocalServer } from '../ipc-handlers/log-request';
import { Log } from '../../../angular-app/shared/backend-api/app';

let loggingEnabled = true;

export function disableLogging() {
	loggingEnabled = false;
}

export function enableLogging() {
	loggingEnabled = true;
}

// Define a single path for the log file
// @todo change this so its not fixed
const LOG_FILE_PATH = join(homedir(), '.singulatron', 'all.log');

const logStream = fs.createWriteStream(LOG_FILE_PATH, { flags: 'a' });

const originalConsole: { [key: string]: (...data: any[]) => void } = {
	log: console.log,
	error: console.error,
	warn: console.warn,
	info: console.info,
	debug: console.debug,
	trace: console.trace,
};

['log', 'error', 'warn', 'info', 'debug', 'trace'].forEach((mn) => {
	const methodName: keyof Console = mn as any;

	console[methodName] = ((...args: any[]) => {
		if (!loggingEnabled) {
			return;
		}

		let log: Log;

		if (args?.length == 0) {
			originalConsole.trace('no arguments for logging');
			return;
		}
		// local Nodejs logs are coming like
		// console.log("Log message", extraObject)
		// however Localtron logs come prepared as single object as first arg
		if (typeof args[0] === 'object') {
			log = args[0];
		} else {
			log = {
				level: mn,
				source: 'nodejs',
				message: args[0],
				fields: args[1],
			};
		}
		if (!log.time) {
			log.time = new Date().toISOString();
		}

		const localMessage = formatForLocalLogs(log);

		originalConsole[methodName as keyof Console](localMessage);
		sendLogToLocalServer(log);
		logStream.write(localMessage + '\n');
	}) as any;
});

function formatForLocalLogs(log: Log): string {
	return `${log.time} [${log.level?.toUpperCase()}] [${log.source}] ${log.message} ${log.fields && Object.keys(log.fields)?.length ? JSON.stringify(log.fields) : ''}`;
}

export const logger = {};

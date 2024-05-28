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
import sudo from 'sudo-prompt';
import { EventService } from '../services/event-service';
import { promises as fs } from 'fs';
import path from 'path';
import os from 'os';
import { join } from 'path';

const username = os.userInfo().username;

export class OSManager {
	logFilePath: string;
	tempScriptPath: string;

	constructor(
		private assetFolder: string,
		private eventService: EventService
	) {
		this.eventService.installRuntimeRequest$.subscribe(() => {
			this.initializeEnvironment();
		});

		// @todo fix path, config service is now in localtron
		this.logFilePath = path.join(os.homedir(), 'singulatron_install.log');

		this.init();

		new FileWatcher(this.logFilePath, (log) => {
			this.eventService.onRuntimeInstallLogSubject.next(log);
		});
	}

	async init() {
		try {
			let fd = await fs.open(this.logFilePath, 'w');
			fd.close();
		} catch (err) {
			console.error('Error opening log file:', err);
		}
	}

	async initializeEnvironment(): Promise<void> {
		const platform = os.platform();

		if (platform === 'win32') {
			await this.setupWindows();
		} else if (platform == 'linux' || platform == 'darwin') {
			await this.setupLinux();
		} else {
			console.log(`Running on ${platform}, no special setup required.`);
		}
	}

	async setupWindows(): Promise<void> {
		let exePath = join(this.assetFolder, 'dapper.exe');

		const command = `"${exePath}" --var-username=${username} --var-assetfolder=${this.assetFolder} run "${join(this.assetFolder, 'app.json')}" > "${this.logFilePath}" 2>&1`;
		await this.executeCommand(command);
	}

	async setupLinux(): Promise<void> {
		let exePath = join(this.assetFolder, 'dapper');

		const command = `"${exePath}" --var-username=${username} --var-assetfolder=${this.assetFolder} run "${join(this.assetFolder, 'app.json')}" > "${this.logFilePath}" 2>&1`;
		await this.executeCommand(command);
	}

	private async executeCommand(command: string): Promise<void> {
		return new Promise((resolve, reject) => {
			sudo.exec(
				command,
				{
					name: 'Singulatron Environment Setup',
					// @todo fix icon
					// icns: '/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns'
					icns: getIconPath(this.assetFolder),
				},
				(error, stdout, stderr) => {
					if (error) {
						console.error('Error:', error);
						reject(error);
					} else {
						console.log('stdout:', stdout);
						if (stderr) console.error('stderr:', stderr);
						resolve();
					}
				}
			);
		});
	}
}

function getIconPath(assetFolder: string): string {
	switch (os.platform()) {
		case 'darwin':
			return path.join(assetFolder, 'main', 'icons', 'icon.icns');
		case 'win32':
			return path.join(assetFolder, 'main', 'icons', 'icon.ico');
		case 'linux':
			return path.join(assetFolder, 'main', 'icons', 'icon.png');
		default:
			return '';
	}
}

class FileWatcher {
	filePath: string;
	lastContent: string;

	constructor(
		filePath: string,
		private callback: (s: string) => void
	) {
		this.filePath = filePath;
		this.lastContent = '';
		this.watchFile();
	}

	async readFile() {
		try {
			let data = await fs.readFile(this.filePath, 'utf8');

			if (data !== this.lastContent) {
				this.lastContent = data;
			}

			this.callback(this.lastContent);
		} catch (err) {
			console.error('Error reading file:', err);
		}
	}

	watchFile() {
		setInterval(() => this.readFile(), 500);
	}
}

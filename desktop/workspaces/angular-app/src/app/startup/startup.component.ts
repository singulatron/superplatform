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
import { Component, OnInit, ViewChild, ElementRef } from '@angular/core';
import { ElectronIpcService } from '../services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { ElectronAppService } from '../services/electron-app.service';
import { combineLatest, Subscription } from 'rxjs';
import { ApiService } from '../../../shared/stdlib/api.service';
import { DownloadService, DownloadDetails } from '../services/download.service';
import { ModelService } from '../services/model.service';
import { DockerService } from '../services/docker.service';
import { models } from '../../../shared/stdlib/api.service';
import { ConfigService, Config } from '../services/config.service';

@Component({
	selector: 'app-startup',
	templateUrl: './startup.component.html',
	styleUrl: './startup.component.scss',
})
export class StartupComponent implements OnInit {
	@ViewChild('logContainer') private logContainer!: ElementRef;
	log = 'Installation logs will be streamed here. Please wait...\n';
	scrollToBottom(): void {
		try {
			this.logContainer.nativeElement.scrollTop =
				this.logContainer.nativeElement.scrollHeight;
		} catch (err) {}
	}

	models = models;
	allIsWell = false;
	isDownloading = false;
	downloaded = false;
	restartIsRequired = false;

	downloadFolder: string = '';

	showSections = {
		model: false,
		dependencies: false,
		starting: false,
	};

	toggleSection(section: string) {
		type ShowSectionsKeys = keyof typeof this.showSections;
		this.showSections[section as ShowSectionsKeys] =
			!this.showSections[section as ShowSectionsKeys];
	}

	constructor(
		private ipcService: ElectronIpcService,
		public lapi: ElectronAppService,
		public configService: ConfigService,
		public downloadService: DownloadService,
		public dockerService: DockerService,
		public modelService: ModelService,
		private apiService: ApiService
	) {}

	handleDownloadStatus(data: DownloadDetails) {
		this.isDownloading = data.status == 'inProgress' || data.status == 'paused';
		this.downloaded = data.status == 'completed';
	}

	selectedModelName(cu: Config) {
		let mod = this.models?.find((v) => v.id == cu?.model?.currentModelId);
		let displayName = [mod?.name, mod?.flavour, mod?.version].join(' ');
		return displayName;
	}

	selectedModel(cu: Config) {
		return this.models?.find((v) => v.id == cu?.model?.currentModelId);
	}

	ngOnDestroy() {
		this.subscriptions.forEach((v) => v.unsubscribe());
	}

	private subscriptions: Subscription[] = [];

	async ngOnInit(): Promise<void> {
		this.subscriptions.push(
			this.lapi.onRuntimeInstallLog$.subscribe((data) => {
				if (data == this.log) {
					return;
				}

				data
					.replace(this.log, '')
					.trim()
					.split('\n')
					.forEach((l) => {
						l = l?.trim();
						if (l) {
							console.log('Install log: ' + l);
						}
					});

				this.log = data;
				if (
					this.log?.includes('RESTART REQUIRED') ||
					this.log?.includes('restart is required')
				) {
					this.restartIsRequired = true;
				}

				this.scrollToBottom();
			})
		);
		this.models = await this.apiService.getModels();

		this.subscriptions.push(
			this.lapi.onFolderSelect$.subscribe((data) => {
				this.downloadFolder = data.location;
			})
		);

		let selectedExists = false;
		this.subscriptions.push(
			this.modelService.onModelCheck$.subscribe((data) => {
				if (data.selectedExists === undefined) {
					return;
				}
				if (data.selectedExists !== selectedExists) {
					selectedExists = data.selectedExists;
				}
			})
		);

		combineLatest([
			this.dockerService.onDockerInfo$,
			this.modelService.onModelCheck$,
		]).subscribe(([dockerInfo, modelCheck]) => {
			if (this.allIsWell) {
				return;
			}
			if (!dockerInfo.hasDocker) {
				this.showSections.dependencies = true;
			} else if (!modelCheck.selectedExists) {
				this.showSections.model = true;
			} else {
				this.showSections.starting = true;
			}
		});

		this.subscriptions.push(
			this.modelService.onModelLaunch$.subscribe(async () => {
				if (this.allIsWell) {
					return;
				}
				this.showSections.starting = false;
				console.log('Received model launch event');
				this.allIsWell = true;
			})
		);
	}

	openFolderSelect() {
		this.ipcService.send(WindowApiConst.SELECT_FOLDER_REQUEST, {});
	}

	async download() {
		const config = this.configService.lastConfig;
		if (!config?.model?.currentModelId) {
			throw 'Model id is empty';
		}
		this.downloadService.downloadDo(config?.model?.currentModelId);
	}

	isRuntimeInstalling = false;
	async installRuntime() {
		this.ipcService.send(WindowApiConst.INSTALL_RUNTIME_REQUEST, {});
		this.isRuntimeInstalling = true;
	}
}

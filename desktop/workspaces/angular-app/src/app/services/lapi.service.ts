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
import { ElectronIpcService } from './electron-ipc.service';
import { ReplaySubject, Observable, lastValueFrom } from 'rxjs';
import { take } from 'rxjs/operators';
import { combineLatest } from 'rxjs';

import { Config } from 'shared-lib/models/types';
import {
	OnDockerInfo,
	FileDownloadRequest,
	OnFileDownloadStatus,
	SelectFolderRequest,
	OnFolderSelect,
	GraphicsInfoRequest,
	OnGraphicsInfo,
	ModelLaunchRequest,
	OnModelLaunch,
	OnModelCheck,
	OnOSInfo,
	onModelReady,
	OnSystemLanguage,
	DockerImagePullStatus,
} from 'shared-lib/models/event-request-response';
import { WindowApiConst } from 'shared-lib';
import { ApiService } from '../../../shared/stdlib/api.service';

@Injectable({
	providedIn: 'root',
})
export class LapiService {
	lastConfig: Config;

	onDockerInfoSubject = new ReplaySubject<OnDockerInfo>(1);
	onDockerInfo$ = this.onDockerInfoSubject.asObservable();

	onGraphicsInfoSubject = new ReplaySubject<OnGraphicsInfo>(1);
	onGraphicsInfo$ = this.onGraphicsInfoSubject.asObservable();

	onOsInfoSubject = new ReplaySubject<OnOSInfo>(1);
	onOsInfo$ = this.onOsInfoSubject.asObservable();

	onFolderSelectSubject = new ReplaySubject<OnFolderSelect>(1);
	/** Emitted when a folder selection finished  */
	onFolderSelect$ = this.onFolderSelectSubject.asObservable();

	onConfigUpdateSubject = new ReplaySubject<Config>(1);
	/** Config emitted whenever it's loaded (on startup) or saved */
	onConfigUpdate$ = this.onConfigUpdateSubject.asObservable();

	onFileDownloadStatusSubject = new ReplaySubject<OnFileDownloadStatus>(1);
	onFileDownloadStatus$ = this.onFileDownloadStatusSubject.asObservable();

	onModelCheckSubject = new ReplaySubject<OnModelCheck>(1);
	/** Emitted any time when the currently selected model is checked */
	onModelCheck$ = this.onModelCheckSubject.asObservable();

	onModelLaunchSubject = new ReplaySubject<OnModelLaunch>(1);
	/** Emitted when the model is launched and available shortly */
	onModelLaunch$ = this.onModelLaunchSubject.asObservable();

	onModelReadySubject = new ReplaySubject<onModelReady>(1);
	onModelReady$ = this.onModelReadySubject.asObservable();

	onRuntimeInstallLogSubject = new ReplaySubject<string>(1);
	onRuntimeInstallLog$ = this.onRuntimeInstallLogSubject.asObservable();

	onLLMAddressChangeSubject = new ReplaySubject<string>(1);
	onLLMAddressChange$ = this.onLLMAddressChangeSubject.asObservable();

	onDockerImagePullStatusSubject = new ReplaySubject<DockerImagePullStatus>(1);
	onDockerImagePullStatus$ = this.onDockerImagePullStatusSubject.asObservable();

	constructor(
		private ipcService: ElectronIpcService,
		public apiService: ApiService
	) {
		this.listenToIpcEvents();
		this.listenToModelReady();
		this.ipcService.send(WindowApiConst.FRONTEND_READY_REQUEST, {});
	}

	private listenToIpcEvents(): void {
		this.ipcService.receive<OnFolderSelect>(
			WindowApiConst.ON_FOLDER_SELECT,
			(data) => {
				this.onFolderSelectSubject.next(data);
			}
		);

		this.ipcService.receive<OnSystemLanguage>(
			WindowApiConst.ON_SYSTEM_LANGUAGE,
			(data) => {
				this.apiService.setLocale(data.systemLanguage);
			}
		);

		this.ipcService.receive<string>(
			WindowApiConst.ON_RUNTIME_INSTALL_LOG,
			(data) => {
				this.onRuntimeInstallLogSubject.next(data);
			}
		);
	}

	private listenToModelReady(): void {
		combineLatest([this.onDockerInfo$, this.onModelCheck$]).subscribe(
			([dockerInfo, modelCheck]) => {
				if (dockerInfo.hasDocker && modelCheck.selectedExists) {
					this.onModelReadySubject.next({ modelReady: true });
				}
			}
		);
	}
}

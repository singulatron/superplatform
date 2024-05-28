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
import { ReplaySubject } from 'rxjs';
import { Config } from 'shared-lib/models/types';
import {
	OnDockerInfo,
	OnFileDownloadStatus,
	OnModelCheck,
	OnGraphicsInfo,
	OnOSInfo,
	OnModelLaunch,
	ModelLaunchRequest,
	DockerImagePullStatus,
} from 'shared-lib/models/event-request-response';

export class EventService {
	constructor() {}

	onConfigUpdateSubject = new ReplaySubject<Config>(1);
	onConfigUpdate$ = this.onConfigUpdateSubject.asObservable();

	onModelCheckSubject = new ReplaySubject<OnModelCheck>(1);
	onModelCheck$ = this.onModelCheckSubject.asObservable();

	onDockerInfoSubject = new ReplaySubject<OnDockerInfo>(1);
	onDockerInfo$ = this.onDockerInfoSubject.asObservable();

	fileDownloadStopRequest = new ReplaySubject<void>(1);
	fileDownloadStopRequest$ = this.fileDownloadStopRequest.asObservable();

	onFileDownloadStatusSubject = new ReplaySubject<OnFileDownloadStatus>(1);
	onFileDownloadStatus$ = this.onFileDownloadStatusSubject.asObservable();

	onGraphicsInfoSubject = new ReplaySubject<OnGraphicsInfo>(1);
	onGraphicsInfo$ = this.onGraphicsInfoSubject.asObservable();

	onOsInfoSubject = new ReplaySubject<OnOSInfo>(1);
	onOsInfo$ = this.onOsInfoSubject.asObservable();

	modelLaunchRequestSubject = new ReplaySubject<ModelLaunchRequest>(1);
	modelLaunchRequest$ = this.modelLaunchRequestSubject.asObservable();

	onModelLaunchSubject = new ReplaySubject<OnModelLaunch>(1);
	onModelLaunch$ = this.onModelLaunchSubject.asObservable();

	installRuntimeRequest = new ReplaySubject<void>(1);
	installRuntimeRequest$ = this.installRuntimeRequest.asObservable();

	onRuntimeInstallLogSubject = new ReplaySubject<string>(1);
	onRuntimeInstallLog$ = this.onRuntimeInstallLogSubject.asObservable();

	onLLMAddressChangeSubject = new ReplaySubject<string>(1);
	onLLMAddressChange$ = this.onLLMAddressChangeSubject.asObservable();

	onDockerImagePullStatusSubject = new ReplaySubject<DockerImagePullStatus>(1);
	onDockerImagePullStatus$ = this.onDockerImagePullStatusSubject.asObservable();
}

export const eventService = new EventService();

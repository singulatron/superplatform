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
import { DockerService } from './docker.service';
import { ReplaySubject, combineLatest } from 'rxjs';
import {
	OnModelLaunch,
	OnModelCheck,
} from 'shared-lib/models/event-request-response';

@Injectable({
	providedIn: 'root',
})
export class ModelService {
	onModelCheckSubject = new ReplaySubject<OnModelCheck>(1);
	/** Emitted any time when the currently selected model is checked */
	onModelCheck$ = this.onModelCheckSubject.asObservable();

	onModelLaunchSubject = new ReplaySubject<OnModelLaunch>(1);
	/** Emitted when the model is launched and available shortly */
	onModelLaunch$ = this.onModelLaunchSubject.asObservable();

	onModelReadySubject = new ReplaySubject<OnModelReady>(1);
	onModelReady$ = this.onModelReadySubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private dockerService: DockerService
	) {
		// @todo nothing to trigger model start so we resolve to polling
		setInterval(() => {
			this.init();
		  }, 2000);

		this.listenToModelReady();
	}

	private listenToModelReady(): void {
		combineLatest([
			this.dockerService.onDockerInfo$,
			this.onModelCheck$,
		]).subscribe(([dockerInfo, modelCheck]) => {
			if (dockerInfo.hasDocker && modelCheck.selectedExists) {
				this.onModelReadySubject.next({ modelReady: true });
			}
		});
	}

	async init() {
		try {
			let rsp = await this.modelStatus();

			if (rsp?.status?.selectedExists) {
				await this.modelStart().catch((e) => {
					console.error('Error starting model', {
						error: e,
					});
				});
			}
			if (rsp?.status?.running) {
				this.onModelLaunchSubject.next({});
			}
			this.onModelCheckSubject.next({
				selectedExists: rsp?.status?.selectedExists,
			});
		} catch (error) {
			console.log(error)
			console.error('Error in pollModelStatus', {
				error: JSON.stringify(error),
			});
		}
	}

	async modelStatus(): Promise<ModelStatusResponse> {
		return this.localtron.call('/model/status', {});
	}

	async modelStart(url?: string) {
		this.localtron.call('/model/start', { url: url });
	}
}

export interface OnModelReady {
	modelReady: boolean;
}

interface ModelStatus {
	selectedExists: boolean;
	currentModelId: string;
	/** Running triggers onModelLaunch on the frontend.
	 * Running is true when the model is both running and answering
	 * - fully loaded.
	 */
	running: boolean;
	modelAddress: string;
}

// {
//   "status": {
//     "selectedExists": false,
//     "currentModelId": "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q3_K_S.gguf"
//   }
// }
interface ModelStatusResponse {
	status: ModelStatus | null;
}

// also duplicated in the api service
export interface Model {
	/** id is the download url of the model */
	id: string;
	name: string;
	parameters?: string;
	flavour?: string;
	version?: string;
	quality?: string;
	extension?: string;
	fullName?: string;
	tags?: string[];
	mirrors?: string[];
	size?: number;
	uncensored?: boolean;
	maxRam?: number;
	description?: string;
	promptTemplate?: string;
	quantComment?: string;
}

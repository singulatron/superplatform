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
import { ApiService } from 'shared/stdlib/api.service';

@Injectable({
	providedIn: 'root',
})
export class ModelService {
	private onModelCheckSubject = new ReplaySubject<OnModelCheck>(1);
	/** Emitted any time when the currently selected model is checked */
	public onModelCheck$ = this.onModelCheckSubject.asObservable();

	private onModelLaunchSubject = new ReplaySubject<OnModelLaunch>(1);
	/** Emitted when the model is launched and available shortly */
	public onModelLaunch$ = this.onModelLaunchSubject.asObservable();

	private onModelReadySubject = new ReplaySubject<OnModelReady>(1);
	public onModelReady$ = this.onModelReadySubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private apiService: ApiService,
		private dockerService: DockerService
	) {
		// @todo nothing to trigger model start so we resolve to polling
		setInterval(() => {
			this.init();
		}, 2000);

		this.listenToModelReady();
	}

	modelsFromServer = [];
	serverModelsChecked = false;

	async getModels(): Promise<Model[]> {
		if (!this.serverModelsChecked) {
			this.serverModelsChecked = true;

			try {
				let rsp: ModelResponse = await this.apiService.getModelsFromServer();
				if (rsp.models && Array.isArray(rsp.models) && rsp.models.length > 0) {
					this.modelsFromServer = rsp.models as any;
				}
			} catch (e) {
				console.error('Error getting models from server', {
					error: JSON.stringify(e),
				});
			}
		}

		models.push(...this.modelsFromServer);
		return models;
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
				await this.modelStartWrapper(rsp?.status.currentModelId).catch((e) => {
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
			console.log(error);
			console.error('Error in pollModelStatus', {
				error: JSON.stringify(error),
			});
		}
	}

	async modelStatus(): Promise<ModelStatusResponse> {
		return this.localtron.call('/model/status', {});
	}

	private async modelStartWrapper(modelId: string) {
		let models = await this.getModels();
		let model = models.find((model) => model.id == modelId);
		if (!model) {
			throw `Model ${modelId} not found`;
		}
		await this.modelStart(model.platform, model.assets);
		return;
	}

	async modelStart(
		platform: Platform,
		assets: { [key: string]: string }
	): Promise<ModelStartResponse> {
		let req: ModelStartRequest = {
			platform: platform,
			assets: assets,
		};
		return this.localtron.call('/model/start', req);
	}

	async makeDefault(url?: string) {
		this.localtron.call('/model/make-default', { url: url });
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

interface ModelStartRequest {
	platform: Platform;
	assets: { [key: string]: string };
}

interface ModelStartResponse {}

export interface Platform {
	id: string;
	name?: string;
	version?: number;
	container: PlatformContainer;
}

export interface PlatformContainer {
	/** Internal port */
	port: number;
	images: PlatformImages;
}

export interface PlatformImages {
	default: string;
	cuda?: string;
}

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

export const PlatformLlamaCpp: Platform = {
	id: 'llama-cpp',
	container: {
		port: 8000,
		images: {
			default: 'crufter/llama-cpp-python-simple',
			cuda: 'crufter/llama-cpp-python-cuda',
		},
	},
};

export const PlatformStableDiffusion: Platform = {
	id: 'stable-diffusion',
	container: {
		port: 7860,
		images: {
			default: 'nicklucche/stable-diffusion',
		},
	},
};

export interface Model {
	id: string;
	platform: Platform;
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
	maxBits?: number;
	bits?: number;
	/** Asset maps asset name to URL, eg:
	 * 	MODEL: "https://huggingface.co/TheBloke/Mistral-7B-Instruct-v0.2-GGUF/resolve/main/mistral-7b-instruct-v0.2.Q2_K.gguf"
	 *
	 *  The asset will be downloaded and passed in to the container
	 * as an envar (under the name MODEL).
	 */
	assets: { [key: string]: string };
}

export interface ModelResponse {
	models: Model[];
}

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
	private initInProgress: boolean = false;

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
		private dockerService: DockerService
	) {
		// @todo nothing to trigger model start so we resolve to polling
		setInterval(() => {
			this.init();
		}, 2000);

		this.listenToModelReady();
	}

	models: Model[] = [];

	async getModels(): Promise<Model[]> {
		if (this.models?.length) {
			return this.models;
		}

		let rsp: GetModelsResponse = await this.localtron.call(
			'/model/get-models',
			{}
		);
		return rsp.models;
	}

	private listenToModelReady(): void {
		combineLatest([
			this.dockerService.onDockerInfo$,
			this.onModelCheck$,
		]).subscribe(([dockerInfo, modelCheck]) => {
			if (dockerInfo.hasDocker && modelCheck.assetsReady) {
				this.onModelReadySubject.next({ modelReady: true });
			}
		});
	}

	async init() {
		try {
			if (this.initInProgress) {
				return;
			}
			this.initInProgress = true;

			this.models = await this.getModels();
			let rsp = await this.modelStatus();

			this.onModelCheckSubject.next({
				assetsReady: rsp?.status?.assetsReady,
			});

			if (rsp?.status?.running) {
				this.onModelLaunchSubject.next({});
			}

			if (rsp?.status?.assetsReady) {
				await this.modelStart();
			}
		} catch (error) {
			console.log(error);
			console.error('Error in model.service init', {
				error: JSON.stringify(error),
			});
		} finally {
			this.initInProgress = false;
		}
	}

	async modelStatus(modelId?: string): Promise<ModelStatusResponse> {
		let req: ModelStatusRequest = {
			modelId: modelId,
		};
		return this.localtron.call('/model/status', req);
	}

	async modelStart(modelId?: string): Promise<ModelStartResponse> {
		let req: ModelStartRequest = {
			modelId: modelId,
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
	assetsReady: boolean;
	/** Running triggers onModelLaunch on the frontend.
	 * Running is true when the model is both running and answering
	 * - fully loaded.
	 */
	running: boolean;
	address: string;
}

interface ModelStatusRequest {
	modelId?: string;
}

interface ModelStatusResponse {
	status: ModelStatus | null;
}

interface ModelStartRequest {
	modelId?: string;
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

export interface GetModelsResponse {
	models: Model[];
}

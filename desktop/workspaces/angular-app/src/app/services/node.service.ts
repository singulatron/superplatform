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
export class NodeService {
	constructor(private localtron: LocaltronService) {}

	async nodesList(request: ListNodesRequest): Promise<ListNodesResponse> {
		return this.localtron.call('/node/list', request);
	}
}

export interface GPU {
	id: string;
	intraNodeId: number;
	name: string;
	busId: string;
	temperature: number;
	performanceState: string;
	powerUsage: number;
	powerCap: number;
	memoryUsage: number;
	memoryTotal: number;
	gpuUtilization: number;
	computeMode: string;
	processDetails?: Process[];
}

export interface Process {
	pid: number;
	processName: string;
	nemoryUsage: number;
}

export interface Node {
	hostname: string;
	gpus: GPU[];
}

export interface Cluster {
	nodes: Node[];
}

// eslint-disable-next-line
export interface ListNodesRequest {}

export interface ListNodesResponse {
	nodes: Node[];
}

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
	Id: string;
	IntraNodeId: number;
	Name: string;
	BusId: string;
	Temperature: number;
	PerformanceState: string;
	PowerUsage: number;
	PowerCap: number;
	MemoryUsage: number;
	MemoryTotal: number;
	GPUUtilization: number;
	ComputeMode: string;
	ProcessDetails: Process[];
}

export interface Process {
	Pid: number;
	ProcessName: string;
	MemoryUsage: number;
}

export interface Node {
	Hostname: string;
	GPUs: GPU[];
}

export interface Cluster {
	Nodes: Node[];
}

export interface ListNodesRequest {}

// eslint-disable-next-line
export interface ListNodesResponse {
	nodes: Node[];
}

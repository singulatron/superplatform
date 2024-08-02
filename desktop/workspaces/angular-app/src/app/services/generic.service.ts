/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { FirehoseService } from './firehose.service';
import { first } from 'rxjs';
import { UserService } from './user.service';
import {
	GenericObject,
	Condition,
	CreateRequest,
	FindRequest,
	FindResponse,
	UpdateRequest,
	UpdateResponse,
	UpsertRequest,
	DeleteRequest,
	DeleteResponse,
} from '@singulatron/types';

@Injectable({
	providedIn: 'root',
})
export class GenericService {
	constructor(
		private localtron: LocaltronService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
			}
			return;
		});
	}

	async create(table: string, object: GenericObject): Promise<void> {
		object.table = table;
		const request: CreateRequest = {
			object: object,
		};

		return this.localtron.post('/generic-service/object', request);
	}

	async find(table: string, conditions: Condition[]): Promise<FindResponse> {
		const request: FindRequest = {
			table: table,
			query: {
				conditions: conditions,
			},
		};

		return this.localtron.post('/generic-service/objects', request);
	}

	async upsert(table: string, object: GenericObject): Promise<void> {
		object.table = table;
		const request: UpsertRequest = {
			object: object,
		};

		return this.localtron.put(`/generic-service/object/${object.id}`, request);
	}

	async update(
		table: string,
		conditions: Condition[],
		object: GenericObject
	): Promise<UpdateResponse> {
		const request: UpdateRequest = {
			table: table,
			conditions: conditions,
			object: object,
		};

		return this.localtron.post('/generic-service/objects/update', request);
	}

	async delete(
		table: string,
		conditions: Condition[]
	): Promise<DeleteResponse> {
		const request: DeleteRequest = {
			table: table,
			conditions: conditions,
		};

		return this.localtron.delete('/generic-service/objects/delete', request);
	}
}

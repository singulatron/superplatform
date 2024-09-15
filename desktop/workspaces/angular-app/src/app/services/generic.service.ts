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
	Configuration,
	DynamicSvcApi,
	DatastoreFilter,
	DynamicSvcCreateObjectRequest as CreateObjectRequest,
	DynamicSvcQueryRequest as QueryRequest,
	DynamicSvcQueryResponse as QueryResponse,
	DynamicSvcObject as DynamicObject,
	DynamicSvcUpdateObjectRequest as UpdateObjectRequest,
	DynamicSvcCreateObjectResponse as UpdateObjectResponse,
	DynamicSvcUpsertObjectRequest as UpsertObjectRequest,
	DynamicSvcCreateObjectResponse,
	DynamicSvcUpsertObjectResponse,
	// DynamicSvcDeleteObjectRequest as DeleteObjectRequest,
	// DynamicSvcDeleteObjectResponse as DeleteObjectResponse,
} from '@singulatron/client';

@Injectable({
	providedIn: 'root',
})
export class DynamicService {
	dynamicService!: DynamicSvcApi;

	constructor(
		private localtron: LocaltronService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
			this.dynamicService = new DynamicSvcApi(
				new Configuration({
					apiKey: this.localtron.token(),
					basePath: this.localtron.addr(),
				})
			);
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
			}
			return;
		});
	}

	async create(
		table: string,
		object: DynamicObject
	): Promise<DynamicSvcCreateObjectResponse> {
		object.table = table;
		const request: CreateObjectRequest = {
			object: object,
		};

		return this.dynamicService.createObject({ body: request });
	}

	async find(
		table: string,
		filters: DatastoreFilter[]
	): Promise<QueryResponse> {
		const request: QueryRequest = {
			table: table,
			query: {
				filters: filters,
			},
		};

		return this.dynamicService.query({ body: request });
	}

	async upsert(
		table: string,
		object: DynamicObject
	): Promise<DynamicSvcUpsertObjectResponse> {
		object.table = table;
		const request: UpsertObjectRequest = {
			object: object,
		};

		return this.dynamicService.upsertObject({
			objectId: object.id!,
			body: request,
		});
	}

	async update(
		table: string,
		filters: DatastoreFilter[],
		object: DynamicObject
	): Promise<UpdateObjectResponse> {
		const request: UpdateObjectRequest = {
			table: table,
			filters: filters,
			object: object,
		};

		return this.dynamicService.updateObjects({
			body: request,
		});
	}

	async delete(table: string, filters: DatastoreFilter[]): Promise<any> {
		console.log(table, filters);

		// const request = {
		// 	table: table,
		// 	filters: filters,
		// };
		//eturn this.dynamicService.deleteObjects({
		//	body: request,
		//);
	}
}

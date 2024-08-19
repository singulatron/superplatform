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
	GenericSvcApi,
	DatastoreCondition,
	GenericSvcCreateObjectRequest as CreateObjectRequest,
	GenericSvcQueryRequest as QueryRequest,
	GenericSvcQueryResponse as QueryResponse,
	GenericSvcGenericObject as GenericObject,
	GenericSvcUpdateObjectRequest as UpdateObjectRequest,
	GenericSvcCreateObjectResponse as UpdateObjectResponse,
	GenericSvcUpsertObjectRequest as UpsertObjectRequest,
	GenericSvcCreateObjectResponse,
	GenericSvcUpsertObjectResponse,
	// GenericSvcDeleteObjectRequest as DeleteObjectRequest,
	// GenericSvcDeleteObjectResponse as DeleteObjectResponse,
} from '@singulatron/client';

@Injectable({
	providedIn: 'root',
})
export class DynamicService {
	dynamicService!: GenericSvcApi;

	constructor(
		private localtron: LocaltronService,
		private userService: UserService,
		private firehoseService: FirehoseService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
			this.dynamicService = new GenericSvcApi(
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
		object: GenericObject
	): Promise<GenericSvcCreateObjectResponse> {
		object.table = table;
		const request: CreateObjectRequest = {
			object: object,
		};

		return this.dynamicService.createObject({ body: request });
	}

	async find(
		table: string,
		conditions: DatastoreCondition[]
	): Promise<QueryResponse> {
		const request: QueryRequest = {
			table: table,
			query: {
				conditions: conditions,
			},
		};

		return this.dynamicService.query({ body: request });
	}

	async upsert(
		table: string,
		object: GenericObject
	): Promise<GenericSvcUpsertObjectResponse> {
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
		conditions: DatastoreCondition[],
		object: GenericObject
	): Promise<UpdateObjectResponse> {
		const request: UpdateObjectRequest = {
			table: table,
			conditions: conditions,
			object: object,
		};

		return this.dynamicService.updateObjects({
			body: request,
		});
	}

	async delete(table: string, conditions: DatastoreCondition[]): Promise<any> {
		console.log(table, conditions);

		// const request = {
		// 	table: table,
		// 	conditions: conditions,
		// };
		//eturn this.dynamicService.deleteObjects({
		//	body: request,
		//);
	}
}

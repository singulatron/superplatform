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
import { FirehoseService } from './firehose.service';
import { first } from 'rxjs';
import { UserService } from './user.service';

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
		const request: CreateRequest = {
			table: table,
			object: object,
		};

		return this.localtron.call('/generic/create', request);
	}

	async find(table: string, conditions: Condition[]): Promise<FindResponse> {
		const request: FindRequest = {
			table: table,
			conditions: conditions,
		};

		return this.localtron.call('/generic/find', request);
	}

	async upsert(table: string, object: GenericObject): Promise<void> {
		const request: UpsertRequest = {
			table: table,
			object: object,
		};

		return this.localtron.call('/generic/upsert', request);
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

		return this.localtron.call('/generic/update', request);
	}

	async delete(
		table: string,
		conditions: Condition[]
	): Promise<DeleteResponse> {
		const request: DeleteRequest = {
			table: table,
			conditions: conditions,
		};

		return this.localtron.call('/generic/delete', request);
	}
}

export interface FieldSelector {
	field?: string;
	oneOf?: string[];
	any?: boolean;
}

export function field(fieldName: string): FieldSelector {
	return {
		field: fieldName,
	};
}

export function fields(fieldNames: string[]): FieldSelector {
	return {
		oneOf: fieldNames,
	};
}

export function anyField(): FieldSelector {
	return {
		any: true,
	};
}

export interface Query {
	conditions?: Condition[];
	after?: any[];
	limit?: number;
	orderBys?: OrderBy[];
	count?: boolean;
}

export interface OrderBy {
	field: string;
	desc: boolean;
}

// this could be a sumtype, eg. EqualCondition | AllCondition
// but it's defined as a product type here to match the backend Go structure
// for easier understanding
export interface Condition {
	equal?: EqualCondition;
	all?: AllCondition;
	startsWith?: StartsWithCondition;
	contains?: ContainsCondition;
}

export function conditionFieldIs(
	condition: Condition,
	fieldName: string
): boolean {
	if (
		condition.equal &&
		(condition.equal.selector.field == fieldName ||
			condition.equal.selector.oneOf?.includes(fieldName))
	) {
		return true;
	}
	if (
		condition.contains &&
		(condition.contains.selector.field == fieldName ||
			condition.contains.selector.oneOf?.includes(fieldName))
	) {
		return true;
	}
	if (
		condition.startsWith &&
		(condition.startsWith.selector.field == fieldName ||
			condition.startsWith.selector.oneOf?.includes(fieldName))
	) {
		return true;
	}

	return false;
}

export function conditionField(condition: Condition): string {
	if (condition.equal) {
		return (
			condition.equal.selector.field! ||
			condition.equal.selector.oneOf?.join(',') ||
			''
		);
	}
	if (condition.contains) {
		return (
			condition.contains.selector.field ||
			condition.contains.selector.oneOf?.join(',') ||
			''
		);
	}
	if (condition.startsWith) {
		return (
			condition.startsWith.selector.field! ||
			condition.startsWith.selector.oneOf?.join(',') ||
			''
		);
	}

	return '';
}

export function conditionsToKeyValue(conditions: Condition[]): {
	[key: string]: any;
} {
	if (!conditions) {
		return {};
	}
	const object: { [key: string]: any } = {};

	for (const condition of conditions) {
		object[conditionField(condition)] = conditionValue(condition);
	}

	return object;
}

export function conditionValue(condition: Condition): any {
	if (condition.equal) {
		return condition.equal.value;
	}
	if (condition.contains) {
		return condition.contains.value;
	}
	if (condition.startsWith) {
		return condition.startsWith.value;
	}

	return '';
}

export function queryHasFieldCondition(
	query: Query,
	fieldName: string
): boolean {
	if (!query.conditions) {
		return false;
	}
	for (const condition of query.conditions) {
		if (conditionFieldIs(condition, fieldName)) {
			return true;
		}
	}

	return false;
}

export interface EqualCondition {
	selector: FieldSelector;
	value: any;
}

export interface StartsWithCondition {
	selector: FieldSelector;
	value: any;
}

export interface ContainsCondition {
	selector: FieldSelector;
	value: any;
}

// eslint-disable-next-line
export interface AllCondition {}

export function equal(selector: FieldSelector, value: any): Condition {
	return {
		equal: {
			selector,
			value,
		},
	};
}

export function startsWith(selector: FieldSelector, value: any): Condition {
	return {
		startsWith: {
			selector,
			value,
		},
	};
}

/* contains creates a Condition for the given fields specifed by the selector
 * eg. 'field1:~something' can be acquired by contains(field("field1"), "something")
 * 'field1,field2:~something' can be acquired by contains(fields("field1", "field2"), "something")
 */
export function contains(selector: FieldSelector, value: any): Condition {
	return {
		contains: {
			selector,
			value,
		},
	};
}

export function all(): Condition {
	return {
		all: {},
	};
}

export function id(id: string): Condition {
	return equal(field('id'), id);
}

export function userId(id: string): Condition {
	return equal(field('userId'), id);
}

export interface GenericObject {
	id: string;
	createdAt: string;
	updatedAt: string;
	userId?: string;
	data: any;
	public?: boolean;
}

export interface CreateRequest {
	table: string;
	object: GenericObject;
}

// eslint-disable-next-line
export interface CreateResponse {}

export interface UpdateRequest {
	table: string;
	conditions: Condition[];
	object: GenericObject;
}

// eslint-disable-next-line
export interface UpdateResponse {}

export interface DeleteRequest {
	table: string;
	conditions: Condition[];
}

// eslint-disable-next-line
export interface DeleteResponse {}

export interface FindRequest {
	table: string;
	conditions: Condition[];
}

export interface FindResponse {
	objects: GenericObject[];
}

export interface UpsertRequest {
	table: string;
	object: GenericObject;
}

// eslint-disable-next-line
export interface UpsertResponse {}

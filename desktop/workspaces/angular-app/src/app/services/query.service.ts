/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import {
	Condition,
	equal,
	contains,
	startsWith,
	field,
	Query,
} from '@singulatron/types/generic';

@Injectable({
	providedIn: 'root',
})
export class QueryService {
	constructor() {}
}

export class QueryParser {
	defaultConditionFunc!: (value: any) => Condition;

	parse(queryString: string): Query {
		const query: Query = {};

		// Extract and remove 'orderBy', 'limit', and 'after' parts from the query string first
		const orderByRegex = /orderBy:([\w,:-]+)/;
		const limitRegex = /limit:(\d+)/;
		const afterRegex = /after:([\w,]+)/;

		const orderByMatch = orderByRegex.exec(queryString);
		const limitMatch = limitRegex.exec(queryString);
		const afterMatch = afterRegex.exec(queryString);

		// Remove these parts from the query string
		queryString = queryString
			.replace(orderByRegex, '')
			.replace(limitRegex, '')
			.replace(afterRegex, '')
			.trim();

		if (orderByMatch) {
			query.orderBys = orderByMatch[1].split(',').map((field) => {
				const [fieldName, order] = field.split(':');
				return { field: fieldName, desc: order === 'desc' };
			});
		}

		if (limitMatch) {
			query.limit = Number.parseInt(limitMatch[1], 10);
		}

		if (afterMatch) {
			query.after = afterMatch[1].split(',');
		}

		if (!queryString) {
			return query;
		}

		if (!queryString.includes(':')) {
			if (this.defaultConditionFunc) {
				query.conditions = [this.defaultConditionFunc(queryString)]
			}

			return query;
		}

		// Regex to match field:value pairs including quoted values with spaces
		const fieldRegex = /(\w+(?:,\w+)*):(".*?"|[^ ]+)/g;
		let match;
		while ((match = fieldRegex.exec(queryString)) !== null) {
			const fields = match[1].split(',');
			let value = match[2];

			// Remove surrounding quotes from the value if they exist
			if (value.startsWith('"') && value.endsWith('"')) {
				value = value.slice(1, -1);
			}

			for (const field of fields) {
				if (!query.conditions) query.conditions = [];
				query.conditions.push(this.createCondition(field, value));
			}
		}

		return query;
	}

	public convertQueryParamsToSearchTerm(parameters: any): string {
		if (Object.keys(parameters)?.length == 1 && parameters['search']) {
			return parameters['search'];
		}

		return Object.entries(parameters)
			.filter((v) => {
				return v[0] !== 'search';
			})
			.map(([key, value]) => `${key}:${value}`)
			.join(' ');
	}


	private createCondition(fieldName: string, value: string): Condition {
		if (value.startsWith('~')) {
			return contains(field(fieldName), value.slice(1));
		} else if (value.startsWith('^')) {
			return startsWith(field(fieldName), value.slice(1));
		} else {
			const numericValue = Number(value);
			return equal(
				field(fieldName),
				Number.isNaN(numericValue) ? value : numericValue
			);
		}
	}
}

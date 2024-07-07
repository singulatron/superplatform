import { Injectable } from '@angular/core';
import {
	Condition,
	equal,
	contains,
	startsWith,
	field,
	fields,
} from './generic.service';

@Injectable({
	providedIn: 'root',
})
export class QueryService {
	constructor() {}
}

export class QueryParser {
	defaultFields = ['name'];

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
			query.orderBy = orderByMatch[1].split(',').map((field) => {
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
			this.addDefaultConditions(query, this.defaultFields, queryString);
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

			fields.forEach((field) => {
				if (!query.conditions) query.conditions = [];
				query.conditions.push(this.createCondition(field, value));
			});
		}

		return query;
	}

	private addDefaultConditions(
		query: Query,
		fieldNames: string[],
		value: string
	) {
		if (!query.conditions) query.conditions = [];
		if (fieldNames?.length > 1) {
			query.conditions?.push(contains(fields(fieldNames), value));
		} else {
			query.conditions?.push(contains(field(fieldNames[0]), value));
		}
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

interface Query {
	conditions?: Condition[];
	orderBy?: { field: string; desc: boolean }[];
	limit?: number;
	after?: string[];
}

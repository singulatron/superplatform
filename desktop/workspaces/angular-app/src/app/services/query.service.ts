import { Injectable } from '@angular/core';
import {
	Condition,
	equal,
	contains,
	startsWith,
	field,
} from './generic.service';

@Injectable({
	providedIn: 'root',
})
export class QueryService {
	constructor() {}
}

export class QueryParser {
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

		// Regex to match field:value pairs including quoted values with spaces
		const fieldRegex = /(\w+):(".*?"|[^ ]+)/g;
		let match;
		while ((match = fieldRegex.exec(queryString)) !== null) {
			const field = match[1];
			let value = match[2];

			// Remove surrounding quotes from the value if they exist
			if (value.startsWith('"') && value.endsWith('"')) {
				value = value.slice(1, -1);
			}

			if (!query.conditions) query.conditions = [];
			query.conditions.push(this.createCondition(field, value));
		}

		return query;
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

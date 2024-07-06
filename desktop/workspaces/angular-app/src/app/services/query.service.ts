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

		const orderByRegex = /orderBy:([\w:,-]+)/;
		const orderByMatch = orderByRegex.exec(queryString);

		if (orderByMatch) {
			query.orderBy = orderByMatch[1].split(',').map((field) => {
				const [fieldName, order] = field.split(':');
				return { field: fieldName, desc: order === 'desc' };
			});
		}

		const limitRegex = /limit:(\d+)/;
		const limitMatch = limitRegex.exec(queryString);
		if (limitMatch) {
			query.limit = Number.parseInt(limitMatch[1], 10);
		}

		const afterRegex = /after:([\w,]+)/;
		const afterMatch = afterRegex.exec(queryString);
		if (afterMatch) {
			query.after = afterMatch[1].split(',');
		}

		return query;
	}

	private createCondition(fieldName: string, value: string): Condition {
		if (value.startsWith('~')) {
			return contains(field(fieldName), value.slice(1));
		} else if (value.startsWith('^')) {
			return startsWith(field(fieldName), value.slice(1));
		} else {
			return equal(
				field(fieldName),
				Number.isNaN(value as any) ? value : Number.parseFloat(value)
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

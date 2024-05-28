// replace local URL with remote URL

import { readFileSync, writeFileSync } from 'fs';
import { join } from 'path';

const filePath = join(__dirname, '.webpack', 'main', 'index.js');
const localURL = 'http://127.0.0.1:8080';
const productionURL = 'https://api.commonagi.com';

try {
	const data = readFileSync(filePath, 'utf8');

	const result = data.replace(new RegExp(localURL, 'g'), productionURL);

	writeFileSync(filePath, result, 'utf8');
} catch (err) {
	console.error(err);
}

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
import { Pipe, PipeTransform, Inject, InjectionToken } from '@angular/core';
import { ApiService } from './api.service';

export const TRANSLATE_OBJECT = new InjectionToken<string>('translateObject');
// key to language to translation
export type TranslateObject = { [key: string]: { [key: string]: string } };

@Pipe({
    name: 'translate',
    pure: true,
    standalone: true,
})
export class TranslatePipe implements PipeTransform {
	constructor(
		private apiService: ApiService,
		@Inject(TRANSLATE_OBJECT) private translations: TranslateObject
	) {}

	transform(key: string): string {
		if (!this.translations[key]) {
			return key;
		}
		const lang = this.apiService.getLocale();
		if (lang) {
			return this.translations[key][lang];
		}
		if (this.translations[key]['en']) {
			return this.translations[key]['en'];
		}
		return key;
	}
}

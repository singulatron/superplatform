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
import { Injectable, Type, ComponentRef } from '@angular/core';
import { ReplaySubject } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class UiService {
	private isMobile = false;
	private footerComponentRef: ComponentRef<any> | null = null;

	isMobileSubject = new ReplaySubject<boolean>(1);
	isMobile$ = this.isMobileSubject.asObservable();

	footerComponentSubject = new ReplaySubject<Type<any> | null>(1);
	footerComponent$ = this.footerComponentSubject.asObservable();

	constructor() {}

	setIsMobile(isMobile: boolean) {
		this.isMobile = isMobile;
		this.isMobileSubject.next(isMobile);
	}

	getIsMobile(): boolean {
		return this.isMobile;
	}

	setFooterComponent(componentType: Type<any> | null) {
		this.footerComponentSubject.next(componentType);
	}

	clearFooterComponent() {
		if (this.footerComponentRef) {
			this.footerComponentRef.destroy();
			this.footerComponentRef = null;
		}
	}
}

import { Injectable, Type, ComponentRef } from '@angular/core';
import { ReplaySubject } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class UiService {
	private isMobile = false;

	private footerComponentRef: ComponentRef<any> | null = null;
	footerComponentSubject = new ReplaySubject<Type<any> | null>(1);
	footerComponent$ = this.footerComponentSubject.asObservable();

	isMobileSubject = new ReplaySubject<boolean>(1);
	isMobile$ = this.isMobileSubject.asObservable();

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

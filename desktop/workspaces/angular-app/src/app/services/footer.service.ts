import { Injectable, Type, ComponentRef } from '@angular/core';
import { ReplaySubject } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class FooterService {
	private hasFooter = false;

	// eslint-disable-next-line
	private footerComponentRef: ComponentRef<any> | null = null;
	footerComponentSubject = new ReplaySubject<Type<any> | null>(1);
	footerComponent$ = this.footerComponentSubject.asObservable();

	constructor() {}
	hasFooterComponent(): boolean {
		return this.hasFooter;
	}

	updateFooterComponent(componentType: Type<any> | null) {
		this.hasFooter = true;
		this.footerComponentSubject.next(componentType);
	}

	removeFooterComponent() {
		this.hasFooter = false;
		// eslint-disable-next-line
		this.footerComponentSubject.next(null);

		if (this.footerComponentRef) {
			this.footerComponentRef.destroy();
			// eslint-disable-next-line
			this.footerComponentRef = null;
		}
	}
}

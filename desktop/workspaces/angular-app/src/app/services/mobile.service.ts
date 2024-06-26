import { Injectable } from '@angular/core';
import { ReplaySubject } from 'rxjs';

@Injectable({
	providedIn: 'root',
})
export class MobileService {
	private isMobile = false;

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
}

import { ReplaySubject } from 'rxjs';

export class MobileService {
	private isMobile = false;

	isMobileSubject = new ReplaySubject<boolean>(1);
	isMobile$ = this.isMobileSubject.asObservable();

	constructor() {}

	setMobileStatus(isMobile: boolean) {
		this.isMobile = isMobile;
		this.isMobileSubject.next(isMobile);
	}

	getMobileStatus(): boolean {
		return this.isMobile;
	}
}

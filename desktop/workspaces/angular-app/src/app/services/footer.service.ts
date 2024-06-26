import { Injectable, Type, ComponentRef } from '@angular/core';
import { ReplaySubject } from 'rxjs';
import { Router, ActivatedRoute, NavigationEnd } from '@angular/router';
import { map, filter } from 'rxjs/operators';
import { ChatInputComponent } from '../chat/chat-box/chat-input/chat-input.component';

@Injectable({
	providedIn: 'root',
})
export class FooterService {
	private hasFooter = false;

	// eslint-disable-next-line
	private footerComponentRef: ComponentRef<any> | null = null;
	footerComponentSubject = new ReplaySubject<Type<any> | null>(1);
	footerComponent$ = this.footerComponentSubject.asObservable();

	constructor(
		private activatedRoute: ActivatedRoute,
		private router: Router
	) {
		// Since ionic lifecycle hooks dont seem to
		// be triggering properly - nor ngonint
		// we need to do this hack here
		// Idealy we would do something like this
		// in the chat box component:
		//  this.subscriptions.push(
		//  	this.mobile.isMobile$.subscribe((isMobile) => {
		//  		if (isMobile) {
		//  			this.footer.updateFooterComponent(ChatInputComponent);
		//  		}
		//  	})
		//  );
		// and then unscubscribe in ngOnDestroy.
		this.router.events
			.pipe(
				filter((event) => event instanceof NavigationEnd),
				map(() => this.activatedRoute),
				map((route) => {
					while (route.firstChild) route = route.firstChild;
					return route;
				}),
				filter((route) => route.outlet === 'primary'),
				map((route) => route.snapshot.url.join('/'))
			)
			.subscribe((path) => {
				console.log('path is', path);
				if (path === 'chat') {
					console.log('updating footer to ChatInputComponent');
					this.updateFooterComponent(ChatInputComponent);
				} else {
					console.log('removing footer');
					this.removeFooterComponent();
				}
			});
	}

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

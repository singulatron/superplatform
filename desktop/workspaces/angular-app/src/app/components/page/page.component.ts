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
import {
	ContentChildren,
	QueryList,
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	Input,
	HostListener,
	AfterContentInit,
	ViewChild,
} from '@angular/core';
import { Subscription } from 'rxjs';
import { CommonModule, NgFor } from '@angular/common';
import { IonicModule, IonMenu } from '@ionic/angular';
import { RouterLink } from '@angular/router';
import { NgStyle, NgIf, NgClass } from '@angular/common';
import { MobileService } from '../../services/mobile.service';
import { FooterService } from '../../services/footer.service';
import { LocaltronService } from '../../services/localtron.service';
import { Router, NavigationStart } from '@angular/router';

@Component({
	selector: 'app-page',
	standalone: true,
	templateUrl: './page.component.html',
	styleUrl: './page.component.scss',
	imports: [
		IonicModule,
		CommonModule,
		RouterLink,
		NgStyle,
		NgIf,
		NgFor,
		NgClass,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PageComponent implements AfterContentInit {
	/** We give unique IDs to the ion-menu and ion-content elements
	 * so multiple of them can coexists, since the page component is not a singleton one.
	 */
	id: string = '';

	@Input() menuWidth = '90%';
	@Input() menuEnabled = true;
	@Input() columnWidths: string[] = [];
	@Input() mobileColumnWidths: string[] = [];
	@Input() title: string = '';
	@Input() breakpoint: number = 768; // Default breakpoint is 768px

	@ContentChildren('columnContent') columnsContent!: QueryList<any>;
	@ContentChildren('mainContent') mainContent!: QueryList<any>;

	@ViewChild(IonMenu, { static: true }) menu!: IonMenu;

	columns: any[] = [];
	main: any;

	private subscriptions: Subscription[] = [];

	constructor(
		public mobile: MobileService,
		public footer: FooterService,
		private cd: ChangeDetectorRef,
		private localtron: LocaltronService,
		private router: Router
	) {
		this.id = this.localtron.uuid();
		this.cd.markForCheck();
	}

	@HostListener('window:resize', ['$event'])
	onResize() {
		this.mobile.setMobileStatus(window.innerWidth < this.breakpoint);
	}

	ngOnInit() {
		this.subscriptions.push(
			this.router.events.subscribe((event) => {
				if (event instanceof NavigationStart) {
					this.footer.removeFooterComponent();
					this.cd.markForCheck();
				}
			}),
			this.mobile.isMobile$.subscribe((isMobile) => {
				if (isMobile) {
					return;
				}

				this.footer.removeFooterComponent();
				this.cd.markForCheck();
			})
		);

		this.mobile.setMobileStatus(window.innerWidth < this.breakpoint);
		this.footer.footerComponent$.subscribe(() => {
			this.cd.markForCheck();
		});
	}

	ionViewWillLeave() {
		for (const s of this.subscriptions) {
			s.unsubscribe();
		}
	}

	ngAfterContentInit(): void {
		this.columns = this.columnsContent.toArray();
		this.main = this.mainContent.first;
	}

	getColumnWidth(index: number): string {
		if (this.mobile.getMobileStatus() && this.mobileColumnWidths[index]) {
			return this.mobileColumnWidths[index];
		}
		return this.columnWidths[index] || 'auto';
	}

	getBreakpointQuery(): string {
		return `(min-width: ${this.breakpoint}px)`;
	}

	toggleMenu() {
		this.menu.toggle();
	}
}

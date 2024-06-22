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
import { CommonModule, NgFor } from '@angular/common';
import { IonicModule, IonMenu } from '@ionic/angular';
import { RouterLink } from '@angular/router';
import { NgStyle, NgIf } from '@angular/common';

@Component({
	selector: 'app-page',
	standalone: true,
	templateUrl: './page.component.html',
	styleUrl: './page.component.scss',
	imports: [IonicModule, CommonModule, RouterLink, NgStyle, NgIf, NgFor],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PageComponent implements AfterContentInit {
	@Input() columnWidths: string[] = [];
  @Input() title: string = ""
	@Input() breakpoint: number = 768; // Default breakpoint is 768px

	@ContentChildren('columnContent') columnsContent!: QueryList<any>;
	@ContentChildren('mainContent') mainContent!: QueryList<any>;

	@ViewChild(IonMenu, { static: true }) menu!: IonMenu;

	columns: any[] = [];
	main: any;
	isMobile = false;

	@HostListener('window:resize', ['$event'])
	onResize() {
		this.isMobile = window.innerWidth < this.breakpoint;
	}

	ngOnInit() {
		this.isMobile = window.innerWidth < this.breakpoint;
	}

	constructor(private cd: ChangeDetectorRef) {
		this.cd.markForCheck();
	}

	ngAfterContentInit(): void {
		this.columns = this.columnsContent.toArray();
		this.main = this.mainContent.first;
	}

	getColumnWidth(index: number): string {
		return this.columnWidths[index] || 'auto';
	}

	getBreakpointQuery(): string {
		return `(min-width: ${this.breakpoint}px)`;
	}

	toggleMenu() {
		this.menu.toggle();
	}
}

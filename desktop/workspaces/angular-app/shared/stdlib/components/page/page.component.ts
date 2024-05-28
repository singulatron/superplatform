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
import { Input, Component, ViewContainerRef, ViewChild } from '@angular/core';
import { NavController } from '@ionic/angular';

@Component({
	selector: 'b-page',
	templateUrl: './page.component.html',
	styleUrls: ['./page.component.css'],
})
export class PageComponent {
	id = Math.random().toString(36).substring(7);

	@ViewChild('template', { static: true }) template;

	@Input() title: string;
	@Input() icon: string;
	@Input() noModal: boolean = false;
	@Input() appsModal: boolean = false;
	@Input() noBackButton: boolean = false;
	@Input() themeKey: string = '';

	constructor(
		private viewContainerRef: ViewContainerRef,
		public navCtrl: NavController
	) {}

	ngOnInit() {
		this.viewContainerRef.createEmbeddedView(this.template);
	}
}

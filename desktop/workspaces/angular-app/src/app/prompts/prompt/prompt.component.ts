/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Component, ChangeDetectionStrategy, Input } from '@angular/core';

import { NgFor, NgIf, NgStyle } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonicModule } from '@ionic/angular';
import { Prompt } from '@singulatron/types';
import { PageComponent } from '../../components/page/page.component';
import { IconMenuComponent } from '../../components/icon-menu/icon-menu.component';
import { CenteredComponent } from '../../components/centered/centered.component';
import { DatePipe } from '@angular/common';
@Component({
	selector: 'app-prompt',
	standalone: true,
	imports: [
		CenteredComponent,
		PageComponent,
		IconMenuComponent,
		IonicModule,
		NgFor,
		NgIf,
		FormsModule,
		NgStyle,
		DatePipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
	templateUrl: './prompt.component.html',
	styleUrl: './prompt.component.scss',
})
export class PromptComponent {
	@Input() prompt!: Prompt;
	@Input() expanded = false;

	constructor() {}
}

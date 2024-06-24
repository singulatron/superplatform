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
import { ChangeDetectionStrategy, Component, Input } from '@angular/core';
import { NgIf, NgTemplateOutlet, NgStyle } from '@angular/common';

@Component({
	selector: 'app-centered',
	templateUrl: './centered.component.html',
	styleUrl: './centered.component.css',
	standalone: true,
	imports: [NgIf, NgTemplateOutlet, NgStyle],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CenteredComponent {
	@Input() headerHeight = '25vh';
	@Input() vertical = false;
}

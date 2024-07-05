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
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';

import { NgFor, NgIf, NgStyle } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { IonicModule } from '@ionic/angular';
import { Prompt, PromptService } from '../services/prompt.service';
import { UserService } from '../services/user.service';
import { first } from 'rxjs';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { CenteredComponent } from '../components/centered/centered.component';
import { PromptComponent } from './prompt/prompt.component';

@Component({
	selector: 'app-prompts',
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
		PromptComponent,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
	templateUrl: './prompts.component.html',
	styleUrl: './prompts.component.scss',
})
export class PromptsComponent {
	prompts: Prompt[] = [];
	afters: any[] = [];
	done = false;
	request = {
		statuses: [
			'scheduled',
			'running',
			'completed',
			'errored',
			'abandoned',
			'canceled',
		],
		desc: true,
	};
	count = 0;

	constructor(
		private cd: ChangeDetectorRef,
		private userService: UserService,
		private promptService: PromptService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async loggedInInit() {
		const rsp = await this.promptService.promptList({
			...this.request,
			count: true,
		});
		this.prompts = rsp.prompts;
		this.count = rsp.count || 0;
		this.afters.push(rsp.after);
		this.cd.markForCheck();
	}

	async prev() {
		if (this.afters.length >= 2) {
			const rsp = await this.promptService.promptList({
				...this.request,
				after: this.afters.at(-1),
				count: true,
			});
			this.prompts = rsp.prompts;
			this.count = rsp.count || 0;
			this.afters.pop();
			this.cd.markForCheck();
		}
	}

	async next() {
		const rsp = await this.promptService.promptList({
			...this.request,
			after: this.afters.at(-1),
			count: true,
		});

		this.prompts = rsp.prompts;
		this.count = rsp.count || 0;
		if (rsp.after) {
			this.afters.push(rsp.after);
		} else {
			this.done = true;
		}

		this.cd.markForCheck();
	}

	pageCount(): number {
		return Math.ceil(this.count / 20);
	}
}

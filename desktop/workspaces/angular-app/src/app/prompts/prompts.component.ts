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
import {
	ListPromptsRequest,
	Prompt,
	PromptService,
} from '../services/prompt.service';
import { UserService } from '../services/user.service';
import { first } from 'rxjs';
import { PageComponent } from '../components/page/page.component';
import { IconMenuComponent } from '../components/icon-menu/icon-menu.component';
import { CenteredComponent } from '../components/centered/centered.component';
import { PromptComponent } from './prompt/prompt.component';
import { QueryParser } from '../services/query.service';
import {
	queryHasFieldCondition,
	field,
	equal,
} from '../services/generic.service';

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
	styleUrls: ['./prompts.component.scss'],
})
export class PromptsComponent {
	prompts: Prompt[] = [];
	done = false;
	after: any;
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
	searchTerm = '';

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
		this.search('');
	}

	async search(value: string) {
		this.searchTerm = value;
		this.done = false;
		this.q();
	}

	async q(after?: any) {
		let query = new QueryParser().parse(this.searchTerm);
		if (!query.conditions) {
			query.conditions = [];
		}
		query.count = true;

		let request: ListPromptsRequest = {
			query: query,
		};
		if (!request.query) {
			request.query = {};
		}

		if (after) {
			request.query.after = [after];
		}

		if (!queryHasFieldCondition(query, 'status')) {
			query.conditions.push(equal(field('status'), this.request.statuses));
		}

		const rsp = await this.promptService.promptList(request);

		if (rsp.prompts) {
			this.prompts = [...this.prompts, ...rsp.prompts];
		} else {
			this.prompts = [];
		}
		this.count = rsp.count || 0;

		if (rsp.after) {
			this.after = rsp.after;
		} else {
			this.done = true;
		}

		this.cd.markForCheck();
	}

	async loadMoreData() {
		if (this.done) {
			console.log('No more prompts to load');
			return;
		}
		await this.q(this.after);
	}
}

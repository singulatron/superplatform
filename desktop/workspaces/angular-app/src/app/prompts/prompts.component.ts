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
	conditionsToKeyValue,
	contains,
	Condition,
	conditionFieldIs,
} from '../services/generic.service';
import { ActivatedRoute, Router } from '@angular/router';

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
	queryParser: QueryParser;

	constructor(
		private cd: ChangeDetectorRef,
		private userService: UserService,
		private promptService: PromptService,
		private activatedRoute: ActivatedRoute,
		private router: Router
	) {
		this.queryParser = new QueryParser();
		this.queryParser.defaultConditionFunc = (value: any): Condition => {
			return contains(field('modelId'), value);
		};

		this.userService.user$.pipe(first()).subscribe(() => {
			this.initializeOnLogin();
		});
	}

	private initializeOnLogin() {
		this.activatedRoute.queryParams.subscribe((parameters) => {
			const search =
				this.queryParser.convertQueryParamsToSearchTerm(parameters);
			this.searchTerm = search;
			this.fetchPrompts();
			this.prompts = [];
			this.cd.markForCheck();
		});
	}

	public redirect() {
		const query = this.queryParser.parse(this.searchTerm);
		const kv = conditionsToKeyValue(
			query.conditions!.filter((v) => {
				return !conditionFieldIs(v, 'modelId');
			})
		);

		if (Object.keys(kv)?.length) {
			this.router.navigate([], {
				queryParams: kv,
			});
			return;
		}
		this.router.navigate([], {
			queryParams: { search: this.searchTerm },
		});
	}

	public async fetchPrompts() {
		const query = this.queryParser.parse(this.searchTerm);
		query.count = true;
		query.conditions = query.conditions || [];

		if (!queryHasFieldCondition(query, 'status')) {
			query.conditions.push(equal(field('status'), this.request.statuses));
		}

		const request: ListPromptsRequest = {
			query: query,
		};

		if (this.after) {
			request.query!.after = [this.after];
		}

		const response = await this.promptService.promptList(request);

		if (response.prompts) {
			this.prompts = [...this.prompts, ...response.prompts];
		} else {
			this.prompts = [];
		}

		this.count = response.count || 0;
		if (response.after && response.after != `0001-01-01T00:00:00Z`) {
			this.after = response.after || null;
		} else {
			this.after = undefined;
		}

		this.cd.markForCheck();
	}

	async loadMoreData() {
		if (!this.after) {
			console.log('No more prompts to load');
			return;
		}
		await this.fetchPrompts();
	}
}

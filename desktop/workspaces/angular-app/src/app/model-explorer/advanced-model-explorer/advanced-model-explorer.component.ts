/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Component,
	HostListener,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { ModelService } from '../../services/model.service';
import { ModelSvcModel as Model } from '@singulatron/client';
import { DownloadService } from '../../services/download.service';
import { DownloadStatusChangeEvent } from '@singulatron/types';
import { ConfigService } from '../../services/config.service';
import { IonicModule } from '@ionic/angular';
import { TranslatePipe } from '../../translate.pipe';
import { DecimalPipe } from '@angular/common';
import { DownloadingComponent } from '../../downloading/downloading.component';
import { FormsModule } from '@angular/forms';

const veryLargeScreenWidth = 2400;

@Component({
	selector: 'app-advanced-model-explorer',
	templateUrl: './advanced-model-explorer.component.html',
	styleUrl: './advanced-model-explorer.component.scss',
	standalone: true,
	imports: [
		CommonModule,
		IonicModule,
		FormsModule,
		DownloadingComponent,
		TranslatePipe,
		DecimalPipe,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AdvancedModelExplorerComponent {
	expandedStates = new Map<string, boolean>();

	allModels: Model[] = [];
	allFilteredModels: Model[] = [];
	models: Model[] = [];
	currentPage = 1;
	itemsPerPage = 9;
	totalItems = 0;
	gridView = true;
	veryLargeScreen = false;

	showOnlyDownloadedModels = false;
	searchQuery = '';
	modelCategoryOptions: ModelCategoryOption[] = [
		{ name: 'Instruct', value: 'Instruct', active: false },
		{ name: 'Code', value: 'Code', active: false },
		{ name: 'Chat', value: 'Chat', active: false },
		{ name: 'Uncensored', value: 'Uncensored', active: false },
	];

	constructor(
		public downloadService: DownloadService,
		private modelService: ModelService,
		public configService: ConfigService,
		private cd: ChangeDetectorRef
	) {
		this.detectLargeScreen();
	}

	@HostListener('window:resize', ['$event'])
	onResize() {
		this.detectLargeScreen();
	}

	detectLargeScreen() {
		const screenWidth = window.innerWidth;
		this.veryLargeScreen = screenWidth > veryLargeScreenWidth;
	}

	async filterModels() {
		if (!this.searchQuery) {
			this.allFilteredModels = await this.getModels();
			this.totalItems = this.allFilteredModels.length;
			this.loadPage(1);
			return;
		}
		const models = await this.getModels();
		this.allFilteredModels = models.filter((model) => {
			const m = {
				...model,
			};

			const subject =
				JSON.stringify(m) +
				(model.uncensored ? ' Uncensored ' : '') +
				` ${Math.floor(model.maxRam || 0)} gb` +
				` ${Math.floor(model.maxRam || 0)}gb` +
				' gb'.toLowerCase();

			return subject.includes(this.searchQuery.toLowerCase());
		});

		// After filtering, reload the pagination with the filtered list
		this.totalItems = this.allFilteredModels.length;
		this.loadPage(1); // Reset to the first page
	}

	modelCategoryClicked(option: ModelCategoryOption) {
		option.active = !option.active;
		this.filterModels();
	}

	async getModels(): Promise<Model[]> {
		const activeCategories = this.modelCategoryOptions.filter(
			(option) => option.active
		);
		let models = this.allModels;
		if (this.showOnlyDownloadedModels) {
			const downloadedModels = [];
			const downloadsResponse = await this.downloadService.downloadList();
			for (const model of models) {
				if (
					downloadsResponse.downloads?.some(
						(download) =>
							download.status === 'completed' &&
							model.assets &&
							Object.values(model.assets)?.includes(download.url!)
					)
				) {
					downloadedModels.push(model);
				}
			}
			models = downloadedModels;
		}

		return this.anyCategorySelected()
			? models.filter((model) => {
					const found = activeCategories.some((option) => {
						switch (option.value) {
							case 'Instruct':
							case 'Code':
							case 'Chat': {
								return option.value === model.flavour;
							}
							case 'Uncensored': {
								return model.uncensored;
							}
							default: {
								break;
							}
						}
						return '';
					});
					return found;
				})
			: models;
	}

	anyCategorySelected(): boolean {
		return !!this.modelCategoryOptions.some((option) => option.active);
	}

	loadPage(page: number) {
		this.currentPage = page;
		const startIndex = (page - 1) * this.itemsPerPage;
		const endIndex = startIndex + this.itemsPerPage;
		this.models = this.allFilteredModels.slice(startIndex, endIndex);
		this.cd.markForCheck();
	}

	async ngOnInit(): Promise<void> {
		this.allModels = await this.modelService.getModels();
		this.allFilteredModels = this.allModels;
		this.totalItems = this.allModels.length;
		this.loadPage(this.currentPage);
	}

	isDownloading(
		model: Model,
		status: DownloadStatusChangeEvent | null
	): boolean {
		if (status === null) {
			return false;
		}
		const c = status?.allDownloads?.find(
			(download) =>
				model.assets && Object.values(model.assets).includes(download.url)
		);
		if (c?.status === 'inProgress' || c?.status === 'paused') {
			return true;
		}
		return false;
	}

	// @todo rename this to make model default...
	// the wording activate makes one think there is only one model can be used at a time
	async makeModelDefault(modelId: string) {
		this.modelService.makeDefault(modelId);
	}

	flavourToolTip(flavour: string): string {
		switch (flavour) {
			case 'Instruct': {
				return 'Instruct models are good at completing tasks.';
			}
			case 'Chat': {
				return 'Chat models are designed for general chat.';
			}
			case 'Code': {
				return 'Code models are designed for programming tasks.';
			}
		}
		return `Flavour ${flavour}`;
	}

	downloaded(model: Model, status: DownloadStatusChangeEvent | null): boolean {
		if (status === null) {
			return false;
		}
		if (
			status?.allDownloads?.find(
				(download) =>
					model.assets && Object.values(model.assets)?.includes(download.url)
			)?.status === 'completed'
		) {
			return true;
		}
		return false;
	}

	progress(id: string, status: DownloadStatusChangeEvent): number {
		return status?.allDownloads?.find((v) => v.id === id)?.progress || 0;
	}

	async download(model: Model) {
		const assetURLs = Object.values(model.assets!);
		if (!assetURLs?.length) {
			throw `No assets to download for ${model.id}`;
		}

		for (const url of assetURLs) {
			this.downloadService.downloadDo(url);
		}
	}

	hasAssets(model: Model): boolean {
		if (!model?.assets) {
			return false;
		}
		return Object.values(model.assets)?.length > 0;
	}

	toggleItem(id: string) {
		const currentState = this.expandedStates.get(id);
		this.expandedStates.set(id, !currentState);
		this.cd.markForCheck();
	}

	getDescription(item: Model): string {
		if (!item.description) {
			return '';
		}
		const maxLength = 0;
		if (this.expandedStates.get(item.id!)) {
			return item.description || '';
		} else {
			return item.description.length > maxLength
				? item.description.slice(0, maxLength)
				: item.description;
		}
	}

	extractValueFromQuality(quality: string): number {
		const match = quality.match(/Q(\d+)\D*/);
		return match ? Number.parseInt(match[1], 10) : 0;
	}

	getStatValue(model: Model) {
		const value: number = model.quality
			? this.extractValueFromQuality(model.quality)
			: 1;

		return value;
	}

	getStatStyle(model: Model) {
		const value: number = model.quality
			? this.extractValueFromQuality(model.quality)
			: 1;

		const maxBits = model.maxBits || 8;

		const percentageValue = (value / maxBits) * 100;

		const hue = (value / maxBits) * 120;

		let backgroundColor = `hsl(${hue}, 100%, 50%)`; // Adjust the lightness and saturation if needed
		backgroundColor = '#aaa';

		return {
			'background-color': backgroundColor,
			width: `${percentageValue}%`,
		};
	}

	getColumnSize(): string {
		const screenWidth = window.innerWidth;

		return screenWidth > 1400 ? '4' : '6';
	}

	switchGridListView() {
		this.gridView = !this.gridView;
		this.filterModels();
	}

	trackById(_: number, message: { id: string }): string {
		return message.id;
	}
}

export interface ModelCategoryOption {
	name: string;
	value: string;
	active: boolean;
}

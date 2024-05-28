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
import { Component } from '@angular/core';
import { LapiService } from '../../services/lapi.service';
import { OnFileDownloadStatus } from 'shared-lib/models/event-request-response';
import { ElectronIpcService } from '../../services/electron-ipc.service';
import { ApiService, Model } from '../../../../shared/stdlib/api.service';
import { LocaltronService } from '../../services/localtron.service';

@Component({
	selector: 'app-advanced-model-explorer',
	templateUrl: './advanced-model-explorer.component.html',
	styleUrl: './advanced-model-explorer.component.scss',
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

	showOnlyDownloadedModels = false;
	searchQuery = '';
	modelCategoryOptions: ModelCategoryOption[] = [
		{ name: 'Instruct', value: 'Instruct', active: false },
		{ name: 'Code', value: 'Code', active: false },
		{ name: 'Chat', value: 'Chat', active: false },
		{ name: 'Uncensored', value: 'uncensored', active: false },
	];

	async filterModels() {
		if (!this.searchQuery) {
			this.allFilteredModels = await this.getModels();
			this.totalItems = this.allFilteredModels.length;
			this.loadPage(1);
			return;
		}
		this.allFilteredModels = (await this.getModels()).filter((model) => {
			let m = {
				...model,
			};
			delete m.uncensored;

			const subject =
				JSON.stringify(m) +
				(model.uncensored ? ' uncensored ' : '') +
				` ${Math.floor(model.maxRam)} gb` +
				` ${Math.floor(model.maxRam)}gb` +
				' gb'.toLowerCase();

			return subject.includes(this.searchQuery.toLowerCase());
		});

		// After filtering, reload the pagination with the filtered list
		this.totalItems = this.allFilteredModels.length;
		this.loadPage(1); // Reset to the first page
		console.log(this.allFilteredModels);
	}

	modelCategoryClicked(option: ModelCategoryOption) {
		option.active = !option.active;
		this.filterModels();
	}

	async getModels() {
		const activeCategories = this.modelCategoryOptions.filter(
			(option) => option.active
		);
		let models = this.allModels;
		if (this.showOnlyDownloadedModels) {
			let downloadsResponse = await this.localtron.downloadList();
			models = models.filter((model) => {
				return downloadsResponse.downloads.find(
					(download) =>
						download.status === 'completed' && download.id === model.id
				);
			});
		}
		return !this.anyCategorySelected()
			? models
			: models.filter((model) => {
					let found = activeCategories.find((option) => {
						switch (option.value) {
							case 'Instruct':
							case 'Code':
							case 'Chat':
								return option.value === model.flavour;
							default:
								break;
						}
					});
					return found;
				});
	}

	anyCategorySelected(): boolean {
		return !!this.modelCategoryOptions.find((option) => option.active);
	}

	loadPage(page: number) {
		this.currentPage = page;
		const startIndex = (page - 1) * this.itemsPerPage;
		const endIndex = startIndex + this.itemsPerPage;
		this.models = this.allFilteredModels.slice(startIndex, endIndex);
	}

	constructor(
		public lapi: LapiService,
		private api: ApiService,
		private ipcService: ElectronIpcService,
		private localtron: LocaltronService
	) {}

	async ngOnInit(): Promise<void> {
		this.allModels = await this.api.getModels();
		this.allFilteredModels = this.allModels;
		this.totalItems = this.allModels.length;
		this.loadPage(this.currentPage);
	}

	isDownloading(id: string, status: OnFileDownloadStatus): boolean {
		let c = status?.allDownloads?.find((v) => v.id === id);
		if (c?.status === 'inProgress' || c?.status === 'paused') {
			return true;
		}
		return false;
	}

	async activateModel(modelId: string) {
		this.localtron.modelStart(modelId);
	}

	flavourToolTip(flavour: string): string {
		switch (flavour) {
			case 'Instruct':
				return 'Instruct models are good at completing tasks.';
			case 'Chat':
				return 'Chat models are designed for general chat.';
			case 'Code':
				return 'Code models are designed for programming tasks.';
		}
		return `Flavour ${flavour}`;
	}

	downloaded(id: string, status: OnFileDownloadStatus): boolean {
		if (
			status?.allDownloads?.find((v) => v.id === id)?.status === 'completed'
		) {
			return true;
		}
	}

	progress(id: string, status: OnFileDownloadStatus): number {
		return status?.allDownloads?.find((v) => v.id === id)?.progress || 0;
	}

	async download(id: string) {
		this.localtron.downloadDo(id);
	}

	toggleItem(id: string) {
		const currentState = this.expandedStates.get(id);
		this.expandedStates.set(id, !currentState); // Toggle the state
	}

	// Method to decide which description to show (full or preview)
	getDescription(item) {
		const maxLength = 0;
		if (this.expandedStates.get(item.id)) {
			return item.description; // Item is expanded, show full description
		} else {
			// Item is not expanded, show preview (if necessary)
			return item.description.length > maxLength
				? item.description.substring(0, maxLength)
				: item.description;
		}
	}

	getStatValue(model, statType) {
		let value;
		const minSizeGB = 1; // Minimum model size in GB for "speed"
		const maxSizeGB = 30; // Maximum model size in GB for "quality"

		if (model.maxRam !== undefined) {
			const clampedRam = Math.max(minSizeGB, Math.min(model.maxRam, maxSizeGB));
			switch (statType) {
				case 'speed':
					value =
						100 * (1 - (clampedRam - minSizeGB) / (maxSizeGB - minSizeGB));
					break;
				case 'quality':
					value = (100 * (clampedRam - minSizeGB)) / (maxSizeGB - minSizeGB);
					break;
			}
		} else {
			value = 50; // Default to 50% if maxRam is undefined
		}

		return Math.round(value / 10);
	}

	getStatStyle(model, statType) {
		let value;
		const minSizeGB = 1; // Minimum model size in GB for "speed"
		const maxSizeGB = 30; // Maximum model size in GB for "quality"

		if (model.maxRam !== undefined) {
			const clampedRam = Math.max(minSizeGB, Math.min(model.maxRam, maxSizeGB));
			switch (statType) {
				case 'speed':
					value =
						100 * (1 - (clampedRam - minSizeGB) / (maxSizeGB - minSizeGB));
					break;
				case 'quality':
					value = (100 * (clampedRam - minSizeGB)) / (maxSizeGB - minSizeGB);
					break;
			}
		} else {
			value = 50; // Default to 50% if maxRam is undefined
		}

		// Ensure value is between 0 and 100
		value = Math.max(0, Math.min(100, value));

		// Map value to hue from 240 (blue) to 120 (green)
		const hue = 240 - (value * (240 - 120)) / 100;

		const backgroundColor = `hsl(${hue}, 100%, 50%)`; // Adjust the lightness and saturation if needed
		return {
			'background-color': backgroundColor,
			width: `${value}%`,
		};
	}

	switchGridListView() {
		this.gridView = !this.gridView;
		this.filterModels();
	}
}

export interface ModelCategoryOption {
	name: string;
	value: string;
	active: boolean;
}

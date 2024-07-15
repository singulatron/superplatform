/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import { Injectable } from '@angular/core';
import { LocaltronService } from './localtron.service';
import { FirehoseService } from './firehose.service';
import { ReplaySubject, first } from 'rxjs';
import { UserService } from './user.service';

@Injectable({
	providedIn: 'root',
})
export class DownloadService {
	onFileDownloadStatusSubject = new ReplaySubject<DownloadStatusChangeEvent>(1);
	onFileDownloadStatus$ = this.onFileDownloadStatusSubject.asObservable();

	constructor(
		private localtron: LocaltronService,
		private firehoseService: FirehoseService,
		private userService: UserService
	) {
		this.init();
		this.userService.user$.pipe(first()).subscribe(() => {
			this.loggedInInit();
		});
	}

	async init() {
		this.firehoseService.firehoseEvent$.subscribe(async (event) => {
			switch (event.name) {
				case 'downloadStatusChange': {
					const rsp = await this.downloadList();
					this.onFileDownloadStatusSubject.next({
						allDownloads: rsp.downloads,
					});
					break;
				}
			}
		});
	}

	async loggedInInit() {
		try {
			const rsp = await this.downloadList();
			this.onFileDownloadStatusSubject.next({
				allDownloads: rsp?.downloads as DownloadDetails[],
			});
		} catch (error) {
			console.error('Error in pollFileDownloadStatus', {
				error: JSON.stringify(error),
			});
		}
	}

	async downloadDo(url: string) {
		this.localtron.call('/download/do', { url: url });
	}

	async downloadPause(url: string) {
		this.localtron.call('/download/pause', { url: url });
	}

	async downloadList(): Promise<DownloadsResponse> {
		return this.localtron.call('/download/list', {});
	}
}

export interface DownloadDetails {
	id: string;
	url: string;
	fileName: string;
	dir?: string;
	progress?: number;
	downloadedBytes: number;
	fullFileSize?: number;
	status: 'inProgress' | 'completed' | 'paused' | 'cancelled' | 'failed';
	filePath?: string;
	paused?: boolean;
	cancelled?: boolean;
	error?: string;
}

type DownloadsResponse = {
	downloads: DownloadDetails[];
};

export interface DownloadStatusChangeEvent {
	allDownloads: DownloadDetails[];
}

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
import { ChangeDetectionStrategy, Component } from '@angular/core';
import { ElectronIpcService } from '../services/electron-ipc.service';
import { WindowApiConst } from 'shared-lib';
import { enableLogging, disableLogging } from '../app.component';
import { LogService } from '../services/log.service';
import { TranslatePipe } from '../../../shared/stdlib/translate.pipe';
import { TranslateModule } from '@ngx-translate/core';
import { RouterLink } from '@angular/router';
import { NgClass, NgIf, NgStyle } from '@angular/common';
import { IonicModule } from '@ionic/angular';
import { CenteredComponent } from '../../../shared/stdlib/components/centered/centered.component';
import { PageComponent } from '../../../shared/stdlib/components/page/page.component';
import { IconMenuComponent } from '../../../shared/stdlib/components/icon-menu/icon-menu.component';

@Component({
	selector: 'app-home',
	templateUrl: './home.component.html',
	styleUrl: './home.component.css',
	standalone: true,
	imports: [
		IconMenuComponent,
		PageComponent,
		CenteredComponent,
		IonicModule,
		NgClass,
		NgIf,
		RouterLink,
		TranslateModule,
		TranslatePipe,
		NgStyle,
	],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent {
	loggingEnabled = true;

	constructor(
		private logService: LogService,
		private ipcService: ElectronIpcService
	) {}

	async ngOnInit() {
		const logStatus = await this.logService.logStatus();
		this.loggingEnabled = logStatus.enabled;
		if (!this.loggingEnabled) {
			console.log('Logging is disabled');
			this.ipcService.send(WindowApiConst.DISABLE_LOGGING_REQUEST, {});
		}
	}

	async toggleLogging() {
		if (this.loggingEnabled) {
			this.disableLog();
			return;
		}
		this.enableLog();
	}

	async enableLog() {
		this.loggingEnabled = true;
		await this.logService.logEnable();
		const rsp = await this.logService.logStatus();
		this.ipcService.send(WindowApiConst.ENABLE_LOGGING_REQUEST, {});
		this.loggingEnabled = rsp.enabled;
		enableLogging();
		console.log('Enabled logging');
	}

	async disableLog() {
		console.log('Disabled logging');
		disableLogging();
		this.loggingEnabled = false;
		await this.logService.logDisable();
		const rsp = await this.logService.logStatus();
		this.ipcService.send(WindowApiConst.DISABLE_LOGGING_REQUEST, {});
		this.loggingEnabled = rsp.enabled;
	}
}

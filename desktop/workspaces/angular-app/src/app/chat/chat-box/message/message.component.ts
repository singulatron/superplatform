/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
import {
	Component,
	Input,
	Output,
	EventEmitter,
	ChangeDetectionStrategy,
} from '@angular/core';
import { ChatService } from '../../../services/chat.service';
import {
	ChatSvcMessage as Message,
	ChatSvcAsset as Asset,
} from '@singulatron/client';
import { PromptService } from '../../../services/prompt.service';
import { LocaltronService } from '../../../services/server.service';
import { MarkdownComponent } from 'ngx-markdown';
import { IonicModule } from '@ionic/angular';
import { NgIf, DatePipe, AsyncPipe } from '@angular/common';
import { MobileService } from '../../../services/mobile.service';
import { UserService } from '../../../services/user.service';

@Component({
	selector: 'app-message',
	templateUrl: './message.component.html',
	styleUrl: './message.component.scss',
	standalone: true,
	imports: [NgIf, IonicModule, MarkdownComponent, DatePipe, AsyncPipe],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MessageComponent {
	constructor(
		private chatService: ChatService,
		private promptService: PromptService,
		public userService: UserService,
		private server: LocaltronService,
		public mobile: MobileService
	) {}
	hasAsset = false;

	@Input() message!: Message;
	@Input() assets: Asset[] = [];
	@Input() streaming: boolean = false;
	@Input() modelId: string = '';

	@Output() onCopyToClipboard = new EventEmitter<string>();

	ngOnInit() {
		if (this.assets?.length) {
			this.hasAsset = true;
		}
	}

	async regenerateAnswer(message: Message) {
		if (message.userId) {
			return;
		}

		await this.promptService.promptAdd({
			id: this.server.id('msg'),
			prompt: message.content!,
			threadId: message.threadId as string,
			modelId: this.modelId as string,
		});
	}

	asset(message: Message): string {
		return ('data:image/png;base64,' +
			this.assets.find((a) => message.assetIds?.includes(a.id!))
				?.content) as string;
	}

	deleteMessage(messageId: string | undefined) {
		if (messageId === undefined) {
			return;
		}
		this.chatService.chatMessageDelete(messageId);
	}

	propagateCopyToClipboard(text: string | undefined) {
		if (text === undefined) {
			return;
		}
		this.onCopyToClipboard.emit(text);
	}
}

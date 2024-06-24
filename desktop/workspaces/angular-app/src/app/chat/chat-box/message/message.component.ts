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
	Input,
	Output,
	EventEmitter,
	ChangeDetectionStrategy,
} from '@angular/core';
import {
	ChatService,
	ChatMessage,
	Asset,
} from '../../../services/chat.service';
import { PromptService } from '../../../services/prompt.service';
import { LocaltronService } from '../../../services/localtron.service';
import { MarkdownComponent } from 'ngx-markdown';
import { IonicModule } from '@ionic/angular';
import { NgIf, DatePipe } from '@angular/common';

@Component({
	selector: 'app-message',
	templateUrl: './message.component.html',
	styleUrl: './message.component.scss',
	standalone: true,
	imports: [NgIf, IonicModule, MarkdownComponent, DatePipe],
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MessageComponent {
	constructor(
		private chatService: ChatService,
		private promptService: PromptService,
		private localtron: LocaltronService
	) {}
	hasAsset = false;

	@Input() message!: ChatMessage;
	@Input() assets: Asset[] = [];
	@Input() streaming: boolean = false;
	@Input() modelId: string = '';

	@Output() onCopyToClipboard = new EventEmitter<string>();

	ngOnInit() {
		if (this.assets?.length) {
			this.hasAsset = true;
		}
	}

	async regenerateAnswer(message: ChatMessage) {
		if (message.userId) {
			return;
		}

		await this.promptService.promptAdd({
			id: this.localtron.uuid(),
			prompt: message.content,
			threadId: message.threadId as string,
			modelId: this.modelId as string,
		});
	}

	asset(message: ChatMessage): string {
		return ('data:image/png;base64,' +
			this.assets.find((a) => message.assetIds?.includes(a.id))
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

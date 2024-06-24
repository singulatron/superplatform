import {
	Component,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
	Output,
	EventEmitter,
	ViewChild,
} from '@angular/core';
import { IonicModule } from '@ionic/angular';
import { Subscription } from 'rxjs';
import { TranslatePipe } from '../../../translate.pipe';
import { NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import {
	CharacterService,
	Character,
} from '../../../services/character.service';
import { CharacterComponent } from '../../character/character.component';
import { ModelService, Model } from '../../../services/model.service';
import { ConfigService } from '../../../services/config.service';

export interface SendOutput {
	message: string;
	characterId: string;
	modelId: string;
}

@Component({
	selector: 'app-chat-input',
	standalone: true,
	imports: [IonicModule, NgIf, FormsModule, TranslatePipe],
	templateUrl: './chat-input.component.html',
	styleUrl: './chat-input.component.scss',
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatInputComponent {
	private model: Model | undefined;
	private models: Model[] = [];
	public message: string = '';
	private subscriptions: Subscription[] = [];
	@ViewChild(CharacterComponent) characterModal!: CharacterComponent;

	@Output() sends = new EventEmitter<SendOutput>();

	constructor(
		private modelService: ModelService,
		private characterService: CharacterService,
		private configService: ConfigService,
		private cd: ChangeDetectorRef
	) {}

	async ngOnInit() {
		this.models = await this.modelService.getModels();
		this.cd.markForCheck();

		this.subscriptions.push(
			this.configService.onConfigUpdate$.subscribe(async (config) => {
				this.model = this.models?.find(
					(m) => m.id == config?.model?.currentModelId
				);
			})
		);
	}

	public hasNonWhiteSpace(value: string): boolean {
		if (!value) {
			return false;
		}
		return /\S/.test(value);
	}

	// Handle keydown event to differentiate between Enter and Shift+Enter
	handleKeydown(event: KeyboardEvent): void {
		if (event.key === 'Enter' && !event.shiftKey) {
			event.preventDefault();
			if (this.hasNonWhiteSpace(this.message)) {
				this.send();
			}
		} else if (event.key === 'Enter' && event.shiftKey) {
			event.preventDefault();
			this.message += '\n';
		}
	}

	async send() {
		let msg = this.message;
		this.message = '';

		const character = this.getSelectedCharacter();
		this.sends.next({
			characterId: character.id,
			message: msg,
			modelId: this.model?.id || '',
		});
	}

	public getSelectedCharacter(): Character {
		return this.characterService.selectedCharacter;
	}

	public showCharacterModal() {
		this.characterModal.show();
	}

	public getSelectedCharacterText(): string {
		const selectedCharacter = this.getSelectedCharacter();
		return selectedCharacter?.data?.name ? selectedCharacter.data.name : 'None';
	}
}

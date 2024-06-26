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
	Output,
	EventEmitter,
	ViewChild,
	ViewContainerRef,
	ComponentFactoryResolver,
	AfterViewInit,
	OnInit,
	ViewEncapsulation,
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
	styleUrls: ['./chat-input.component.scss'],
	encapsulation: ViewEncapsulation.None,
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ChatInputComponent implements OnInit, AfterViewInit {
	private model: Model | undefined;
	private models: Model[] = [];
	public message: string = '';
	private subscriptions: Subscription[] = [];
	private characterModal!: CharacterComponent;

	@ViewChild('modalContainer', { read: ViewContainerRef })
	modalContainer!: ViewContainerRef;
	@Output() sends = new EventEmitter<SendOutput>();

	constructor(
		private modelService: ModelService,
		private characterService: CharacterService,
		private configService: ConfigService,
		private cd: ChangeDetectorRef,
		private resolver: ComponentFactoryResolver
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

	ngAfterViewInit() {
		// Initial modal load if needed
		this.loadCharacterModal();
	}

	ionViewWillLeave() {
		for (const sub of this.subscriptions) {
			sub.unsubscribe();
		}
	}

	private loadCharacterModal() {
		try {
			const factory = this.resolver.resolveComponentFactory(CharacterComponent);
			this.modalContainer.clear();
			const componentReference = this.modalContainer.createComponent(factory);
			this.characterModal = componentReference.instance;
		} catch (error) {
			console.error('Error loading character modal:', error);
		}
	}

	public hasNonWhiteSpace(value: string): boolean {
		if (!value) {
			return false;
		}
		return /\S/.test(value);
	}

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
		const message = this.message;
		this.message = '';

		const character = this.getSelectedCharacter();
		this.sends.emit({
			characterId: character?.id,
			message: message,
			modelId: this.model?.id || '',
		});
	}

	public getSelectedCharacter(): Character {
		return this.characterService.selectedCharacter;
	}

	public showCharacterModal() {
		this.loadCharacterModal();
		if (this.characterModal) {
			this.characterModal.show();
		} else {
			console.error('Character modal is not available.');
		}
	}

	public getSelectedCharacterText(): string {
		const selectedCharacter = this.getSelectedCharacter();
		return selectedCharacter.data.name || 'None';
	}
}

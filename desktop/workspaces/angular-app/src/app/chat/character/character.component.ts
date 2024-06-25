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
	ViewEncapsulation,
	ViewChild,
	ChangeDetectionStrategy,
	ChangeDetectorRef,
} from '@angular/core';
import { IonModal } from '@ionic/angular';
import { NgFor, NgIf } from '@angular/common';
import { IonicModule } from '@ionic/angular';
import { FormsModule } from '@angular/forms';
import {
	CharacterService,
	Character,
	initCharacter,
} from '../../services/character.service';

@Component({
	selector: 'app-ai-character',
	templateUrl: './character.component.html',
	styleUrl: './character.component.scss',
	imports: [IonicModule, NgFor, NgIf, FormsModule],
	standalone: true,
	encapsulation: ViewEncapsulation.None,
	changeDetection: ChangeDetectionStrategy.OnPush,
})
export class CharacterComponent {
	@ViewChild(IonModal) modal!: IonModal;
	public isOpen: boolean = false;
	public segment = 'select';
	public editingCharacter: Character = initCharacter();
	public characters: Character[] = [];

	constructor(
		private characterService: CharacterService,
		private cd: ChangeDetectorRef
	) {}


	async ngOnInit() {
		await this.loadCharacters();
	}

	async loadCharacters() {
		this.characters = await this.characterService.loadCharacters();
	}

	public clearEditingCharacter(): void {
		this.editingCharacter = initCharacter();
	}

	selectCharacter(character: Character) {
		this.characterService.selectCharacter(character);
	}

	async upsertCharacter(character: Character) {
		await this.characterService.upsertCharacter(character);
		this.clearEditingCharacter();
		await this.loadCharacters();
	}

	async deleteCharacter(character: Character) {
		await this.characterService.deleteCharacter(character);
		await this.loadCharacters();
	}

	async selectCharacterForEdit(character: Character) {
		this.editingCharacter = character;
		this.segment = 'create';
	}

	getModeText(): string {
		return this.editingCharacter?.id ? 'Edit' : 'Create';
	}

	show(): void {
		this.isOpen = true;
		this.cd.markForCheck();
	}

	close(): void {
		this.isOpen = false;
		this.cd.markForCheck();
	}
}

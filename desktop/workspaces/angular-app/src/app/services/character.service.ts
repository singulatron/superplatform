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

import { Injectable } from '@angular/core';
import { GenericObject, GenericService, all, id as idCondition, userId as userIdCondition } from './generic.service';
import { LocaltronService } from './localtron.service';
import { UserService } from './user.service';
import { first } from 'rxjs';

const CHARACTERS_TABLE_NAME = "characters"
const SELECTED_CHARACTERS_TABLE_NAME = "selected-characters";

@Injectable({
	providedIn: 'root'
})
export class CharacterService {
	public selectedCharacter!: Character;

	constructor(
		private localtron: LocaltronService,
		private genericService: GenericService,
		private userService: UserService
	) {
		this.userService.user$.pipe(first()).subscribe(() => {
			this.init();
		});
	}

	async init() {
		let sc = await this.getSelectedCharacter();
		if (sc) {
			this.selectedCharacter = sc;
		}
	}

	async loadCharacters(): Promise<Character[]> {
		const response = await this.genericService.find(CHARACTERS_TABLE_NAME, [all()]);
		return response?.objects as Character[];
	}

	async upsertCharacter(character: Character) {
		const exists = await this.getCharacter(character.id);
		if (exists) {
			await this.updateCharacter(character);
		} else {
			await this.createNewCharacter(character);
		}
	}

	async createNewCharacter(character: Character) {
		const id = this.localtron.uuid();
		const now = new Date().toISOString();
		character.id = id;
		await this.genericService.create(CHARACTERS_TABLE_NAME, {
			...character,
			id,
			createdAt: now,
			updatedAt: now
		});
	}

	async deleteCharacter(character: Character) {
		if (character.id === this.selectedCharacter?.id) {
			this.selectedCharacter = {} as any;
			await this.deleteCharacterSelection(character.id)
		}
		await this.genericService.delete(CHARACTERS_TABLE_NAME, [idCondition(character.id)]);
	}

	async updateCharacter(character: Character) {
		const now = new Date().toISOString();
		character.updatedAt = now;
		this.genericService.update(CHARACTERS_TABLE_NAME, [idCondition(character.id)], {
			...character
		})
	}

	async getCharacter(characterId: string): Promise<Character | null> {
		try {
			const response = await this.genericService.find(CHARACTERS_TABLE_NAME, [idCondition(characterId)]);
			return response?.objects?.[0] as any as Character;
		} catch (e) {
			return null;
		}
	}

	async getSelectedCharacter(): Promise<Character | null> {
		const characterSelection = await this.getCharacterSelection();
		const selectedCharacterId = characterSelection?.data?.selectedCharacterId as any as string;
		if (!selectedCharacterId) {
			return null;
		}
		const character = await this.getCharacter(selectedCharacterId);
		if (character) {
			this.selectedCharacter = character;
		}
		return character;
	}

	async selectCharacter(character: Character): Promise<void> {
		this.selectedCharacter = character;
		this.upsertCharacterSelection(character.id);
	}

	async upsertCharacterSelection(selectedCharacterId: string): Promise<void> {
		const now = new Date().toISOString();
		let characterSelection = await this.getCharacterSelection();
		if (!characterSelection) {
			characterSelection = initCharacterSelection();
			characterSelection.id = this.localtron.uuid();
			characterSelection.createdAt = now;
		}
		characterSelection.updatedAt = now;
		characterSelection.data.selectedCharacterId = selectedCharacterId;
		this.genericService.upsert(SELECTED_CHARACTERS_TABLE_NAME, {
			...characterSelection
		})
	}

	async getCharacterSelection(): Promise<SelectedCharacter | null> {
		const userId = await this.userService.getUserId();
		const response = await this.genericService.find(SELECTED_CHARACTERS_TABLE_NAME, [userIdCondition(userId)]);
		return response?.objects?.[0] as SelectedCharacter;
	}

	/**
	 * This will remove only the selection from SELECTED_CHARACTERS_TABLE_NAME
	 */
	async deleteCharacterSelection(characterId: string) {
		this.genericService.delete(SELECTED_CHARACTERS_TABLE_NAME, [idCondition(characterId)]);
	}
}

export interface Character extends GenericObject {
	data: {
		name: string;
		behaviour: string;
		private: boolean;
	}
}

export interface SelectedCharacter  extends GenericObject{
	data: {
		selectedCharacterId: string;
	}
}

export function initCharacter(): Character {
	return {
		data: {
			name: '',
			behaviour: '',
			private: false,
		},
		id: '',
		createdAt: '',
		updatedAt: ''
	}
}

export function initCharacterSelection(): SelectedCharacter {
	return {
		data: {
			selectedCharacterId: ''
		},
		id: '',
		createdAt: '',
		updatedAt: ''
	}
}
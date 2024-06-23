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
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ChatBoxComponent } from './chat-box/chat-box.component';
import { CharacterComponent } from './character/character.component';
import { IonicModule } from '@ionic/angular';
import { StdlibModule } from '../stdlib/stdlib.module';
import { MarkdownModule, MarkdownService } from 'ngx-markdown';
import { MessageComponent } from './chat-box/message/message.component';

@NgModule({
    providers: [MarkdownService],
    exports: [ChatBoxComponent, CharacterComponent],
    imports: [CommonModule, StdlibModule, IonicModule, MarkdownModule.forChild(), ChatBoxComponent, MessageComponent],
})
export class AiModule {}

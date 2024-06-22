import { Component } from '@angular/core';
import { AdvancedModelExplorerComponent } from './advanced-model-explorer/advanced-model-explorer.component';
import { PageComponent } from '../../../shared/stdlib/components/page/page.component';
import { CenteredComponent } from '../../../shared/stdlib/components/centered/centered.component';
import { IconMenuComponent } from '../../../shared/stdlib/components/icon-menu/icon-menu.component';

@Component({
	selector: 'app-model-explorer',
	templateUrl: './model-explorer.component.html',
	styleUrl: './model-explorer.component.scss',
	standalone: true,
	imports: [
		IconMenuComponent,
		PageComponent,
		CenteredComponent,
		AdvancedModelExplorerComponent,
	],
})
export class ModelExplorerComponent {
	constructor() {}
	selectedPage: string = 'default';
}

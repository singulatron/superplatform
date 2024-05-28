import { Component } from '@angular/core';

@Component({
	selector: 'app-model-explorer',
	templateUrl: './model-explorer.component.html',
	styleUrl: './model-explorer.component.scss',
})
export class ModelExplorerComponent {
	constructor() {}
	selectedPage: string = 'default';
}

import { Component } from '@angular/core';
import { AdvancedModelExplorerComponent } from './advanced-model-explorer/advanced-model-explorer.component';
import { SidebarPageComponent } from '../../../shared/stdlib/components/sidebar-page/sidebar-page.component';

@Component({
    selector: 'app-model-explorer',
    templateUrl: './model-explorer.component.html',
    styleUrl: './model-explorer.component.scss',
    standalone: true,
    imports: [SidebarPageComponent, AdvancedModelExplorerComponent],
})
export class ModelExplorerComponent {
	constructor() {}
	selectedPage: string = 'default';
}

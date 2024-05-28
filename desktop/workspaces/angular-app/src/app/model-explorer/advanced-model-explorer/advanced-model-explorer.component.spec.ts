import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AdvancedModelExplorerComponent } from './advanced-model-explorer.component';

describe('AdvancedModelExplorerComponent', () => {
	let component: AdvancedModelExplorerComponent;
	let fixture: ComponentFixture<AdvancedModelExplorerComponent>;

	beforeEach(async () => {
		await TestBed.configureTestingModule({
			imports: [AdvancedModelExplorerComponent],
		}).compileComponents();

		fixture = TestBed.createComponent(AdvancedModelExplorerComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});

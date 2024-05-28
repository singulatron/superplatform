import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ModelExplorerComponent } from './model-explorer.component';

describe('ModelExplorerComponent', () => {
	let component: ModelExplorerComponent;
	let fixture: ComponentFixture<ModelExplorerComponent>;

	beforeEach(async () => {
		await TestBed.configureTestingModule({
			imports: [ModelExplorerComponent],
		}).compileComponents();

		fixture = TestBed.createComponent(ModelExplorerComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});

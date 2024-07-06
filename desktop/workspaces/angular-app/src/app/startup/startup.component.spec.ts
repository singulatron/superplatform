import { ComponentFixture, TestBed } from '@angular/core/testing';

import { StartupComponent } from './startup.component';

describe('StartupComponent', () => {
	let component: StartupComponent;
	let fixture: ComponentFixture<StartupComponent>;

	beforeEach(async () => {
		await TestBed.configureTestingModule({
			imports: [StartupComponent],
		}).compileComponents();

		fixture = TestBed.createComponent(StartupComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});

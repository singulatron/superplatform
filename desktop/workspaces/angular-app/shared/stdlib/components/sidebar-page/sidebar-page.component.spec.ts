import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PageComponent } from './page.component';

describe('PageComponent', () => {
	let component: PageComponent;
	let fixture: ComponentFixture<PageComponent>;

	beforeEach(() => {
		TestBed.configureTestingModule({
			declarations: [PageComponent],
		});
		fixture = TestBed.createComponent(PageComponent);
		component = fixture.componentInstance;
		fixture.detectChanges();
	});

	it('should create', () => {
		expect(component).toBeTruthy();
	});
});

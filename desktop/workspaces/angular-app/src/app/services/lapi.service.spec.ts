import { TestBed } from '@angular/core/testing';

import { LapiService } from './lapi.service';

describe('LapiService', () => {
	let service: LapiService;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		service = TestBed.inject(LapiService);
	});

	it('should be created', () => {
		expect(service).toBeTruthy();
	});
});

import { TestBed } from '@angular/core/testing';

import { LocaltronService } from './server.service';

describe('LocaltronService', () => {
	let service: LocaltronService;

	beforeEach(() => {
		TestBed.configureTestingModule({});
		service = TestBed.inject(LocaltronService);
	});

	it('should be created', () => {
		expect(service).toBeTruthy();
	});
});

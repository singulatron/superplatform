import { TestBed } from '@angular/core/testing';

import { GenericService } from './generic.service';

describe('GenericService', () => {
  let service: GenericService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GenericService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

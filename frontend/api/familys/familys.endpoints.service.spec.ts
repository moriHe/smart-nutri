import { TestBed } from '@angular/core/testing';

import { FamilysEndpointsService } from './familys.endpoints.service';

describe('FamilysEndpointsService', () => {
  let service: FamilysEndpointsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FamilysEndpointsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

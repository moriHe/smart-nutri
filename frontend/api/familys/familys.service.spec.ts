import { TestBed } from '@angular/core/testing';

import { FamilysService } from './familys.service';

describe('FamilysService', () => {
  let service: FamilysService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(FamilysService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

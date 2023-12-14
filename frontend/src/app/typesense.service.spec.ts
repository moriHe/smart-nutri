import { TestBed } from '@angular/core/testing';

import { TypesenseService } from './typesense.service';

describe('TypesenseService', () => {
  let service: TypesenseService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TypesenseService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

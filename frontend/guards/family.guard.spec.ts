import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { familyGuard } from './family.guard';

describe('familyGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) => 
      TestBed.runInInjectionContext(() => familyGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});

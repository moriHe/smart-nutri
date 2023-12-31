import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { fireAuthGuard } from './fire-auth.guard';

describe('fireAuthGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) => 
      TestBed.runInInjectionContext(() => fireAuthGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});

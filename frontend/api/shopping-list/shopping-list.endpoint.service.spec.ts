import { TestBed } from '@angular/core/testing';

import { ShoppingListEndpointService } from './shopping-list.endpoint.service';

describe('ShoppingListEndpointService', () => {
  let service: ShoppingListEndpointService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ShoppingListEndpointService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

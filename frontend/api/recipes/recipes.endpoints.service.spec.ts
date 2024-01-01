import { TestBed } from '@angular/core/testing';

import { RecipesEndpointsService } from './recipes.endpoints.service';

describe('RecipesEndpointsService', () => {
  let service: RecipesEndpointsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RecipesEndpointsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

import { TestBed } from '@angular/core/testing';

import { MealplansEndpointsService } from './mealplans.endpoints.service';

describe('MealplansEndpointsService', () => {
  let service: MealplansEndpointsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MealplansEndpointsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

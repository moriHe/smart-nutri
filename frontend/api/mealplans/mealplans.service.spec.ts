import { TestBed } from '@angular/core/testing';

import { MealplansService } from './mealplans.service';

describe('MealplansService', () => {
  let service: MealplansService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MealplansService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

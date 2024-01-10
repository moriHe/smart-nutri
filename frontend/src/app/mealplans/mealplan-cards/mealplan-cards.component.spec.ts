import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MealplanCardsComponent } from './mealplan-cards.component';

describe('MealplanCardsComponent', () => {
  let component: MealplanCardsComponent;
  let fixture: ComponentFixture<MealplanCardsComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [MealplanCardsComponent]
    });
    fixture = TestBed.createComponent(MealplanCardsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateMealplanBottomsheetComponent } from './create-mealplan-bottomsheet.component';

describe('CreateMealplanBottomsheetComponent', () => {
  let component: CreateMealplanBottomsheetComponent;
  let fixture: ComponentFixture<CreateMealplanBottomsheetComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateMealplanBottomsheetComponent]
    });
    fixture = TestBed.createComponent(CreateMealplanBottomsheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

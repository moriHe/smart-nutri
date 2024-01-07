import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateMealplanDialogComponent } from './create-mealplan-dialog.component';

describe('CreateMealplanDialogComponent', () => {
  let component: CreateMealplanDialogComponent;
  let fixture: ComponentFixture<CreateMealplanDialogComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateMealplanDialogComponent]
    });
    fixture = TestBed.createComponent(CreateMealplanDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

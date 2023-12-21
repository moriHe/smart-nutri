import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateRecipeDialogComponent } from './create-recipe-dialog.component';

describe('CreateRecipeDialogComponent', () => {
  let component: CreateRecipeDialogComponent;
  let fixture: ComponentFixture<CreateRecipeDialogComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateRecipeDialogComponent]
    });
    fixture = TestBed.createComponent(CreateRecipeDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

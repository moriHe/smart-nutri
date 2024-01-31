import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IngredientDatabaseComponent } from './ingredient-database.component';

describe('IngredientDatabaseComponent', () => {
  let component: IngredientDatabaseComponent;
  let fixture: ComponentFixture<IngredientDatabaseComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [IngredientDatabaseComponent]
    });
    fixture = TestBed.createComponent(IngredientDatabaseComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

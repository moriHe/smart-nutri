import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IngredientSourceComponent } from './ingredient-source.component';

describe('IngredientSourceComponent', () => {
  let component: IngredientSourceComponent;
  let fixture: ComponentFixture<IngredientSourceComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [IngredientSourceComponent]
    });
    fixture = TestBed.createComponent(IngredientSourceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

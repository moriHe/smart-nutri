import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SearchIngredientDialogComponent } from './search-ingredient-dialog.component';

describe('SearchIngredientDialogComponent', () => {
  let component: SearchIngredientDialogComponent;
  let fixture: ComponentFixture<SearchIngredientDialogComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [SearchIngredientDialogComponent]
    });
    fixture = TestBed.createComponent(SearchIngredientDialogComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

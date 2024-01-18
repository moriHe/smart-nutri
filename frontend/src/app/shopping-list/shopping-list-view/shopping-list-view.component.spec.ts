import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShoppingListViewComponent } from './shopping-list-view.component';

describe('ShoppingListViewComponent', () => {
  let component: ShoppingListViewComponent;
  let fixture: ComponentFixture<ShoppingListViewComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ShoppingListViewComponent]
    });
    fixture = TestBed.createComponent(ShoppingListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

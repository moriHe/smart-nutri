import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ShoppingListBellComponent } from './shopping-list-bell.component';

describe('ShoppingListBellComponent', () => {
  let component: ShoppingListBellComponent;
  let fixture: ComponentFixture<ShoppingListBellComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ShoppingListBellComponent]
    });
    fixture = TestBed.createComponent(ShoppingListBellComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

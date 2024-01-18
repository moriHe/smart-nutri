import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NotOnShoppingListViewComponent } from './not-on-shopping-list-view.component';

describe('NotOnShoppingListViewComponent', () => {
  let component: NotOnShoppingListViewComponent;
  let fixture: ComponentFixture<NotOnShoppingListViewComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [NotOnShoppingListViewComponent]
    });
    fixture = TestBed.createComponent(NotOnShoppingListViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

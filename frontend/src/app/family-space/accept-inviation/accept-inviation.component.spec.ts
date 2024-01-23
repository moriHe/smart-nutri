import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AcceptInviationComponent } from './accept-inviation.component';

describe('AcceptInviationComponent', () => {
  let component: AcceptInviationComponent;
  let fixture: ComponentFixture<AcceptInviationComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [AcceptInviationComponent]
    });
    fixture = TestBed.createComponent(AcceptInviationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

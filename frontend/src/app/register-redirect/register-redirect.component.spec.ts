import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterRedirectComponent } from './register-redirect.component';

describe('RegisterRedirectComponent', () => {
  let component: RegisterRedirectComponent;
  let fixture: ComponentFixture<RegisterRedirectComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [RegisterRedirectComponent]
    });
    fixture = TestBed.createComponent(RegisterRedirectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

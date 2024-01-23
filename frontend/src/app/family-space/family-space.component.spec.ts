import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FamilySpaceComponent } from './family-space.component';

describe('FamilySpaceComponent', () => {
  let component: FamilySpaceComponent;
  let fixture: ComponentFixture<FamilySpaceComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [FamilySpaceComponent]
    });
    fixture = TestBed.createComponent(FamilySpaceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GypoComponent } from './gypo.component';

describe('GypoComponent', () => {
  let component: GypoComponent;
  let fixture: ComponentFixture<GypoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GypoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(GypoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

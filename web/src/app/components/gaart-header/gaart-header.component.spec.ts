import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GaartHeaderComponent } from './gaart-header.component';

describe('GaartHeaderComponent', () => {
  let component: GaartHeaderComponent;
  let fixture: ComponentFixture<GaartHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GaartHeaderComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GaartHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

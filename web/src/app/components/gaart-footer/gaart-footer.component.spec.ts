import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GaartFooterComponent } from './gaart-footer.component';

describe('GaartFooterComponent', () => {
  let component: GaartFooterComponent;
  let fixture: ComponentFixture<GaartFooterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GaartFooterComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GaartFooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

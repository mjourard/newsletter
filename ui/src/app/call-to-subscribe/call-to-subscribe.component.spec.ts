import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CallToSubscribeComponent } from './call-to-subscribe.component';

describe('CallToSubscribeComponent', () => {
  let component: CallToSubscribeComponent;
  let fixture: ComponentFixture<CallToSubscribeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CallToSubscribeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CallToSubscribeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

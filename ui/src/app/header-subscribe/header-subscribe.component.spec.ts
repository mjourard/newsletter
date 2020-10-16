import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HeaderSubscribeComponent } from './header-subscribe.component';

describe('HeaderSubscribeComponent', () => {
  let component: HeaderSubscribeComponent;
  let fixture: ComponentFixture<HeaderSubscribeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HeaderSubscribeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderSubscribeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

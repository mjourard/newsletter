import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HeaderSubscribeComponent } from './header-subscribe.component';
import {SubscriptionService} from "../subscription.service";
import {ToastService} from "../toast.service";

describe('HeaderSubscribeComponent', () => {
  let component: HeaderSubscribeComponent;
  let fixture: ComponentFixture<HeaderSubscribeComponent>;

  beforeEach(async(() => {
    const subService = jasmine.createSpyObj('SubscriptionService', ['subscribeEmail']);
    const toastService = jasmine.createSpyObj('ToastService', ['success']);
    TestBed.configureTestingModule({
      declarations: [ HeaderSubscribeComponent ],
      providers: [
        {provide: SubscriptionService, useValue: subService},
        {provide: ToastService, useValue: toastService}
      ]
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

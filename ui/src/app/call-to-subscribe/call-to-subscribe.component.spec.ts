import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CallToSubscribeComponent } from './call-to-subscribe.component';
import {Subscription} from "rxjs";
import {SubscriptionService} from "../subscription.service";
import {ToastService} from "../toast.service";

describe('CallToSubscribeComponent', () => {
  let component: CallToSubscribeComponent;
  let fixture: ComponentFixture<CallToSubscribeComponent>;

  beforeEach(async(() => {
    const subService = jasmine.createSpyObj('SubscriptionService', ['subscribeEmail']);
    const toastService = jasmine.createSpyObj('ToastService', ['success']);
    TestBed.configureTestingModule({
      declarations: [ CallToSubscribeComponent ],
      providers: [
        {provide: SubscriptionService, useValue: subService},
        {provide: ToastService, useValue: toastService}
        ]
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

import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AboutUsComponent } from './about-us.component';
import {SubscriptionService} from "../subscription.service";
import {Observable, of} from "rxjs";
import {Subscriber} from "../subscriber";

describe('AboutUsComponent', () => {
  let component: AboutUsComponent;
  let fixture: ComponentFixture<AboutUsComponent>;

  let mockSubs: Subscriber[] = [
    {email: 'testing@test.com', addedtsutc: '2020-11-12 02:19:22'}
  ];


  beforeEach(async(() => {
    const subscriptionServiceStub = jasmine.createSpyObj('SubscriptionService', ['listSubscribers', 'subscribeEmail']);
    subscriptionServiceStub.listSubscribers.and.returnValue(of(mockSubs));
    TestBed.configureTestingModule({
      declarations: [ AboutUsComponent ],
      providers: [{provide: SubscriptionService, useValue: subscriptionServiceStub}]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AboutUsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

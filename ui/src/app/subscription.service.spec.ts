import { TestBed } from '@angular/core/testing';
import { SubscriptionService } from './subscription.service';
import {HttpClientModule} from "@angular/common/http";
import {EnvServiceProvider} from "./env.service.provider";

describe('SubscriptionService', () => {
  let service: SubscriptionService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientModule],
      providers: [EnvServiceProvider, SubscriptionService]
    });
    service = TestBed.inject(SubscriptionService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('#subscribeEmail should notify the email was successfully subscribed', (done: DoneFn) => {
    let email = `test_email${Date.now()}@hotmail.com`;
    service.subscribeEmail(email).subscribe(value => {
      expect(value).toEqual({success: true, message: 'Successfully subscribed ' + email});
      done();
    })
  })
});

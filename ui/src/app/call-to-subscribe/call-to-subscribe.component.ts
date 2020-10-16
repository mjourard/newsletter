import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {SubscriptionService} from "../subscription.service";
import {ToastService} from "../toast.service";
import {AppConstants} from "../app-constants";

@Component({
  selector: 'app-call-to-subscribe',
  templateUrl: './call-to-subscribe.component.html',
  styleUrls: ['./call-to-subscribe.component.css']
})
export class CallToSubscribeComponent implements OnInit {
  newEmail: string;
  loadingAnimation: string;
  constructor(
    private subscriptionService: SubscriptionService,
    private toastService: ToastService
  ) {
    this.loadingAnimation = AppConstants.LOADING_ANIMATION;
  }

  @Output() hideCallToSubscribe = new EventEmitter();

  ngOnInit(): void {
  }

  collapse() {
    this.hideCallToSubscribe.emit();
  }

  subscribe(email: string): void {
    email = email.trim();
    if (!email) {
      return;
    }
    this.subscriptionService.subscribeEmail(email).subscribe(opResult => {
      if (opResult.success) {
        this.toastService.success(opResult.message);
        this.hideCallToSubscribe.emit();
      } else {
        this.toastService.error(opResult.message);
      }
    });
  }

}

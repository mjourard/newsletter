import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {SubscriptionService} from "../subscription.service";
import {ToastService} from "../toast.service";

@Component({
  selector: 'app-call-to-subscribe',
  templateUrl: './call-to-subscribe.component.html',
  styleUrls: ['./call-to-subscribe.component.css']
})
export class CallToSubscribeComponent implements OnInit {
  newEmail: string;
  constructor(
    private subscriptionService: SubscriptionService,
    private toastService: ToastService
  ) {
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
    this.subscriptionService.subscribeEmail(email).subscribe(message => {
      this.toastService.add(message);
      this.hideCallToSubscribe.emit();
    });
  }

}

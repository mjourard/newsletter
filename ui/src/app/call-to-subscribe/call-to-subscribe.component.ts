import { Component, OnInit } from '@angular/core';
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

  ngOnInit(): void {
  }

  collapse() {
    //TODO: have this call a function from the home component for proper separation
  }

  subscribe(email: string): void {
    email = email.trim();
    if (!email) {
      return;
    }
    this.subscriptionService.subscribeEmail(email).subscribe(message => {
      this.toastService.add(message);
    });
  }

}

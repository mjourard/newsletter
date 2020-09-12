import { Component, OnInit } from '@angular/core';
import {SubscriptionService} from "../subscription.service";
import {Subscriber} from "../subscriber";
import {ToastService} from "../toast.service";


@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  newEmail: string;
  subscribers: Subscriber[];

  constructor(
    private subscriptionService: SubscriptionService,
    private toastService: ToastService
  ) {
  }

  ngOnInit(): void {
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

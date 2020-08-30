import {Component, OnInit} from '@angular/core';
import {SubscriptionService} from "../subscription.service";
import {Subscriber} from "../subscriber";
import {ToastService} from "../toast.service";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  newEmail: string;
  subscribers: Subscriber[];

  constructor(
    private subscriptionService: SubscriptionService,
    private toastService: ToastService
  ) {
  }

  ngOnInit(): void {
    this.listSubscribers();
  }

  subscribe(email: string): void {
    email = email.trim();
    if (!email) {
      return;
    }
    console.log("subscribing " + email);
    this.subscriptionService.subscribeEmail(email).subscribe(message => {
      this.toastService.add(message);
    });
  }

  listSubscribers(): void {
    this.subscriptionService.listSubscribers().subscribe(
      subs => this.subscribers = subs
    );
  }


}

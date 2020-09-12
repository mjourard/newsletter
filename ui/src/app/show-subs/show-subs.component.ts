import { Component, OnInit } from '@angular/core';
import {Subscriber} from "../subscriber";
import {SubscriptionService} from "../subscription.service";

@Component({
  selector: 'app-show-subs',
  templateUrl: './show-subs.component.html',
  styleUrls: ['./show-subs.component.css']
})
export class ShowSubsComponent implements OnInit {

  subscribers: Subscriber[];

  constructor(
    private subscriptionService: SubscriptionService
  ) {
  }

  ngOnInit(): void {
    this.listSubscribers();
  }

  listSubscribers(): void {
    this.subscriptionService.listSubscribers().subscribe(
      subs => this.subscribers = subs
    );
  }
}

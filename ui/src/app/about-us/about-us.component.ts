import { Component, OnInit } from '@angular/core';
import {SubscriptionService} from "../subscription.service";
import {Subscriber} from "../subscriber";

@Component({
  selector: 'app-about-us',
  templateUrl: './about-us.component.html',
  styleUrls: ['./about-us.component.css']
})
export class AboutUsComponent implements OnInit {

  subscribers: Subscriber[];

  constructor(
    private subscriptionService: SubscriptionService
  ) { }

  ngOnInit(): void {
    this.subscriptionService.listSubscribers().subscribe(subscribers => {
      this.subscribers = subscribers;
    })
  }

}

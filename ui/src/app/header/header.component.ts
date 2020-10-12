import {Component, OnInit, Output, EventEmitter} from '@angular/core';
import { Router, NavigationEnd }  from '@angular/router';
import {SubscriptionService} from "../subscription.service";
import {Subscriber} from "../subscriber";
import {ToastService} from "../toast.service";
declare var $: any;

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  newEmail: string;
  subscribers: Subscriber[];
  jqMobileHeader: any;

  @Output() mobileHeaderClicked = new EventEmitter();
  @Output() mobileHeaderClosed = new EventEmitter();

  constructor(
    private subscriptionService: SubscriptionService,
    private toastService: ToastService,
    private router: Router
  ) {
  }

  ngOnInit(): void {
    this.jqMobileHeader = $("#navbarToggleExternalContent");
    this.jqMobileHeader.on('hidden.bs.collapse', () => {
        this.mobileHeaderClosed.emit();
    });
    this.router.events.subscribe(e => {
      if (e instanceof NavigationEnd) {
        this.mobileHeaderClose();
      }
    })
  }

  mobileHeaderClose(): void {
    this.jqMobileHeader.collapse('hide');
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

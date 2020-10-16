import {Component, OnInit, Output, EventEmitter, ViewChild} from '@angular/core';
import { Router, NavigationEnd }  from '@angular/router';
import {Subscriber} from "../subscriber";
declare var $: any;

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {
  subscribers: Subscriber[];
  jqMobileHeader: any;

  @Output() mobileHeaderClicked = new EventEmitter();
  @Output() mobileHeaderClosed = new EventEmitter();

  constructor(
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


}

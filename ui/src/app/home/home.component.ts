import {Component, ElementRef, OnInit, AfterViewInit, Renderer2, ViewChild} from '@angular/core';
import {EnvService} from "../env.service";
declare var $: any;

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit, AfterViewInit {

  jqCallToSubscribeContainer: any;
  @ViewChild("callToSubscribeContainer") callToSubscribeContainer: ElementRef;
  @ViewChild("showCallToSubscribeBtn") showCallToSubscribeBtn: ElementRef;

  constructor(
    private renderer: Renderer2,
    private env: EnvService
  ) {
  }

  ngOnInit(): void {
    this.jqCallToSubscribeContainer = $("#collapse-call-to-action");
    this.jqCallToSubscribeContainer.on("hidden.bs.collapse", () => {
      this.renderer.addClass(this.callToSubscribeContainer.nativeElement, 'd-none');
    })
  }

  ngAfterViewInit(): void {
    if (this.env.callToSubscribeDevelopment === true) {
      this.renderer.removeClass(this.showCallToSubscribeBtn.nativeElement, 'd-none');
    }
  }

  hideCallToSubscribe(): void {
    this.jqCallToSubscribeContainer.collapse('hide');
  }

  showCallToSubscribe(): void {
    this.renderer.removeClass(this.callToSubscribeContainer.nativeElement, 'd-none');
    this.renderer.addClass(this.callToSubscribeContainer.nativeElement, 'show');
  }
}

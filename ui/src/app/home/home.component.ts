import {Component, ElementRef, OnInit, Renderer2, ViewChild} from '@angular/core';
declare var $: any;

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  jqCallToSubscribeContainer: any;
  @ViewChild("callToSubscribeParent") callToSubscribeParent: ElementRef;

  constructor(private renderer: Renderer2) {
  }

  ngOnInit(): void {
    this.jqCallToSubscribeContainer = $("#collapse-call-to-action");
    this.jqCallToSubscribeContainer.on("hidden.bs.collapse", () => {
      this.renderer.addClass(this.callToSubscribeParent.nativeElement, 'd-none');
    })
  }

  hideCallToSubscribe(): void {
    this.jqCallToSubscribeContainer.collapse('hide');
  }

  showCallToSubscribe(): void {
    this.renderer.removeClass(this.callToSubscribeParent.nativeElement, 'd-none');
  }
}

import {Component, ElementRef, Renderer2, ViewChild} from '@angular/core';
import {LogService} from "./log.service";
import {HeaderComponent} from "./header/header.component";
declare var $: any;
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'newsletter';
  @ViewChild("mobileOverlay") overlayDiv: ElementRef;
  @ViewChild("headerComponent") headerComp: HeaderComponent;

  constructor(
    private elementRef: ElementRef,
    private renderer: Renderer2,
    private logger: LogService
  ) {
  }

  mobileHeaderOpen(): void {
    this.logger.log("opening mobile header");
    this.renderer.addClass(this.overlayDiv.nativeElement, "active");
  }

  mobileHeaderClose(): void {
    this.logger.log("closing mobile header");
    this.headerComp.mobileHeaderClose();
  }

  removeMobileOverlay(): void {
    this.renderer.removeClass(this.overlayDiv.nativeElement, "active");
  }
}

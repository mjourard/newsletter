import {Component, ElementRef, OnInit, Renderer2, ViewChild} from '@angular/core';
import {SubscriptionService} from "../subscription.service";
import {ToastService} from "../toast.service";
import {AppConstants} from "../app-constants";

@Component({
  selector: 'app-header-subscribe',
  templateUrl: './header-subscribe.component.html',
  styleUrls: ['./header-subscribe.component.css']
})
export class HeaderSubscribeComponent implements OnInit {



  readonly HIDDEN_CSS_CLASS = 'd-none';
  readonly timeoutMS = 3000;
  loadingAnimationSrc: string;
  newEmail: string;
  @ViewChild("loadingAnimation") loadingAnimation: ElementRef;

  constructor(
    private subscriptionService: SubscriptionService,
    private toastService: ToastService,
    private renderer: Renderer2
  ) {
    this.loadingAnimationSrc = AppConstants.LOADING_ANIMATION;
  }

  ngOnInit(): void {
  }

  private setStateLoading() {
    this.renderer.removeClass(this.loadingAnimation.nativeElement, this.HIDDEN_CSS_CLASS);
  }

  private clearStateLoading() {
    this.renderer.addClass(this.loadingAnimation.nativeElement, this.HIDDEN_CSS_CLASS);
  }

  subscribe(email: string): void {
    email = email.trim();
    if (!email) {
      return;
    }
    this.setStateLoading();
    this.subscriptionService.subscribeEmail(email).subscribe(opResult => {
      if (opResult.success) {
        this.toastService.success(opResult.message)
      } else {
        this.toastService.error(opResult.message);
      }
    }, error => {
      this.toastService.error(error);
    }, () => {
      this.clearStateLoading();
    });
  }

}

import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { EnvServiceProvider } from "./env.service.provider";

import { AppRoutingModule } from './app-routing.module';
import { FormsModule } from "@angular/forms";
import { HttpClientModule } from "@angular/common/http";
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { NoopAnimationsModule, BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FooterComponent } from './footer/footer.component';
import { HomeComponent } from './home/home.component';
import { FaqComponent } from './faq/faq.component';
import { AboutUsComponent } from './about-us/about-us.component';
import { ContactUsComponent } from './contact-us/contact-us.component';
import { PrivacyPolicyComponent } from './privacy-policy/privacy-policy.component';
import { ToastComponent } from './toast/toast.component';
import { CallToSubscribeComponent } from './call-to-subscribe/call-to-subscribe.component';
import { ShowSubsComponent } from './show-subs/show-subs.component';
import { NewsletterHighlightsComponent } from './newsletter-highlights/newsletter-highlights.component';
import { HeaderSubscribeComponent } from './header-subscribe/header-subscribe.component';
import { FontAwesomeModule, FaIconLibrary } from '@fortawesome/angular-fontawesome';

//fontawesome imports - different file probably
import {faAngleDoubleRight, faTimes} from "@fortawesome/free-solid-svg-icons";

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    HomeComponent,
    FaqComponent,
    AboutUsComponent,
    ContactUsComponent,
    PrivacyPolicyComponent,
    ToastComponent,
    CallToSubscribeComponent,
    ShowSubsComponent,
    NewsletterHighlightsComponent,
    HeaderSubscribeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NoopAnimationsModule,
    BrowserAnimationsModule,
    FormsModule,
    HttpClientModule,
    FontAwesomeModule,
  ],
  providers: [EnvServiceProvider],
  bootstrap: [AppComponent]
})
export class AppModule {
  constructor(private library: FaIconLibrary) {
    library.addIcons(faTimes, faAngleDoubleRight);
  }
}
